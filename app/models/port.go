package models

import (
	"encoding/json"
	"time"
)

type PortMap map[string]*Port

func (p *PortMap) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

type Port struct {
	// Object key in payload. Purposefully not called it Id to avoid collisions with aa real DB client
	Key         string
	Name        string      `json:"name"`
	City        string      `json:"city"`
	Country     string      `json:"country"`
	Alias       Alias       `json:"alias"`
	Regions     Regions     `json:"regions"`
	Coordinates Coordinates `json:"coordinates"`
	Province    string      `json:"province"`
	Timezone    string      `json:"timezone"`
	Unlocs      Unlocs      `json:"unlocs"`
	Code        string      `json:"code"`
}

func (p *Port) GetTimezoneLocation() (*time.Location, error) {
	return time.LoadLocation(p.Timezone)
}

type Alias []string
type Regions []string
type Coordinates []float32
type Unlocs []string
