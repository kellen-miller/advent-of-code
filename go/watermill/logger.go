package watermill

import (
	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/zap"
)

// Logger implements watermill.LoggerAdapter with *zap.Logger.
type Logger struct {
	base      *zap.Logger
	logFields watermill.LogFields
}

// NewLogger returns new watermill.LoggerAdapter using passed *zap.Logger as base.
func NewLogger(z *zap.Logger) watermill.LoggerAdapter {
	return &Logger{base: z}
}

// Error writes error log with message, error and some logFields.
func (l *Logger) Error(msg string, err error, wmFields watermill.LogFields) {
	zapFields := l.convFields(wmFields)
	zapFields = append(zapFields, zap.Error(err))

	l.base.Error(msg, zapFields...)
}

// Info writes info log with message and some logFields.
func (l *Logger) Info(msg string, fields watermill.LogFields) {
	l.base.Info(msg, l.convFields(fields)...)
}

// Debug writes debug log with message and some logFields.
func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	l.base.Debug(msg, l.convFields(fields)...)
}

// Trace writes debug log instead of trace log because zap does not support trace level logging.
func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	l.base.Debug(msg, l.convFields(fields)...)
}

// With returns new LoggerAdapter with passed logFields.
func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Logger{
		base:      l.base,
		logFields: l.logFields.Add(fields),
	}
}

func (l *Logger) convFields(logFields watermill.LogFields) []zap.Field {
	zfs := make([]zap.Field, len(l.logFields)+len(logFields))
	i := 0

	for k, v := range l.logFields { // add base fields first
		zfs[i] = zap.Any(k, v)
		i++
	}

	for k, v := range logFields { // add passed fields
		zfs[i] = zap.Any(k, v)
		i++
	}

	return zfs
}
