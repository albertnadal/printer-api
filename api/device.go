package api

import (
  "fmt"
  "encoding/json"
	"net/http"
  "strings"
  "time"
	"github.com/gorilla/mux"
  "printer-api/middleware"
  "printer-api/models"
  "printer-api/managers"
)

func RegisterDeviceRoutes(router *mux.Router, printerManager managers.PrinterManager, config models.Configuration) {
  router.Handle("/info", middleware.BasicHandler(GetDeviceInfo, printerManager, config)).Methods("GET") // Deliver device information
}

func GetDeviceInfo(w http.ResponseWriter, r *http.Request, printerManager managers.PrinterManager) (int, uint64, error) {
  info := printerManager.GetDeviceInfo()
  content, err := json.Marshal(info);
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "")
		return http.StatusInternalServerError, 0, err
  }

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
