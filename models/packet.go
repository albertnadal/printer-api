package models

import (
  "strconv"
)

type PacketParameter struct {
  Key string `json:"key"`
  Value string `json:"value"`
}

type Packet struct {
  Action string `json:"action"`
  Labels []string `json:"labels"`
  Model string `json:"model"`
	Payload []byte `json:"payload"`
  SourceEndpoint string `json:"sourceEndpoint"`
	Parameters []PacketParameter
  Sequence uint64
}

func (p *Packet) GetContractId() string {
  if (p.Parameters != nil) {
    for _, parameter := range p.Parameters {
      if (parameter.Key == "contract_id") {
    		return parameter.Value
    	}
    }
  }
  return ""
}

func (p *Packet) GetCompanyId() int64 {
  return p.GetIntValueWithKey("company_id")
}

func (p *Packet) GetYear() int64 {
  return p.GetIntValueWithKey("year")
}

func (p *Packet) GetMonth() int64 {
  return p.GetIntValueWithKey("month")
}

func (p *Packet) GetIntValueWithKey(key string) int64 {
  if (p.Parameters != nil) {
    for _, parameter := range p.Parameters {
      if (parameter.Key == key) {
        value, err := strconv.ParseInt(parameter.Value, 10, 64)
        if(err == nil) {
          return value
        }
    	}
    }
  }
  return 0
}
