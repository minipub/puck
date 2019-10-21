package puck

import (
	"context"
	"fmt"
	stdlog "log"
	"os"
	"sort"
	"strings"
)

const (
	CtxLogKey = "ctxLog"
	OFF       = iota
	DEBUG
	INFO
	WARN
	ERROR
	DefaultLevel = DEBUG
)

type loggerFields map[string]string

func (f loggerFields) Set(key, value string) {
	f[key] = value
}

type Logger struct {
	level      int
	logger     *stdlog.Logger
	fieldValue string
	loggerFields
}

func GetLogger(ctx context.Context) *Logger {
	if ctx == nil {
		return nil
	}
	if l, ok := ctx.Value(CtxLogKey).(*Logger); ok {
		return l
	}
	return nil
}

func NewLogger() *Logger {
	return &Logger{
		level:        DefaultLevel,
		logger:       stdlog.New(os.Stdout, "", stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile),
		loggerFields: make(map[string]string),
	}
}

func (l *Logger) WrapContextLogger(ctx context.Context) context.Context {
	l.buildFieldValue()
	return context.WithValue(ctx, CtxLogKey, l)
}

func (l *Logger) SetField(key, value string) {
	l.loggerFields.Set(key, value)
}

func (l *Logger) buildFieldValue() {
	if len(l.loggerFields) == 0 {
		return
	}

	fieldKeys := make([]string, 0, len(l.loggerFields))
	for k := range l.loggerFields {
		fieldKeys = append(fieldKeys, k)
	}

	sort.Strings(fieldKeys)

	for k, v := range fieldKeys {
		if k == 0 {
			l.fieldValue = fmt.Sprintf("%s", l.loggerFields[v])
		} else {
			l.fieldValue = fmt.Sprintf("%s: %s", l.loggerFields[v], l.fieldValue)
		}
	}
}

func (l *Logger) SetLevel(level string) {
	l.level = getLevel(level)
}

func getLevel(level string) int {
	level = strings.ToLower(level)

	switch level {
	case "off":
		return OFF
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	default:
		fmt.Printf("Level[ %s ] not found, use [ DEBUG ] as default.\n", level)
		return DefaultLevel
	}
}

func (l *Logger) outputf(format string, v ...interface{}) {
	if l.fieldValue != "" {
		format = fmt.Sprintf("%s: %s", l.fieldValue, format)
	}
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) output(v ...interface{}) {
	if l.fieldValue != "" {
		v = append([]interface{}{
			fmt.Sprintf("%s: ", l.fieldValue)}, v...)
	}
	l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) setDebugPrefix() {
	l.logger.SetPrefix("[ DEBUG ] ")
}

func (l *Logger) setInfoPrefix() {
	l.logger.SetPrefix("[ INFO ] ")
}

func (l *Logger) setWarnPrefix() {
	l.logger.SetPrefix("[ WARN ] ")
}

func (l *Logger) setErrorPrefix() {
	l.logger.SetPrefix("[ ERROR ] ")
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l == nil || DEBUG < l.level {
		return
	}
	l.setDebugPrefix()
	l.outputf(format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	if l == nil || DEBUG < l.level {
		return
	}
	l.setDebugPrefix()
	l.output(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l == nil || INFO < l.level {
		return
	}
	l.setInfoPrefix()
	l.outputf(format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	if l == nil || INFO < l.level {
		return
	}
	l.setInfoPrefix()
	l.output(v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l == nil || WARN < l.level {
		return
	}
	l.setWarnPrefix()
	l.outputf(format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	if l == nil || WARN < l.level {
		return
	}
	l.setWarnPrefix()
	l.output(v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l == nil || ERROR < l.level {
		return
	}
	l.setErrorPrefix()
	l.outputf(format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	if l == nil || ERROR < l.level {
		return
	}
	l.setErrorPrefix()
	l.output(v...)
}
