package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

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

func addHook(level logrus.Level, hks ...logrus.Hook) {
	log.hooks[level] = hks
}

func caller(skip int) func(f *runtime.Frame) (string, string) {
	return func(f *runtime.Frame) (string, string) {
		_, file, line, ok := runtime.Caller(skip)
		fileline := "unknown"
		if ok {
			filePath := strings.ReplaceAll(file, fmt.Sprintf("%s/pkg/mod/", os.Getenv("GOPATH")), "")
			// 去除路径中版本号
			versionIndex := strings.Index(filePath, "@")
			if versionIndex != -1 {
				subPath := filePath[versionIndex:]
				version := subPath[:strings.Index(subPath, "/")]
				filePath = strings.ReplaceAll(filePath, version, "")
			}
			fileline = fmt.Sprintf("%v:%v", filePath, line)
			if config.IsDevEnv() {
				fileline += "\t"
			}
		}
		return "", fileline
	}
}
