package models

type MQTTBrokerData struct {
	Host string `json:"host"`
	Port int32 `json:"port"`
}

type DeviceInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	MqttTopic string `json:"mqttTopic"`
	MqttBroker MQTTBrokerData `json:"mqttBroker"`
	//Jobs []Job `json:"jobs"`
}
