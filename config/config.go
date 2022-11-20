package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type appConfig struct {
	MongoURI  string `mapstructure:mongo_uri`
	DBName    string `mapstructure:db_name`
	JWTSecret string `mapstructure:jwt_secret`
}

var (
	AppConfig *appConfig
)

// LoadConfig load config from config.yaml
func LoadConfig(configPath string) error {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AutomaticEnv()

	if configPath != "" {
		v.AddConfigPath(configPath)
	} else {
		return fmt.Errorf("failed to load config file. please set config")
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	if err := v.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("failed to unmarshall app config: %v", err)
	}

	return nil
}
