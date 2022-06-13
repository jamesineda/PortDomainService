package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPort_GetTimezoneLocation(t *testing.T) {
	port := Port{
		Key:         "AEAJM",
		Name:        "",
		City:        "",
		Country:     "United Arab Emirates",
		Alias:       Alias{},
		Regions:     Regions{},
		Coordinates: Coordinates{},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      Unlocs{"AEAJM"},
		Code:        "52000",
	}

	timezone, err := port.GetTimezoneLocation()
	assert.Equal(t, nil, err)
	assert.Equal(t, "Asia/Dubai", timezone.String())
}

func TestPort_GetChanges(t *testing.T) {
	port1 := &Port{
		Key:         "AEAJM",
		Name:        "",
		City:        "",
		Country:     "United Arab Emirates",
		Alias:       Alias{},
		Regions:     Regions{},
		Coordinates: Coordinates{},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      Unlocs{"AEAJM"},
		Code:        "52000",
	}
	port2 := &Port{
		Key:     "AEAJM",
		Name:    "Ajman",
		City:    "Ajman",
		Country: "United Arab Emirates",
		Alias:   Alias{},
		Regions: Regions{},
		Coordinates: Coordinates{
			55.5136433,
			25.4052165,
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs:   Unlocs{"AEAJM"},
		Code:     "52000",
	}

	changes := map[string]interface{}{
		"city":        "Ajman",
		"coordinates": Coordinates{55.513645, 25.405216},
		"name":        "Ajman",
	}
	assert.Equal(t, changes, port1.GetChanges(port2))
}
