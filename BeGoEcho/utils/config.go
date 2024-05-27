package utils

import "github.com/spf13/viper"

type Config struct {
	Enviroment string `mapstructure:ENVIROMENT`
	DbDriver   string `mapstructure:DB_DRIVER`
	DbSource   string `mapstructure:DB_SOURCE`
	LoggerPath string `mapstructure:LOGGER_PATH`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	// viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
