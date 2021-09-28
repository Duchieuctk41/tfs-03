package util

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	DBDriver string `mapstructure:"DBDriver"`
	DSN      string `mapstructure:"DSN"`
}

func LoadDBConfig(path string) (config DBConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
