package logger

import (
	"github.com/sirupsen/logrus"

	"github.com/longhaoteng/wineglass/config"
)

const (
	PanicLevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
	TraceLevel = logrus.TraceLevel
)

type Level = logrus.Level

type Logger struct {
	fields logrus.Fields
}

func Init() error {
	if err := log.init(); err != nil {
		return err
	}
	return nil
}

func V(lv Level) bool {
	level, err := logrus.ParseLevel(config.Log.Level)
	if err != nil {
		return true
	}
	return level <= lv
}

func Field(key string, value interface{}) *Logger {
	return &Logger{
		fields: logrus.Fields{key: value},
	}
}

func Fields(keysAndValues ...interface{}) *Logger {
	fields := logrus.Fields{}

	if len(keysAndValues) == 0 {
		return &Logger{fields: fields}
	}

	for i := 0; i < len(keysAndValues); {
		key := keysAndValues[i]
		if keyStr, ok := key.(string); ok {
			if i+1 < len(keysAndValues) {
				fields[keyStr] = keysAndValues[i+1]
			} else {
				fields[keyStr] = ""
			}
		}
		i += 2
	}

	return &Logger{fields: fields}
}

func Log(level Level, v ...interface{}) {
	log.entry.Log(level, v...)
}

func Trace(args ...interface{}) {
	log.entry.Trace(args...)
}

func Debug(args ...interface{}) {
	log.entry.Debug(args...)
}

func Info(args ...interface{}) {
	log.entry.Info(args...)
}

func Warn(args ...interface{}) {
	log.entry.Warn(args...)
}

func Error(args ...interface{}) {
	log.entry.Error(args...)
}

func Fatal(args ...interface{}) {
	log.entry.Fatal(args...)
}

func Panic(args ...interface{}) {
	log.entry.Panic(args...)
}

func Logf(level Level, format string, v ...interface{}) {
	log.entry.Logf(level, format, v...)
}

func (l *Logger) Log(level Level, v ...interface{}) {
	log.entry.WithFields(l.fields).Log(level, v...)
}

func (l *Logger) Logf(level Level, format string, v ...interface{}) {
	log.entry.WithFields(l.fields).Logf(level, format, v...)
}
