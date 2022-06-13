package main

import (
	"flag"
	"github.com/jamesineda/PortDomainService/app/db"
	"github.com/jamesineda/PortDomainService/app/service"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	FILEPATH = "FILEPATH"
)

func BindCommandLineArgs() {
	flag.String(FILEPATH, "ports.json", "path to ports file")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {

	BindCommandLineArgs()
	filepath := viper.GetString(FILEPATH)

	log.Println("Starting Port Domain Service")
	log.Printf("Reading file from specified location: %s", filepath)

	stopChannel := make(chan bool, 1)
	finishedChannel := make(chan bool, 1)
	processCountChannel := make(chan bool, 0)
	processCount := 0
	sigC := make(chan os.Signal)

	streamService := service.NewService(db.NewFakeDatabase())
	streamService.Start(filepath, stopChannel, finishedChannel, processCountChannel)

	// go routine that listens for Interrupt/ Terminate signal which notifies the decoder routine
	go func() {
		log.Println("Ctrl + C to QUIT")
		signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
		<-sigC
		signal.Stop(sigC)

		stopChannel <- true
	}()

	// wait to receive flag down finishedChannel before exiting program
	for {
		select {
		case <-processCountChannel:
			processCount += 1
		case <-finishedChannel:
			log.Printf("Processed %d port objects", processCount)
			log.Println("Shutting down service")
			time.Sleep(10 * time.Second)
			return
		}
	}
}
