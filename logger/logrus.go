package logger

import (
	"os"

	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/consts/timef"
	"github.com/sirupsen/logrus"
)

var (
	log = &Logrus{hooks: make(logrus.LevelHooks)}
)

type Logrus struct {
	hooks logrus.LevelHooks
	entry *logrus.Entry
}

func (l *Logrus) init() error {
	level, err := logrus.ParseLevel(config.Log.Level)
	if err != nil {
		return err
	}

	std := logrus.New()

	std.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timef.YearMonthDayHourMinuteSecond,
	})
	if config.IsDevEnv() {
		std.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: timef.YearMonthDayHourMinuteSecond,
		})
	}

	std.SetOutput(os.Stdout)
	std.SetLevel(level)
	std.ReplaceHooks(log.hooks)

	l.entry = logrus.NewEntry(std)

	return nil
}

func GetEntry() *logrus.Entry {
	return log.entry
}

func AddHook(hook logrus.Hook) {
	for _, level := range hook.Levels() {
		log.hooks[level] = append(log.hooks[level], hook)
	}
}
