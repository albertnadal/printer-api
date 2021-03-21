package models

import (
	"log"
	"plugin"
)

// IPlugin is the interface a plugin MUST implement, otherwise the plugin will crash
type IPlugin interface {
	Init(string) error
	Info() (*PluginInfo, error)
  ProcessPacket(packet *Packet) error
	Close() error
}

type PluginInfo struct {
	Name    string
	Version string
	Labels  []string
}

type Plugin struct {
	Filename    string
	Folder      string
	Information *PluginInfo
	Instance    *plugin.Plugin
	IPlugin
}

type PluginCatalog struct {
	Plugins   []Plugin
	QueuesMap map[string][]Plugin
}

/* This functions closes all active plugins */
func (c *PluginCatalog) Close() {
	for _, plugin := range c.Plugins {
		plugin.Close()
	}
}

/* This function initializes the plugin */
func (p *Plugin) Init(pluginFolder string) error {
	initFunc, err := p.Instance.Lookup("Init")
	if err != nil {
		log.Fatalf("Plugin does not implement IPlugin interface: %s: %v", p.Filename, err)
		return err
	}

	return initFunc.(func(string) error)(pluginFolder)
}

/* This function returns information of the plugin */
func (p *Plugin) Info() (*PluginInfo, error) {
	infoFunc, err := p.Instance.Lookup("Info")
	if err != nil {
		log.Fatalf("Plugin does not implement IPlugin interface: %s: %v", p.Filename, err)
		return nil, err
	}

	return infoFunc.(func() (*PluginInfo, error))()
}

/* This function process a packet of data */
func (p *Plugin) ProcessPacket(packet *Packet) error {
  processPacketFunc, err := p.Instance.Lookup("ProcessPacket")
	if err != nil {
		log.Fatalf("Plugin does not implement IPlugin interface: %s: %v", p.Filename, err)
		return err
	}

	return processPacketFunc.(func(*Packet) (error))(packet)
}

/* This functions should be called at the end of the lifecicle of the plugin */
func (p *Plugin) Close() error {
	closeFunc, err := p.Instance.Lookup("Close")
	if err != nil {
		log.Fatalf("Plugin does not implement IPlugin interface: %s: %v", p.Filename, err)
		return err
	}

	return closeFunc.(func() error)()
}
