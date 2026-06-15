package config

import (
	"errors"

	"github.com/spf13/viper"
)

var conf Config

// Init reads your configuration from $PWD/config.yml
func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// first, we check whether there is ~/.manboster/ or not
	viper.AddConfigPath(Path(""))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return ErrNoConfig
	} else if err != nil {
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
