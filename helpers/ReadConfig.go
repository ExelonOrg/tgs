package helpers

import (
	"log"

	"github.com/davoodharun/terragrunt-scaffolder/structs"
	"github.com/spf13/viper"
)

func ReadConfig() structs.Config {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".tgs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var config structs.Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to form into struct, %s", err)
	}

	return config

}
