package simplelogger

import (
	"bufio"
)

const (
	levelDebug int = iota
	levelInfo
	levelWarn
	levelError
	levelCrit
)

// A Logger writes log messages to a log file.
type Logger struct {
	logLevel int
	encoding string
	writer   bufio.Writer
}

func (l Logger) log(logLevel int, message string) {}

func (l Logger) Custom(lvl string, message string) {}

func (l Logger) Debug(message string) {
	if l.logLevel > levelError {
		return
	}
}

func (l Logger) Info(message string) {
	if l.logLevel > levelInfo {
		return
	}
}

func (l Logger) Warn(message string) {}

func (l Logger) Error(message string) {}

func (l Logger) Crit(message string) {}
