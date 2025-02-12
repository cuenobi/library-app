package configs

import (
	"embed"
	"errors"
	"io/fs"
	"sync"

	"library-service/internal/adapter/postgres"

	"github.com/spf13/viper"
)

// Embed config files
//
//go:embed config.yaml
var configs embed.FS

var (
	once     = &sync.Once{}
	instance *Config
)

type Config struct {
	Postgres     *postgres.PostgresConfig
	ServerConfig *ServerConfig
}

type ServerConfig struct {
	Port string
}

func loadConfig() *Config {
	viper.SetConfigType("yaml")

	// Load base config
	baseConfig, err := configs.Open("config.yaml")
	if err != nil {
		panic(err.Error())
	}
	err = viper.ReadConfig(baseConfig)
	if err != nil {
		panic(err.Error())
	}

	// Load dev config (optional)
	devConfig, err := configs.Open("config-dev.yaml")
	if err == nil {
		err = viper.MergeConfig(devConfig)
		if err != nil {
			panic(err.Error())
		}
	} else if !errors.Is(err, fs.ErrNotExist) {
		panic(err.Error())
	}

	viper.AutomaticEnv()

	// Server config
	ServerConfig := &ServerConfig{
		Port: viper.GetString("SERVER.PORT"),
	}

	// Postgres config
	PostGresConfig := &postgres.PostgresConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Name:     viper.GetString("DB_NAME"),
		Username: viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
	}

	return &Config{
		Postgres:     PostGresConfig,
		ServerConfig: ServerConfig,
	}
}

func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}
