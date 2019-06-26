package logger

import (
	"net/http"
	"os"

	"go.elastic.co/apm/module/apmlogrus"

	"github.com/sirupsen/logrus"
)

func NewMyLogger(req *http.Request) *logrus.Entry {
	return Log.WithFields(apmlogrus.TraceContext(req.Context()))
}

var Log = &logrus.Logger{
	Out:   os.Stderr,
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
	Formatter: &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "function.name", //non-ECS
		},
	},
}
