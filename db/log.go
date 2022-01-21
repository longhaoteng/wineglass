package db

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/consts"
	"github.com/longhaoteng/wineglass/logger"
)

type dbLog struct {
	entry                 *logrus.Entry
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func newLog() *dbLog {
	return &dbLog{
		entry:                 logger.GetEntry(),
		SlowThreshold:         time.Duration(config.DB.LowThreshold) * time.Millisecond,
		SourceField:           consts.SourceField,
		SkipErrRecordNotFound: true,
	}
}

func (l *dbLog) LogMode(gormlog.LogLevel) gormlog.Interface {
	return l
}

func (l *dbLog) Info(ctx context.Context, s string, args ...interface{}) {
	l.entry.WithContext(ctx).Infof(s, args)
}

func (l *dbLog) Warn(ctx context.Context, s string, args ...interface{}) {
	l.entry.WithContext(ctx).Warnf(s, args)
}

func (l *dbLog) Error(ctx context.Context, s string, args ...interface{}) {
	l.entry.WithContext(ctx).Errorf(s, args)
}

func (l *dbLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	level := logrus.DebugLevel
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[consts.ErrKey] = err
		level = logrus.ErrorLevel
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		level = logrus.WarnLevel
	}

	l.entry.WithContext(ctx).WithFields(fields).Logf(level, "%s [%s]", sql, elapsed)
}
