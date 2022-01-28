package config

import "os"

type GemininConfig struct {
	Key    string
	Secret string
}

func (config *GemininConfig) New() *GemininConfig {
	config.Key = os.Getenv("GEMININ_KEY")
	config.Secret = os.Getenv("GEMININ_SECRET")
	return config
}
