package models

import (
	"time"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Server struct {
		Name string `yaml:"name"`
		Port int32 `yaml:"port"`
		APIPathPrefix string `yaml:"apiPathPrefix"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
		ReadTimeout time.Duration `yaml:"readTimeout"`
		IdleTimeout time.Duration `yaml:"idleTimeout"`
		AllowedMethods string `yaml:"allowedMethods"`
		AllowedOrigins string `yaml:"allowedOrigins"`
	} `yaml:"server"`
	MqttBroker struct {
		Host string `yaml:"host"`
		Port int32 `yaml:"port"`
	} `yaml:"mqttBroker"`
	Logger struct {
		Verbose bool `yaml:"verbose"`
	} `yaml:"logger"`
}

func (c *Configuration) LoadConfiguration(filename string) error {
    yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        return err
    }
    return nil
}
