package config

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

// Write writes config into a yml file.
func Write(conf Config) error {
	var m map[string]any

	if err := mapstructure.Decode(conf, &m); err != nil {
		return err
	}

	if err := viper.MergeConfigMap(m); err != nil {
		return err
	}

	viper.SetConfigName("config.yaml")
	return viper.WriteConfigAs("config.yaml")
}
