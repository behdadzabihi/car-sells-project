package logging

import (
	"golang-project/src/config"
)

func NewZapLogger(cfg *config.Config) Logger {
	if cfg.Logger.Logger == "zap" {
		return newZapLogger(cfg)
	} else if cfg.Logger.Logger == "zero" {
		return newZeroLogger(cfg)
	}
	panic("Invalid logger type")
}
