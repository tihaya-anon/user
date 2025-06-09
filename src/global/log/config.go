package log

import (
	"MVC_DI/global"
	"runtime"

	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
)

var (
	timeColor   = global.Green
	callerColor = global.Grey
)

var levelColorMap = map[logrus.Level]*color.RGBStyle{
	logrus.PanicLevel: global.Red,
	logrus.FatalLevel: global.Red,
	logrus.ErrorLevel: global.Red,
	logrus.WarnLevel:  global.Yellow,
	logrus.InfoLevel:  global.Blue,
	logrus.DebugLevel: global.Grey,
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
