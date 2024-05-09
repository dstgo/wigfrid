package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"log/slog"
	"slices"
)

type SlogAdapter struct {
	logger *slog.Logger
}

func (l *SlogAdapter) Log(level log.Level, keyvals ...interface{}) error {

	// find out the msg attribute and remove it from kvs
	var msg string
	for i := 0; i < len(keyvals); i++ {
		if keyvals[i] == "msg" {
			msg = keyvals[i+1].(string)
			keyvals = slices.Delete(keyvals, i, i+2)
			break
		}
	}

	switch level {
	case log.LevelDebug:
		l.logger.Debug(msg, keyvals...)
	case log.LevelInfo:
		l.logger.Info(msg, keyvals...)
	case log.LevelWarn:
		l.logger.Warn(msg, keyvals...)
	case log.LevelError, log.LevelFatal:
		l.logger.Error(msg, keyvals...)
	}

	return nil
}
