package api

import (
  "fmt"
  //"encoding/json"
	"net/http"
  "strings"
  "time"
	"github.com/gorilla/mux"
  "printer-api/middleware"
  "printer-api/models"
)

func RegisterStatsRoutes(router *mux.Router, config models.Configuration) {
  router.Handle("/stats", middleware.BasicHandler(GetStats, config)).Methods("GET") // Deliver aggregation API stats
}

func GetStats(w http.ResponseWriter, r *http.Request) (int, uint64, error) {
  //processingSpeedValue := "?"
  //processingTimeLeft := "?"
  /*if pubSubConn.ProcessingSpeed > 0 {
    processingSpeedValue = fmt.Sprintf("%d ops/min.", pubSubConn.ProcessingSpeed)
    processingTimeLeft = durationToShortString(pubSubConn.ProcessingTimeLeft)
  }

  stats := &models.AggregationAPIStats{
    TotalReceivedOps: pubSubConn.TotalReceivedOps,
    TotalProcessedOps: pubSubConn.TotalProcessedOps,
    TotalOpsAwaiting: pubSubConn.TotalOpsAwaiting,
    ProcessingSpeed: processingSpeedValue,
    ProcessingTimeLeft: processingTimeLeft,
  }

  content, err := json.Marshal(stats);
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "{}")
		return http.StatusInternalServerError, uint64(len(content)), err
  }*/

  content := "hello world!"

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  fmt.Fprintf(w, string(content))
  return http.StatusOK, uint64(len(content)), nil
}

func durationToShortString(d time.Duration) string {
    s := d.String()
    if strings.HasSuffix(s, "m0s") {
        s = s[:len(s)-2]
    }
    if strings.HasSuffix(s, "h0m") {
        s = s[:len(s)-2]
    }
    return s
}
