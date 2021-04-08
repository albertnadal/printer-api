package api

import (
  "fmt"
  "encoding/json"
	"net/http"
	"github.com/gorilla/mux"
  "printer-api/middleware"
  "printer-api/models"
  "printer-api/managers"
)

func RegisterDeviceRoutes(router *mux.Router, printerManager managers.PrinterManager, config models.Configuration) {
  router.Handle("/info", middleware.BasicHandler(GetDeviceInfo, printerManager, config)).Methods("GET") // Deliver device information
  router.Handle("/reset", middleware.BasicHandler(ResetDevice, printerManager, config)).Methods("POST") // Resets the device
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

func ResetDevice(w http.ResponseWriter, r *http.Request, printerManager managers.PrinterManager) (int, uint64, error) {
  info := printerManager.ResetDevice()
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