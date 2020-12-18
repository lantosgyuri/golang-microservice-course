package log

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	log *zap.SugaredLogger
}

func (z *zapLogger) ErrorF(format string, args ...interface{}) {
	z.log.Errorf(format, args)
}

func (z *zapLogger) InfoF(format string, args ...interface{}) {
	z.log.Infof(format, args)
}

func (z *zapLogger) WarnF(format string, args ...interface{}) {
	z.log.Warnf(format, args)
}
