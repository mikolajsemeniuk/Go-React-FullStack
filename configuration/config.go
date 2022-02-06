package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
		message := fmt.Sprintf("\nenv environment variable is not set, setting to: %s\n", env)
		fmt.Println("\033[33m" + message + "\033[0m")
	}

	Config = viper.New()
	Config.SetConfigType("yaml")
	Config.SetConfigName(env)
	Config.AddConfigPath("../configuration/")
	Config.AddConfigPath("configuration/")
	err := Config.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
