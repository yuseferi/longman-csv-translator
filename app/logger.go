package app

import (
	"fmt"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(level string) (logger *zap.Logger, err error) {

	atom := zap.NewAtomicLevel()
	err = atom.UnmarshalText([]byte(level))
	if err != nil {
		return nil, err
	}
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	return logger, err
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Println(v ...interface{}) {
	l.logger.Warn(fmt.Sprint(v...))
}
