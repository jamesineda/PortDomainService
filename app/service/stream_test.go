package service

import (
	"github.com/jamesineda/PortDomainService/app/db"
	"github.com/jamesineda/PortDomainService/app/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StreamServiceTestSuite struct {
	suite.Suite
	AjmanPort        *models.Port
	UpdatedAjmanPort *models.Port
	JebelPort        *models.Port

	TestPorts    []*models.Port
	FakeDatabase *db.FakeDatabase
}

func (suite *StreamServiceTestSuite) SetupTest() {
	suite.AjmanPort = &models.Port{
		Key:         "AEAJM",
		Name:        "",
		City:        "",
		Country:     "United Arab Emirates",
		Alias:       models.Alias{},
		Regions:     models.Regions{},
		Coordinates: models.Coordinates{},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      models.Unlocs{"AEAJM"},
		Code:        "52000",
	}

	suite.UpdatedAjmanPort = &models.Port{
		Key:     "AEAJM",
		Name:    "Ajman",
		City:    "Ajman",
		Country: "United Arab Emirates",
		Alias:   models.Alias{},
		Regions: models.Regions{},
		Coordinates: models.Coordinates{
			55.5136433,
			25.4052165,
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs:   models.Unlocs{"AEAJM"},
		Code:     "52000",
	}

	suite.JebelPort = &models.Port{
		Key:     "AEJEA",
		Name:    "Jebel Ali",
		City:    "Jebel Ali",
		Country: "United Arab Emirates",
		Alias:   models.Alias{},
		Regions: models.Regions{},
		Coordinates: models.Coordinates{
			55.0272904,
			24.9857145,
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs:   models.Unlocs{"AEJEA"},
		Code:     "52051",
	}

	suite.TestPorts = []*models.Port{
		suite.AjmanPort,
		suite.UpdatedAjmanPort,
		suite.JebelPort,
	}

	suite.FakeDatabase = db.NewFakeDatabase()
}

func (suite *StreamServiceTestSuite) TestService_CreateOrUpdatePort() {
	service := NewService(suite.FakeDatabase)

	for _, port := range suite.TestPorts {
		suite.Equal(nil, service.CreateOrUpdatePort(port))
	}

	suite.Equal(suite.FakeDatabase.Get("AEAJM", "port"), suite.UpdatedAjmanPort)
	suite.Equal(suite.FakeDatabase.Get("AEJEA", "port"), suite.JebelPort)
}

func TestStreamServiceTestSuite(t *testing.T) {
	suite.Run(t, new(StreamServiceTestSuite))
}
