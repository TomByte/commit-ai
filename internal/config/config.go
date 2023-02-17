package config

import (
	"errors"
	"os"
)

type Config struct {
	OpenAIKey string
	OpenAIUrl string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	if os.Getenv("OPENAI_KEY") == "" {
		return errors.New("OPENAI_KEY environment variable not set")
	}

	c.OpenAIKey = os.Getenv("OPENAI_KEY")
	c.OpenAIUrl = "https://api.openai.com/v1"
	return nil
}
