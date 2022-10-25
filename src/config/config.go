package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	DBSource string `mapstructure:"DB_SOURCE"`
	DBDriver string `mapstructure:"DB_DRIVER"`
}

// LoadConfig loads config from env
func LoadConfig(path string, env string) (AppConfig, error) {
	var config AppConfig
	configPath := fmt.Sprintf("../%s/env", path)
	viper.AddConfigPath(configPath)
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}

// LoadEnvConfig load app configuration based on file
func LoadEnvConfig() AppConfig {
	cfg, err := LoadConfig("./.", "app")
	if err != nil {
		log.Fatal("Error loading application config: ", err)
	}
	return cfg
}
