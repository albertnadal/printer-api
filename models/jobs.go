package models

type Job struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Status string `json:"status"`
  Mesh string `json:"mesh"`
}

type JobDetails struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Status string `json:"status"`
  Mesh string `json:"mesh"`
  ETA int32 `json:"eta"`
  Elapsed int32 `json:"elapsed"`
  LayerThickness float32 `json:"layer_thickness"`
  TotalLayers int32 `json:"total_layers"`
  CurrentLayer int32 `json:"current_layer"`
}

type JobsList struct {
  Jobs []string `json:"jobs"`
}