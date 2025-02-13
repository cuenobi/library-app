package configs

import (
	"embed"
	"errors"
	"io/fs"
	"os"
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
	JwtConfig    *JwtConfig
}

type ServerConfig struct {
	Port string
}

type JwtConfig struct {
	Secret string
	Exp    int
}

// loadConfig loads the configuration from the config files.
func loadConfig() *Config {
	viper.SetConfigType("yaml")

	isDocker := os.Getenv("RUNNING_IN_DOCKER") == "true"

	// Load the base config
	baseConfig, err := configs.Open("config.yaml")
	if err != nil {
		panic(err.Error())
	}
	err = viper.ReadConfig(baseConfig)
	if err != nil {
		panic(err.Error())
	}

	// If the config-dev.yaml file exists, merge it with the base config
	devConfig, err := configs.Open("config-dev.yaml")
	if err == nil {
		err = viper.MergeConfig(devConfig)
		if err != nil {
			panic(err.Error())
		}
	} else if !errors.Is(err, fs.ErrNotExist) {
		panic(err.Error())
	}

	// Enable automatic environment variable loading
	viper.AutomaticEnv()

	// Server config
	ServerConfig := &ServerConfig{
		Port: viper.GetString("SERVER.PORT"),
	}

	// Postgres config
	var PostGresConfig *postgres.PostgresConfig
	if isDocker {
		PostGresConfig = &postgres.PostgresConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			Name:     viper.GetString("DB_NAME"),
			Username: viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
		}
	} else {
		PostGresConfig = &postgres.PostgresConfig{
			Host:     viper.GetString("CONNECTION.POSTGRES.HOST"),
			Port:     viper.GetString("CONNECTION.POSTGRES.PORT"),
			Name:     viper.GetString("CONNECTION.POSTGRES.NAME"),
			Username: viper.GetString("CONNECTION.POSTGRES.USERNAME"),
			Password: viper.GetString("CONNECTION.POSTGRES.PASSWORD"),
		}
	}

	// JWT config
	JwtConfig := &JwtConfig{
		Secret: viper.GetString("JWT.SECRET_KEY"),
		Exp:    viper.GetInt("JWT.EXP"),
	}

	return &Config{
		Postgres:     PostGresConfig,
		ServerConfig: ServerConfig,
		JwtConfig:    JwtConfig,
	}
}

func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}
