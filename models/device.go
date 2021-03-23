package models

type MQTTBrokerData struct {
	Host string `json:"host"`
	Port int64 `json:"port"`
}

type DeviceInfo struct {
	Name string `json:"name"`
	Status string `json:"status"`
	MqttBroker MQTTBrokerData `json:"mqttBroker"`
	Jobs []Job `json:"jobs"`
}
