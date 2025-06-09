package log

import (
	"runtime"

	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
)

var (
	timeColor   = color.HEX("#018025")
	callerColor = color.HEX("#888888")
)

var levelColorMap = map[logrus.Level]*color.RGBStyle{
	logrus.PanicLevel: color.HEXStyle("#d32f2f"),
	logrus.FatalLevel: color.HEXStyle("#d32f2f"),
	logrus.ErrorLevel: color.HEXStyle("#d32f2f"),
	logrus.WarnLevel:  color.HEXStyle("#fbc02d"),
	logrus.InfoLevel:  color.HEXStyle("#2196f3"),
	logrus.DebugLevel: color.HEXStyle("#777777"),
}
var levelNameMap = map[logrus.Level]string{
	logrus.PanicLevel: "PANC",
	logrus.FatalLevel: "FATL",
	logrus.ErrorLevel: "ERRO",
	logrus.WarnLevel:  "WARN",
	logrus.InfoLevel:  "INFO",
	logrus.DebugLevel: "DEBU",
}

type StackTraceHook struct{}

func (hook *StackTraceHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
}

func (hook *StackTraceHook) Fire(entry *logrus.Entry) error {
	entry.Data["stack"] = getStackTrace()
	return nil
}

func getStackTrace() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
