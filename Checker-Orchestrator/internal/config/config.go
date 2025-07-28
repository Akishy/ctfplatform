package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AllowEmptyEnv(false)
	viper.SetEnvPrefix(viper.GetString("serviceName")) // serviceName will be uppercased automatically

	setDefaults()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./build/")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Printf("Config: config file not found: %s\n", err.Error())
		} else {
			fmt.Printf("Config: read config file error: %s\n", err.Error())
		}
	}
}

func setDefaults() {
	viper.SetDefault("serviceName", "Checker-Orchestrator")
	err := viper.BindEnv("serviceName", "SERVICE_NAME")
	if err != nil {
		fmt.Printf("Config: Error binding env var: %s\n", err.Error())
	}

	viper.SetDefault("loggerType", "production")
	err = viper.BindEnv("loggerType", "LOGGER_TYPE")
	if err != nil {
		fmt.Printf("Config: Error binding env var: %s\n", err.Error())
	}

	viper.SetDefault("loggerLogLevel", "info")
	err = viper.BindEnv("loggerLogLevel", "LOGGER_LOG_LEVEL")
	if err != nil {
		fmt.Printf("Config: Error binding env var: %s\n", err.Error())
	}

	viper.SetDefault("loggerLogType", "json")
	err = viper.BindEnv("loggerLogType", "LOGGER_LOG_TYPE")
	if err != nil {
		fmt.Printf("Config: Error binding env var: %s\n", err.Error())
	}

	viper.SetDefault("loggerWriter", "stdout")
	err = viper.BindEnv("loggerWriter", "LOGGER_WRITER")
	if err != nil {
		fmt.Printf("Config: Error binding env var: %s\n", err.Error())
	}
}
