package infrastructure

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Infrastructure struct {
	Config *Config
	Router *echo.Echo
}

func NewInfrastructure() *Infrastructure {
	return new(Infrastructure)
}

func (i *Infrastructure) LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config not found")
		}
		return err
	}
	newConfig := Config{}
	err = viper.Unmarshal(&newConfig)
	if err != nil {
		return err
	}
	i.Config = &newConfig

	return nil
}
