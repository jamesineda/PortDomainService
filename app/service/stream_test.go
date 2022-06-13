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

// AEAJM is created and then updated to the most recent version from the test data above
// AEJEA is created
func (suite *StreamServiceTestSuite) TestService_CreateOrUpdatePort() {
	service := NewService(suite.FakeDatabase)

	for _, port := range suite.TestPorts {
		suite.Equal(nil, service.CreateOrUpdatePort(port))
	}

	suite.Equal(suite.UpdatedAjmanPort, suite.FakeDatabase.Get("AEAJM", "port"))
	suite.Equal(suite.JebelPort, suite.FakeDatabase.Get("AEJEA", "port"))
	suite.Equal(2, suite.FakeDatabase.Count("port"))
}

func TestStreamServiceTestSuite(t *testing.T) {
	suite.Run(t, new(StreamServiceTestSuite))
}
