package db

import (
	"github.com/jamesineda/PortDomainService/app/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FakeDatabaseTestSuite struct {
	suite.Suite
	DubaiPort    *models.Port
	FakeDatabase *FakeDatabase
}

func (suite *FakeDatabaseTestSuite) SetupTest() {
	suite.DubaiPort = &models.Port{
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

	suite.FakeDatabase = &FakeDatabase{
		database: map[string]map[string]interface{}{"port": {"AEAJM": suite.DubaiPort}},
	}
}

func (suite *FakeDatabaseTestSuite) TestFakeDatabase_Create() {
	newPort := &models.Port{
		Key:     "AERKT",
		Name:    "Ras al Khaimah",
		City:    "Ras al Khaimah",
		Country: "United Arab Emirates",
		Alias:   models.Alias{},
		Regions: models.Regions{},
		Coordinates: models.Coordinates{
			55.95,
			25.78,
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs:   models.Unlocs{"AERKT"},
		Code:     "52000",
	}

	err := suite.FakeDatabase.Create("AERKT", newPort)
	suite.Equal(err, nil)
}

func (suite *FakeDatabaseTestSuite) TestFakeDatabase_Get() {
	portFromDb := suite.FakeDatabase.Get("AEAJM", "port")
	suite.Equal(portFromDb, suite.DubaiPort)
}

func (suite *FakeDatabaseTestSuite) TestFakeDatabase_Update() {
	changes := map[string]interface{}{
		"name": "Ajman",
		"city": "Ajman",
		"coordinates": models.Coordinates{
			55.5136433,
			25.4052165,
		},
		"timezone": "Europe/London",
	}

	portFromDb := suite.FakeDatabase.Get("AEAJM", "port")
	portObject, _ := portFromDb.(*models.Port)
	err := suite.FakeDatabase.Update("AEAJM", portObject, changes)
	suite.Equal(err, nil)
	suite.Equal("Ajman", portObject.Name)
	suite.Equal("Ajman", portObject.City)
	suite.Equal(models.Coordinates{
		55.5136433,
		25.4052165,
	}, portObject.Coordinates)
	suite.Equal("Europe/London", portObject.Timezone)
}

func (suite *FakeDatabaseTestSuite) TestFakeDatabase_Delete() {
	portFromDb := suite.FakeDatabase.Get("AEAJM", "port")
	portObject, _ := portFromDb.(*models.Port)

	err := suite.FakeDatabase.Delete(portObject.Key, "port")
	suite.Equal(err, nil)
	suite.Equal(nil, suite.FakeDatabase.Get("AEAJM", "port"))
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(FakeDatabaseTestSuite))
}
