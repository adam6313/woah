package logger

import (
	"woah/config"

	"github.com/tyr-tech-team/hawk/log"
	"go.uber.org/zap"
)

// NewLogger -
// TODO - result logger interface
func NewLogger(ic config.IConfig) *zap.Logger {
	return log.NewZapLogger(ic.Log())
}
