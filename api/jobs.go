package api

import (
  //"fmt"
  //"encoding/json"
  //"io/ioutil"
	"net/http"
  //"strconv"
	"github.com/gorilla/mux"
  "printer-api/middleware"
  "printer-api/models"
  "printer-api/managers"
)

func RegisterJobsRoutes(router *mux.Router, printerManager managers.PrinterManager, config models.Configuration) {
  router.Handle("/jobs/{job_id}/start", middleware.BasicHandler(StartJob, printerManager, config)).Methods("POST", "PUT") // Creates a monthly aggregate
}

func StartJob(w http.ResponseWriter, r *http.Request, printerManager managers.PrinterManager) (int, uint64, error) {
/*
  statusCode := http.StatusOK
	content := "OK"

  vars := mux.Vars(r)
  var parameters []models.PacketParameter
  parameters = append(parameters, models.PacketParameter{Key: "company_id", Value: vars["company_id"]})
  parameters = append(parameters, models.PacketParameter{Key: "contract_id", Value: vars["contract_id"]})
	parameters = append(parameters, models.PacketParameter{Key: "year", Value: vars["year"]})
	parameters = append(parameters, models.PacketParameter{Key: "month", Value: vars["month"]})

  value, _ := strconv.ParseInt(vars["year"], 10, 32)
  if year := int(value); (year < 2000) || (year >= 2050) {
    content = "Invalid year"
    fmt.Fprintf(w, content)
    return http.StatusBadRequest, uint64(len(content)), nil
  }

  value, _ = strconv.ParseInt(vars["month"], 10, 32)
  if year := int(value); (year < 1) || (year > 12) {
    content = "Invalid month"
    fmt.Fprintf(w, content)
    return http.StatusBadRequest, uint64(len(content)), nil
  }

  defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		content = "Cannot read request body"
		fmt.Fprintf(w, content)
		return http.StatusInternalServerError,  uint64(len(content)), err
	}

  var aggregate_contract_month models.AggregateContractMonth
	if err := json.Unmarshal(body, &aggregate_contract_month); err != nil {
    w.WriteHeader(http.StatusBadRequest)
		content = "Invalid JSON"
    fmt.Fprintf(w, content)
		return http.StatusBadRequest, uint64(len(content)), err
	}

  //packet := &models.Packet{Action: "upsert", Labels: []string{"reports"}, Model: "aggregate_contract_month", Payload: body, SourceEndpoint: r.URL.Path, Parameters: parameters}
  //pubSubConn.Publish(packet)

  fmt.Fprintf(w, content)
  return statusCode, uint64(len(content)), nil*/
  return 200, 0, nil
}
