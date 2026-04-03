package config

import (
	"os"
	"path/filepath"

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

	// if there is a homedir, write into homedir
	homedir, err := os.UserHomeDir()
	if err == nil {
		dir := filepath.Join(homedir, ".manboster")

		// make directory first
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		// write configuration
		return viper.WriteConfigAs(filepath.Join(dir, "config.yaml"))
	}

	return viper.WriteConfigAs("config.yaml")
}
