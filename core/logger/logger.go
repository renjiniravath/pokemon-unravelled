package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

var (
	Info    *logrus.Logger
	Error   *logrus.Logger
	Success *logrus.Logger
)

func LogData(f io.Writer) {
	Info = logrus.New()
	Error = logrus.New()
	Success = logrus.New()

	Info.SetReportCaller(true)
	Error.SetReportCaller(true)
	Success.SetReportCaller(true)

	// Log.SetFormatter()
	Info.SetOutput(f)
	Error.SetOutput(f)
	Success.SetOutput(f)
}
