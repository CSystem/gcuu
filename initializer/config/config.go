package config

import (
	"log"
	"github.com/spf13/viper"
)

// Init - Initialization config options.
// Read the config file from the current directory and marshal into the conf config struct.
func Init() *Configuration {
	// ... leave this block untouched...
	viper.AddConfigPath("./configs")
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	conf := &Configuration{}
	err := viper.Unmarshal(conf)

	if err != nil {
		log.Fatalf("unable to decode into config struct, %v", err)
	}

	return conf
}
