package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is a package level variable, every program should acces logging through "Log"
var Log Logger

// Logger is an interface to represent the logging fucntions
type Logger interface {
	ErrorF(format string, args ...interface{})
	InfoF(format string, args ...interface{})
	WarnF(format string, args ...interface{})
}

// InitLogger should be the only way to a init the package level variable
func InitLogger() {

	cfg := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var zLogger *zap.Logger
	var err error
	zLogger, err = cfg.Build()

	if err != nil {
		panic(err)
	}

	defer zLogger.Sync()
	zSugarLog := zLogger.Sugar()

	Log = &zapLogger{log: zSugarLog}
}
