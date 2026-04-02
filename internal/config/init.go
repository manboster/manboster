package config

import (
	"errors"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var conf Config

// Init reads your configuration from $PWD/config.yml
func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		// if no args, execute guide configuration.
		if errors.As(err, &viper.ConfigFileNotFoundError{}) && len(os.Args) == 1 {
			color.Yellow("config.yaml is not found, now guide you to create one...\n")
			return ErrNoConfig
		}
		return err
	}
	if err := viper.Unmarshal(&conf); err != nil {
		return err
	}
	return nil
}

// Read provides a whole data using viper to application.
func Read() Config {
	return conf
}
