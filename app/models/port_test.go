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
