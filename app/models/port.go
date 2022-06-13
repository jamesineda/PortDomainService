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

// GetChanges TODO make dynamic to struct field change using reflection to grab fields and their underlying values
func (p *Port) GetChanges(p2 *Port) map[string]interface{} {
	changes := map[string]interface{}{}

	if p.Name != p2.Name {
		changes["name"] = p2.Name
	}

	if p.City != p2.City {
		changes["city"] = p2.City
	}

	if p.Country != p2.Country {
		changes["country"] = p2.Country
	}

	if !p.Alias.Equals(p2.Alias) {
		changes["alias"] = p2.Alias
	}

	if !p.Regions.Equals(p2.Regions) {
		changes["regions"] = p2.Regions
	}

	if !p.Coordinates.Equals(p2.Coordinates) {
		changes["coordinates"] = p2.Coordinates
	}

	if p.Province != p2.Province {
		changes["province"] = p2.Province
	}

	if p.Timezone != p2.Timezone {
		changes["timezone"] = p2.Timezone
	}

	if !p.Unlocs.Equals(p2.Unlocs) {
		changes["unlocs"] = p2.Unlocs
	}

	if p.Code != p2.Code {
		changes["code"] = p2.Code
	}
	return changes
}

// horrible code duplication
type Alias []string

func (a *Alias) Equals(b Alias) bool {
	for i, aa := range *a {
		if aa != b[i] {
			return false
		}
	}

	if len(*a) != len(b) {
		return false
	}
	return true
}

type Regions []string

func (a *Regions) Equals(b Regions) bool {
	for i, aa := range *a {
		if aa != b[i] {
			return false
		}
	}

	if len(*a) != len(b) {
		return false
	}
	return true
}

type Coordinates []float32

func (a *Coordinates) Equals(b Coordinates) bool {
	for i, aa := range *a {
		if aa != b[i] {
			return false
		}
	}

	if len(*a) != len(b) {
		return false
	}
	return true
}

type Unlocs []string

func (a *Unlocs) Equals(b Unlocs) bool {
	for i, aa := range *a {
		if aa != b[i] {
			return false
		}
	}

	if len(*a) != len(b) {
		return false
	}
	return true
}
