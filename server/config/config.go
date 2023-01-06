package config

import (
	"flag"
)

type AppConfig struct {
	Port      int
	SecretKey string
}

var config AppConfig

const defaultPort = 5678
const defaultSecretKey = "superstrongpassword"

func Init() {
	// NOTE: For demo purposes, the config is either hard-coded, or provided through CLI.
	// In a real scenario, there would be some kind of environment file parsing here instead.
	config.Port = *flag.Int("port", defaultPort, "")
	config.SecretKey = *flag.String("secretkey", defaultSecretKey, "")

	flag.Parse()
}

func GetConfig() *AppConfig {
	return &config
}
