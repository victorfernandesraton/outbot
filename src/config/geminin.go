package config

import "os"

type GemininConfig struct {
	Key    string
	Domain string
}

func (config *GemininConfig) New() *GemininConfig {
	config.Key = os.Getenv("GEMININ_KEY")
	config.Key = os.Getenv("GEMININ_DOMAIN")
	return config
}
