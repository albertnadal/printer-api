package api

import (
  "fmt"
  "encoding/json"
  "encoding/binary"
	"net/http"
	"github.com/gorilla/mux"
  "printer-api/middleware"
  "printer-api/models"
  "printer-api/managers"
)

func RegisterJobsRoutes(router *mux.Router, printerManager managers.PrinterManager, config models.Configuration) {
  router.Handle("/jobs", middleware.BasicHandler(GetJobs, printerManager, config)).Methods("GET") // Deliver list of available jobs
  router.Handle("/jobs/{job_id}", middleware.BasicHandler(GetJobDetails, printerManager, config)).Methods("GET") // Deliver the job details
  router.Handle("/jobs/{job_id}/stl", middleware.BasicHandler(GetJobSTL, printerManager, config)).Methods("GET") // Deliver the STL mesh file of the job
  router.Handle("/jobs/{job_id}/start", middleware.BasicHandler(StartJob, printerManager, config)).Methods("POST") // Starts the job
  router.Handle("/jobs/{job_id}/cancel", middleware.BasicHandler(CancelJob, printerManager, config)).Methods("POST") // Cancel the job
}

func GetJobs(w http.ResponseWriter, r *http.Request, printerManager managers.PrinterManager) (int, uint64, error) {
  jobs := printerManager.GetJobs()
  content, err := json.Marshal(jobs);
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "")
		return http.StatusInternalServerError, 0, err
  }

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  fmt.Fprintf(w, string(content))
  return http.StatusOK, uint64(len(content)), nil
}

func GetJobDetails(w http.ResponseWriter, r *http.Request, printerManager managers.PrinterManager) (int, uint64, error) {
  vars := mux.Vars(r)
  job, err := printerManager.GetJobDetails(vars["job_id"])
  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "")
		return http.StatusNotFound, 0, err
  }

  content, err := json.Marshal(job);
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "")
		return http.StatusInternalServerError, 0, err
  }

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  fmt.Fprintf(w, string(content))
  return http.StatusOK, uint64(len(content)), nil
}

func GetJobSTL(w http.ResponseWriter, r *http.Request, printerManager managers.PrinterManager) (int, uint64, error) {
  vars := mux.Vars(r)
  STLbytes, err := printerManager.GetJobSTLbytes(vars["job_id"])
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "")
		return http.StatusInternalServerError, 0, err
  }

  w.Header().Set("Content-Type", "model/stl")
  w.Header().Set("Content-Disposition", "inline; filename=\""+vars["job_id"]+".stl\"")
  STLbytes.WriteTo(w)
  return http.StatusOK, uint64(binary.Size(STLbytes)), nil
}

func StartJob(w http.ResponseWriter, r *http.Request, printerManager managers.PrinterManager) (int, uint64, error) {
  vars := mux.Vars(r)
  job, err := printerManager.StartJob(vars["job_id"])
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "")
		return http.StatusBadRequest, 0, err
  }

  content, err := json.Marshal(job);
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "")
		return http.StatusInternalServerError, 0, err
  }

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  fmt.Fprintf(w, string(content))
  return http.StatusOK, uint64(len(content)), nil
}

func CancelJob(w http.ResponseWriter, r *http.Request, printerManager managers.PrinterManager) (int, uint64, error) {
  vars := mux.Vars(r)
  job, err := printerManager.CancelJob(vars["job_id"])
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "")
		return http.StatusBadRequest, 0, err
  }

  content, err := json.Marshal(job);
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "")
		return http.StatusInternalServerError, 0, err
  }

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  fmt.Fprintf(w, string(content))
  return http.StatusOK, uint64(len(content)), nil
}