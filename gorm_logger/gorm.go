package gorm_logger

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	logger *logrus.Logger
}

var _ logger.Writer = &GormLogger{}

func (g *GormLogger) Printf(format string, args ...interface{}) {
	g.logger.WithField("level", "Database").Debugf(format, args...)
}

func NewGormLogger(logger *logrus.Logger) *GormLogger {
	return &GormLogger{
		logger: logger,
	}
}
