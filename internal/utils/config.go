package utils

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	DBUser   string `mapstructure:"POSTGRES_USER"`
	DBName   string `mapstructure:"POSTGRES_DB"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return Config{}, errors.New("config File \".env\" not found")
		}
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	return
}
