package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmlogrus"
)

// func NewMyLogger(ctx context.Context) *logrus.Entry {
// 	l := Log.WithFields(apmlogrus.TraceContext(ctx))
// 	return l
// }

func NewLogger(ctx context.Context) *logrus.Entry {
	return log.WithFields(apmlogrus.TraceContext(ctx))
}

var log = &logrus.Logger{
	Out:   os.Stderr,
	Hooks: make(logrus.LevelHooks),
	Level: logrus.InfoLevel,
	Formatter: &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "agro", //non-ECS
		},
	},
}

var Log = &logrus.Logger{
	Out:   os.Stderr,
	Hooks: make(logrus.LevelHooks),
	Level: logrus.InfoLevel,
	Formatter: &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "agro", //non-ECS
		},
	},
}

func init() {
	apm.DefaultTracer.SetLogger(log)
	log.AddHook(&apmlogrus.Hook{})
}
