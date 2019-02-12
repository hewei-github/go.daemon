package main

import (
	"fmt"
	"os"
	"time"

	service "github.com/hewei-github/godaemon"
)

var log service.Logger

func main() {
	var name = "GoServiceTest"
	var displayName = "Go Service Test"
	var desc = "This is a test Go service.  It is designed to run well."

	var s, err = service.NewService(name, displayName, desc)
	log = s

	if err != nil {
		fmt.Printf("%s unable to start: %s", displayName, err)
		return
	}

	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			err = s.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" installed.\n", displayName)
		case "remove":
			err = s.Remove()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" removed.\n", displayName)
		case "run":
			doWork()
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" started.\n", displayName)
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" stopped.\n", displayName)
		}
		return
	}
	// start control
	startFunc = func() error {
		go doWork()
		return nil
	}
	// stop control
	stopFunc = func() error {
		stopWork()
		return nil
	}
	err = s.Run(startFunc, stopFunc)
	if err != nil {
		s.Error(err.Error())
	}
}

var exit = make(chan struct{})

func doWork() {
	log.Info("I'm Running!")
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			log.Info("Still running...")
		case <-exit:
			ticker.Stop()
			return
		}
	}
}
func stopWork() {
	log.Info("I'm Stopping!")
	exit <- struct{}{}
}
