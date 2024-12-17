package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var FxOption fx.Option = fx.Options(fx.Provide(NewLogger), fx.Invoke(func(*zap.Logger) {}))

func NewLogger() (*zap.Logger, error) {
	loggerOutput, err := setLoggerOutput()
	loggerType := setLoggerType()
	if err != nil {
		return nil, err
	}
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(setLoggerLevel()),
		Development: loggerType == "development",

		Encoding:         setLoggerLogType(),
		EncoderConfig:    zapcore.EncoderConfig{},
		OutputPaths:      loggerOutput,
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    nil,
	}
	if loggerType == "development" {
		config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.EncoderConfig = zap.NewProductionEncoderConfig()
	}
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	return logger, nil
}

func setLoggerType() string {
	return viper.GetString("loggerType")
}

func setLoggerLevel() zapcore.Level {
	loggerLevel := viper.GetString("loggerLogLevel")
	switch loggerLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func setLoggerOutput() ([]string, error) {
	loggerWriter := viper.GetString("loggerWriter")
	loggerOutputSources := strings.Split(loggerWriter, ";")
	for i, source := range loggerOutputSources {
		strings.TrimSpace(source)
		if source == "" {
			loggerOutputSources = append(loggerOutputSources[:i], loggerOutputSources[i+1:]...)
		}
		if source != "stdout" && source != "stderr" && source != "stdin" {
			if _, err := os.Stat(source); os.IsNotExist(err) {
				_, createErr := os.Create(source)
				if createErr != nil {
					return nil, err
				}
			}
		}
	}
	return loggerOutputSources, nil
}

func setLoggerLogType() string {
	return viper.GetString("loggerLogType")
}
