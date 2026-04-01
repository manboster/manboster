package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var conf Config

// Init reads your configuration from $PWD/config.yml
func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		// if no args, execute guide configuration.
		if len(os.Args) == 1 {
			fmt.Println("config.yaml is not found, now guide you to create one...")
			conf, err := Form()
			if err != nil {
				panic(err)
			}
			err = Write(conf)
			if err != nil {
				panic(err)
			}
			fmt.Println("Successfully created config.yaml, open Manboster again and enjoy it!")
			os.Exit(0)
		}
	}
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
}

// Read provides a whole data using viper to application.
func Read() Config {
	return conf
}
