package config

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Environment string
type ProtocolType string

const (
	Dev  Environment = "dev"
	Prod Environment = "prod"

	TCP ProtocolType = "tcp"
	UDP ProtocolType = "udp"
)

var Protocols = []ProtocolType{TCP, UDP}

type Protocol struct {
	Type ProtocolType `yaml:"type"`
	Port string       `yaml:"port"`
}

type Config struct {
	Env      Environment `yaml:"env"`
	Protocol Protocol    `yaml:"protocol"`
}

func (c *Config) checkConfig() {
	switch c.Env {
	case Dev, Prod:
		log.Printf("Running in %s environment\n", c.Env)
	default:
		panic(fmt.Sprintf("Invalid environment: %s, Must be one of %v", c.Env, []Environment{Dev, Prod}))
	}

	if !slices.Contains(Protocols, c.Protocol.Type) {
		panic(fmt.Sprintf("Invalid protocol type: %s\nMust be one of %v", c.Protocol.Type, Protocols))
	}
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		panic("CONFIG_PATH environment variable is not set")
	}
	var config Config
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		panic(err)
	}
	config.checkConfig()
	return &config
}
