package config

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func (c Config) NewYt() (*youtube.Service, error) {
	ctx := context.Background()

	return youtube.NewService(ctx, option.WithAPIKey(c.YoutubeKey))
}
