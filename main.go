package main

import (
	"github.com/h-ft/gort-bot/config"
	"github.com/h-ft/gort-bot/discord"
	"github.com/sirupsen/logrus"
)

var (
	cfg  = &config.Config{}
	path = "config/config.yaml"
)

func main() {
	err := config.New(path, cfg)
	if err != nil {
		logrus.Fatalf("[main] Error initializing config: %v", err)
	}

	cfg.NewYt()

	if err = discord.Init(cfg); err != nil {
		logrus.Fatalf("[main] Error initializing discord connection: %v", err)
	}
}
