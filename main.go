package main

import (
	"github.com/sirupsen/logrus"
	"main.go/config"
)

func main() {
	cfg, err := config.New("config/config.yaml")
	if err != nil {
		logrus.Fatalf("[main] Error initializing config: %v", err)
	}
}
