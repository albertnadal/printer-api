package managers

import (
  "fmt"
  //"strconv"
  //"os"
  //"log"
  "encoding/json"
  "io/ioutil"
  "printer-api/models"
)

type PrinterManager struct {
  Device models.DeviceInfo
  Jobs map[string]models.Job
}

func (pm *PrinterManager) LoadJobsFromDisk() {
  files, err := ioutil.ReadDir("./jobs")
  check(err)

  for _, file := range files {
    fmt.Printf("Loading job with id %s... ", file.Name())
    var job models.Job
    jsonData, err := ioutil.ReadFile("./jobs/" + file.Name() + "/data.json")
    check(err)
    err = json.Unmarshal(jsonData, &job)
    check(err)
    pm.Jobs[file.Name()] = job
    fmt.Printf("done\n")
  }
}

func (pm *PrinterManager) GetDeviceInfo() (models.DeviceInfo) {
  return pm.Device
}

func InitPrinterManager(config models.Configuration) (PrinterManager) {
  printer := PrinterManager{
    Jobs: make(map[string]models.Job),
  }
  printer.Device.Id = config.Server.Id
  printer.Device.Name = config.Server.Name
  printer.Device.Status = "idle"
  printer.Device.MqttTopic = config.Server.Id
  printer.Device.MqttBroker.Host = config.MqttBroker.Host
  printer.Device.MqttBroker.Port = config.MqttBroker.Port
  printer.LoadJobsFromDisk()
  return printer
}

func check(e error) {
  if e != nil {
      panic(e)
  }
}
/*
func (pm *PrinterManager) ListenWorkerOpsStats() {
  ps.NatsConn.QueueSubscribe("worker_stats", "workers", func(status_msg *stan.Msg) {
    var worker_stats models.WorkerStats
    err := json.Unmarshal(status_msg.Data, &worker_stats)

    if(err == nil) {
      ps.ProcessingSpeed = worker_stats.ProcessingSpeed
      ps.TotalProcessedOps = worker_stats.TotalProcessedOps
      if ps.TotalProcessedOps > ps.TotalReceivedOps {
        ps.TotalProcessedOps = ps.TotalReceivedOps;
      }
      ps.TotalOpsAwaiting = ps.TotalReceivedOps - ps.TotalProcessedOps
      seconds_left := (ps.TotalOpsAwaiting * 60) / uint64(Max(ps.ProcessingSpeed, 1))
      ps.ProcessingTimeLeft = time.Duration(seconds_left)*time.Second
    }

  }, stan.StartWithLastReceived())
}

func (ps *PubSubConnection) ListenAndProcessQueues(config models.Configuration, processPacket func(*models.Packet, string, error)) {

  pool := slaves.NewPool(20, func(paramsObj interface{}) {
    processPacket(paramsObj.(SlaveParams).Packet, paramsObj.(SlaveParams).Subject, paramsObj.(SlaveParams).Err)
  })

  for _, queue := range config.Worker.Queues {

    sub, err := ps.NatsConn.QueueSubscribe(queue, config.Nats.QueueGroup, func (msg *stan.Msg) {
      if msg.Subject == "reports" {
        var packet models.Packet
        err := json.Unmarshal(msg.Data, &packet)

        if msg.Sequence < config.LastProcessedMessageSequence {
          // Just ignore duplicated old messages
          log.Println(" [ LOWER SEQUENCE ] Received seq:", msg.Sequence, " Expected seq:", config.LastProcessedMessageSequence + 1, " This operation has been processed. Operation ignored.")
          msg.Ack()
        } else {
          if msg.Sequence > config.LastProcessedMessageSequence+1 {
            log.Println(" [ ERROR ] Received seq:", msg.Sequence, " Expected seq:", config.LastProcessedMessageSequence + 1, " Some queued operations may have been dropped due to queue size overflow.")
          }
          packet.Sequence = msg.Sequence
          pool.Serve(SlaveParams{&packet, msg.Subject, err})
          config.LastProcessedMessageSequence = msg.Sequence
          ps.TotalProcessedOps++
          opsProcessedDuringInterval++
          msg.Ack()
        }
      }

    }, stan.DeliverAllAvailable(), stan.DurableName(queue+"_"+config.Nats.DurableName), stan.SetManualAckMode(), stan.MaxInflight(10), stan.AckWait(10*time.Second))

    if err != nil {
  		ps.NatsConn.Close()
  		log.Fatal(err)
  	}

    if queue == "reports" {
      // Send processed operations count every second
      go doEvery(time.Second, func () {
        var worker_stats models.WorkerStats
        worker_stats.TotalProcessedOps = ps.TotalProcessedOps
        worker_stats.ProcessingSpeed = opsProcessedEveryFiveMinutes / 5

        serialized_worker_stats, err := json.Marshal(worker_stats)
        if(err == nil) {
          ps.NatsConn.Publish("worker_stats", serialized_worker_stats)
        }
      })

      // Periodic function used to calculate processing speed
      go doEvery(5*time.Minute, func () {
        opsProcessedEveryFiveMinutes = opsProcessedDuringInterval
        opsProcessedDuringInterval = 0
      })
    }

    ps.Subscriptions = append(ps.Subscriptions, sub)
    log.Printf("Listening on queue=[%s], clientID=[%s], qgroup=[%s] durable=[%s]\n", queue, config.Nats.ClientId, config.Nats.QueueGroup, config.Nats.DurableName)
  }
}

func (ps *PubSubConnection) Publish(packet *models.Packet) error {
  serialized_packet, err := json.Marshal(packet)
  if(err != nil) {
    return err
  }

  for _, label := range packet.Labels {
    if label == "reports" {
      //log.Println("Publishing packet in "+label+" queue")
      ps.NatsConn.Publish(label, serialized_packet)
      ps.TotalReceivedOps++
    }
  }

  return nil
}

func Connect(config models.Configuration) (*PubSubConnection, error) {
  uri := os.Getenv("NATS_URI");
  if uri == "" {
    uri = "nats://"+config.Nats.Host+":"+itoa(config.Nats.Port)
  }

  sc, err := stan.Connect(config.Nats.ClusterId, config.Nats.ClientId, stan.NatsURL(uri), stan.MaxPubAcksInflight(10), stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
  			log.Fatalf("[ ERROR ] NATS connection lost, reason: %v", reason)
  }))

  if(err == nil) {
    fmt.Println("Connected to "+uri)
  }

  conn := &PubSubConnection{Platform: "nats-streaming", Host: config.Nats.Host, Port: config.Nats.Port, URI: uri, NatsConn: sc, TotalReceivedOps: 0, TotalProcessedOps: 0, TotalOpsAwaiting: 0, ProcessingSpeed: 0}
  return conn, err
}

func Disconnect(pubSubConn *PubSubConnection, config models.Configuration) {
  if config.Nats.DurableName == "" || config.Nats.UnsubscribeOnExit {
    log.Println("Unsubscribing from active subscriptions...")
    for _, sub := range pubSubConn.Subscriptions {
        sub.Unsubscribe()
    }
  }

  log.Println("Disconnected from "+pubSubConn.URI)
  pubSubConn.NatsConn.Close()
}

func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func itoa(n int32) string {
    return strconv.FormatInt(int64(n), 10)
}
*/