package logger

import (
	"context"
	"fmt"
	formatter "github.com/antonfisher/nested-logrus-formatter"
	thirdLogger "github.com/sirupsen/logrus"
	"time"
)

var log = thirdLogger.New()

func init() {
	log.SetLevel(6)
	matter := formatter.Formatter{
		CallerFirst:     true,
		TimestampFormat: time.RFC3339,
	}
	log.SetFormatter(&matter)
}

func Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.Warnf(msg, keysAndValues...)
}

func ErrorWithErr(ctx context.Context, msg string, err error) {
	errStr := fmt.Sprintf("%+v\n", err)
	log.WithContext(ctx).WithError(err).WithField("err", errStr).Error(msg)
}

func Err(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.WithContext(ctx).Errorf(msg, keysAndValues...)
}

func Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.WithContext(ctx).Errorf(msg, keysAndValues...)
}

func Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.Infof(msg, keysAndValues...)
}

func Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.Debugf(msg, keysAndValues...)
}
