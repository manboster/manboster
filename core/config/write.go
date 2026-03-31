package config

import (
	"encoding/json"

	"github.com/spf13/viper"
)

// Write writes config into a yml file.
func Write(conf Config) error {
	// dull jobs: get the correct map to feed viper's configuration.
	jsonData, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	var cMap map[string]any
	if err := json.Unmarshal(jsonData, &cMap); err != nil {
		return err
	}

	// write configuration
	if err := viper.MergeConfigMap(cMap); err != nil {
		return err
	}
	viper.SetConfigName("config.yaml")
	return viper.WriteConfigAs("config.yaml")
}
