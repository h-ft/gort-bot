package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		DiscordToken string `yaml:"discord_token"`
		YoutubeKey   string `yaml:"youtube_key"`
	}
)

// Creates a new config instance, reading from yaml file
func New(path string, c *Config) error {

	cfgFile, err := os.Open(path)
	if err != nil {
		logrus.Errorf("[config.Open] Error opening config: %v", err)
		return err
	}
	defer cfgFile.Close()

	// Init new yaml decoder
	decoder := yaml.NewDecoder(cfgFile)

	err = decoder.Decode(c)
	if err != nil {
		logrus.Errorf("[config.Decode] Error decoding config: %v", err)
		return err
	}

	return nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		logrus.Errorf("[validateConfigPath.Stat] Error: %v", err)
		return err
	}
	if s.IsDir() {
		err = fmt.Errorf("'%s' is a directory, not a normal file", path)
		logrus.Errorf("[validateConfigPath] Error: %v", err)
		return err
	}
	return nil
}
