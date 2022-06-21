package main

import (
	"github.com/h-ft/gort-bot/config"
	"github.com/h-ft/gort-bot/discord"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.New("config/config.yaml")
	if err != nil {
		logrus.Fatalf("[main] Error initializing config: %v", err)
	}

	if err = discord.Init(cfg); err != nil {
		logrus.Fatalf("[main] Error initializing discord connection: %v", err)
	}
}
