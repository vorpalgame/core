package lib

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var intid bool = false

// TODO Refactor and remove Viper...
// Could add Reset or Reinitialize TBD...
func InitConfigurator() {
	if !intid {
		//Ish is necessary to support unit tests.
		//We can do better later...
		viper.AddConfigPath("./etc")
		viper.AddConfigPath("./samples/etc")
		viper.AddConfigPath("../etc")
		viper.AddConfigPath("../samples/etc")
		viper.AddConfigPath("../../etc")
		viper.AddConfigPath("../../samples/etc")
		viper.SetConfigType("yaml")
	}
	intid = true
}
func LoadConfiguration(fileName string) {
	InitConfigurator() //No-op most of the time.
	loadViper(fileName)

}

func loadViper(fileName string) {

	viper.SetConfigName(fileName)

	// REQUIRED if the config file does not have the extension in the name
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	viper.WatchConfig()

}
