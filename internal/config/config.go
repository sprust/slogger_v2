package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	mutex sync.Mutex
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
	})

	return instance
}

func (c *Config) GetLogDir() string {
	return os.Getenv("LOG_DIR")
}

func (c *Config) GetLogLevels() string {
	return os.Getenv("LOG_LEVELS")
}

func (c *Config) GetLogKeepDays() int {
	keepDays, err := strconv.Atoi(os.Getenv("LOG_KEEP_DAYS"))

	if err != nil {
		fmt.Println("LOG_KEEP_DAYS is not set or invalid, using default value 3")

		keepDays = 3
	}

	return keepDays
}

func (c *Config) GetGrpcPort() string {
	return os.Getenv("GRPC_PORT")
}
