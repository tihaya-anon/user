package log

import (
	"MVC_DI/config"
	"MVC_DI/global/module"
	"MVC_DI/util/stream"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

type ConsoleFormatter struct{}

func (f *ConsoleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	time_ := timeColor.Sprintf("%s", entry.Time.Format("2006-01-02 15:04:05.000"))
	level := levelColorMap[entry.Level].Sprintf("[%s]", levelNameMap[entry.Level])
	message := ""
	caller := callerColor.Sprintf("%v:%d", entry.Caller.File[len(module.GetRoot()):], entry.Caller.Line)
	arrow := levelColorMap[entry.Level].Sprintf(">>")
	stack := entry.Data["stack"]
	data := stream.NewMapStream(entry.Data).
		Filter(func(key string, val any) bool {
			return key != "stack"
		}).
		Map(func(key string, val any) (string, any) {
			return key, fmt.Sprintf("%v", val)
		}).
		ToMap()

	for key, val := range data {
		message = fmt.Sprintf("%s %s=%s", message, key, val)
	}
	message = fmt.Sprintf("%s `%s`", message, entry.Message)
	logLine := fmt.Sprintf("%s %s %s\n\t%s %s", time_, level, caller, arrow, message)

	if stack != nil {
		stack = levelColorMap[entry.Level].Sprintf("%s", entry.Data["stack"])
		logLine = fmt.Sprintf("%s\n%s", logLine, stack)
	}

	return append([]byte(logLine), '\n'), nil
}

type ProdFormatter struct{}

type Json struct {
	Time    int            `json:"time"`
	Level   string         `json:"level"`
	Caller  string         `json:"caller"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

type FileWriteHook struct {
	level    logrus.Level
	interval time.Duration
	current  time.Time
	file     *os.File
}

func NewFileWriteHook(level logrus.Level, interval time.Duration) *FileWriteHook {
	hook := &FileWriteHook{
		level:    level,
		interval: interval,
		current:  time.Now(),
	}
	file, err := hook.createLogFile()
	if err != nil {
		panic(err)
	}
	hook.file = file
	return hook
}

func (hook *FileWriteHook) getFilePath() string {
	timeStr := hook.current.Format("2006-01-02_15-04-05")
	nameStr := "tmp"
	return path.Join(module.GetRoot(), "log", timeStr, nameStr+".log")
}

func (hook *FileWriteHook) createLogFile() (*os.File, error) {
	filePath := hook.getFilePath()
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return file, err
}

func (hook *FileWriteHook) Levels() []logrus.Level {
	levels := stream.NewListStream(logrus.AllLevels).Filter(func(level logrus.Level) bool { return level <= hook.level }).ToList()
	return levels
}

func (hook *FileWriteHook) Fire(entry *logrus.Entry) error {
	since := time.Since(hook.current)
	if since > hook.interval {
		hook.current = time.Now()
		file, err := hook.createLogFile()
		if err != nil {
			return err
		}
		hook.file = file
	}

	jsonData, err := jsonFormat(entry)
	if err != nil {
		return err
	}
	_, err = hook.file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func jsonFormat(entry *logrus.Entry) ([]byte, error) {
	jsonEntry := Json{
		Time:    int(entry.Time.Unix()),
		Level:   entry.Level.String(),
		Caller:  fmt.Sprintf("%v:%d", entry.Caller.File, entry.Caller.Line),
		Message: entry.Message,
		Data:    entry.Data,
	}
	jsonData, err := json.Marshal(jsonEntry)
	if err != nil {
		return nil, err
	}
	jsonData = append(jsonData, '\n')
	return jsonData, nil
}

func GetLogger(interval time.Duration) *logrus.Logger {
	if config.Application.Env == "prod" {
		return getProdLogger(interval)
	}
	return getDevLogger()
}

func getDevLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&ConsoleFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	logger.AddHook(&StackTraceHook{})
	return logger
}

func getProdLogger(interval time.Duration) *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&ConsoleFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	logger.AddHook(&StackTraceHook{})
	logger.AddHook(NewFileWriteHook(logrus.InfoLevel, interval))
	return logger
}
