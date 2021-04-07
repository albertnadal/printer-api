package managers

import (
  "fmt"
  "time"
  "errors"
  mqtt "github.com/eclipse/paho.mqtt.golang"
  "bytes"
  "encoding/json"
  "io/ioutil"
  "printer-api/models"
)

type PrinterManager struct {
  Config models.Configuration
  Device *models.DeviceInfo
  MQTTClient mqtt.Client
  Jobs map[string]*models.JobDetails
}

func (pm *PrinterManager) PrintJob(job *models.JobDetails) {
  start := time.Now()
  for true {
    message := fmt.Sprintf("{\"status\": \"%s\", \"job\": {\"id\": \"%s\", \"status\": \"%s\", \"layer\": %d}}", pm.Device.Status, job.Id, job.Status, job.CurrentLayer)
    token := pm.MQTTClient.Publish(pm.Config.Server.Id, 0, false, message)
    token.Wait()
  
    // Printing speed is the time in milliseconds needed to print a layer (lower is faster)
    time.Sleep(time.Duration(pm.Config.Server.PrintSpeed) * time.Millisecond)
    job.Elapsed = int32(time.Now().Sub(start).Seconds())
    job.CurrentLayer++

    // Check if the job has been cancelled
    if(job.Status == "cancelled") {
      return
    }
  
    // Check if the job has been completed
    if(job.CurrentLayer >= job.TotalLayers) {
      pm.Device.Status = "idle"
      job.Status = "completed"
      return
    }
  }
}

func (pm *PrinterManager) LoadJobsFromDisk() {
  files, err := ioutil.ReadDir("./jobs")
  check(err)

  pm.Jobs = make(map[string]*models.JobDetails)
  for _, file := range files {
    fmt.Printf("Loading job with id %s... ", file.Name())
    var job models.JobDetails
    jsonData, err := ioutil.ReadFile("./jobs/" + file.Name() + "/data.json")
    check(err)
    err = json.Unmarshal(jsonData, &job)
    check(err)
    job.ETA = int32(float64(job.TotalLayers) * (float64(pm.Config.Server.PrintSpeed) / float64(1000)))
    pm.Jobs[file.Name()] = &job
    fmt.Printf("done\n")
  }
}

func (pm *PrinterManager) GetDeviceInfo() (*models.DeviceInfo) {
  fmt.Printf("GetDeviceInfo DEVICE INFO (%v)\n", pm.Device)
  return pm.Device
}

func (pm *PrinterManager) GetJobs() (models.JobsList) {
  jobs_list := models.JobsList{}
  for _, job := range pm.Jobs {
    jobs_list.Jobs = append(jobs_list.Jobs, job.Id)
  }
  return jobs_list
}

func (pm *PrinterManager) StartJob(jobId string) (models.JobDetails, error) {
  job, found := pm.Jobs[jobId]
  if !found {
    return *job, errors.New("Job does not exists")
  }

  pm.Device.Status = "printing"
  job.Status = "printing"
  job.CurrentLayer = 0
  job.Elapsed = 0

  // Start printing the job in a goroutine
  go pm.PrintJob(job)

  return *job, nil
}

func (pm *PrinterManager) CancelJob(jobId string) (models.JobDetails, error) {
  job, found := pm.Jobs[jobId]
  if !found {
    return *job, errors.New("Job does not exists")
  }

  pm.Device.Status = "idle"
  job.Status = "cancelled"

  return *job, nil
}

func (pm *PrinterManager) GetJobDetails(jobId string) (models.JobDetails, error) {
  job, found := pm.Jobs[jobId]
  if !found {
    return *job, errors.New("Job does not exists")
  }

  return *job, nil
}

func (pm *PrinterManager) GetJobSTLbytes(jobId string) (*bytes.Buffer, error) {
  streamSTLbytes, err := ioutil.ReadFile("./jobs/"+jobId+"/mesh.stl")

  if err != nil {
    return nil, err
  }

  return bytes.NewBuffer(streamSTLbytes), nil
}

func (pm *PrinterManager) onConnectMQTT(client mqtt.Client) {
  //fmt.Printf("Connected to MQTT broker at %s:%d", pm.Config.MqttBroker.Host, pm.Config.MqttBroker.Port)
}

func (pm *PrinterManager) onConnectLostMQTT(client mqtt.Client, err error) {
  fmt.Printf("Connection lost to MQTT broker at %s:%d.\nError: %v\n", pm.Config.MqttBroker.Host, pm.Config.MqttBroker.Port, err)
}

func InitPrinterManager(config models.Configuration) (PrinterManager) {
  printer := PrinterManager{
    Config: config,
    Device: &models.DeviceInfo{
      Id: config.Server.Id,
      Name: config.Server.Name,
      Status: "idle",
      MqttTopic: config.Server.Id,
      MqttBroker: models.MQTTBrokerData{
        Host: config.MqttBroker.Host,
        Port: config.MqttBroker.Port,
      },
    },
  }

  printer.LoadJobsFromDisk()

  opts := mqtt.NewClientOptions()
  opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.MqttBroker.Host, config.MqttBroker.Port))
  opts.SetClientID(config.Server.Id)
  opts.OnConnect = printer.onConnectMQTT
  opts.OnConnectionLost = printer.onConnectLostMQTT
  printer.MQTTClient = mqtt.NewClient(opts)

  fmt.Printf("Connecting to MQTT broker at %s:%d... ", config.MqttBroker.Host, config.MqttBroker.Port)
  if token := printer.MQTTClient.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }
  fmt.Printf("done\n")

  return printer
}

func check(e error) {
  if e != nil {
      panic(e)
  }
}
