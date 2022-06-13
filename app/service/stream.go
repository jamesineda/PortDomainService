package service

import (
	"encoding/json"
	"github.com/jamesineda/PortDomainService/app/db"
	"github.com/jamesineda/PortDomainService/app/models"
	"log"
	"os"
	"time"
)

type Service struct {
	Database db.Client
}

func NewService(client db.Client) *Service {
	return &Service{Database: client}
}

func (s *Service) Start(filepath string, stopChannel <-chan bool, finishedChannel, count chan<- bool) {

	go func(stop <-chan bool, finish, count chan<- bool) {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatalf("Error reading [file=%v]: %v", filepath, err.Error())
		}

		defer func() {
			log.Println("Closing file reader")
			file.Close()
		}()

		if decodeErr := s.DecodePortStream(file, stop, finish, count); decodeErr != nil {
			log.Printf("JSON file process stopped with an error: %s", decodeErr)
		} else {
			log.Println("JSON file processing finished, exiting service")
		}

		finish <- true

	}(stopChannel, finishedChannel, count)
}

func (s *Service) DecodePortStream(file *os.File, stop <-chan bool, finish, count chan<- bool) error {
	decoder := json.NewDecoder(file)
	i := 0

	_, initTokenErr := decoder.Token()
	if initTokenErr != nil {
		return initTokenErr
	}

	ports := make([]*models.Port, 0)

	log.Println("Reading port stream and processing data (this may take a while)")
	for decoder.More() {
		select {
		case <-stop:
			log.Println("Stop command received, shutting down JSON stream")
			finish <- true
			return nil

		default:

			key, tokenErr := decoder.Token()
			if tokenErr != nil {
				return tokenErr
			}

			port := &models.Port{Key: key.(string)}
			decodeErr := decoder.Decode(&port)
			if decodeErr != nil {
				return decodeErr
			}

			// if the process cannot commit to database, exit the program
			if persistenceErr := s.CreateOrUpdatePort(port); persistenceErr != nil {
				return persistenceErr
			}

			ports = append(ports, port)
			count <- true
			i++
		}

	}
	decoder.Token()

	return nil
}

func (s *Service) CreateOrUpdatePort(port *models.Port) error {
	dbPort := s.Database.Get(port.Key, "port")
	if dbPort != nil {
		existingPort := dbPort.(*models.Port)
		changes := existingPort.GetChanges(port)
		if len(changes) > 0 {
			if err := s.Database.Update(port.Key, existingPort, changes); err != nil {
				return err
			}
		}
	} else {
		if err := s.Database.Create(port.Key, port); err != nil {
			return err
		}
	}

	// for simulation purposes, as this runs incredibly quickly
	time.Sleep(15 * time.Millisecond)
	return nil
}
