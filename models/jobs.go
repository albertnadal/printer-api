package models

type Job struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Status string `json:"status"`
  MqttTopic string `json:"mqttTopic"`
  Thumbnail string `json:"thumbnail"`
  Mesh string `json:"mesh"`
}
