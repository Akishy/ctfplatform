package models

import "go.uber.org/zap"

type FlagGenerator struct {
	Logger     *zap.Logger
	FlagFormat string
}
