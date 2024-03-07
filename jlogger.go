package main

import (
	"encoding/json"
	"os"

	//"fmt"
	"io"
	//"os"
	"time"
)

const (
	// Increment log levels by 10 from 0.
	// This leaves room for custom severities that may sit between them in
	// the hierarchy.

	levelDebug = iota * 10

	levelInfo
	levelWarn
	levelError
	levelCrit
)

// A Log represents a message log and it's properties.
type Log struct {
	Prefix  string `json:"prefix"`
	Level   int    `json:"level"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	f, _ := os.Create("Example.json")
	w := io.Writer(f)

	log := NewLogger(levelDebug, w)

	for i := 0; i <= 20; i++ {
		log.Info("Testing the logger to json!")
	}
}

// A Logger writes log messages to a log destination in json format.
type Logger struct {
	threshold    int // threshold represents the lower bound for log levels that will be logged.
	out          io.Writer
	customLevels map[string]int
}

// NewLogger returns a new Logger.
func NewLogger(threshold int, out io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		out:       out,
	}
}

// NewLog returns a new Log.
func newLog(prefix string, level int, message string, time string) *Log {
	return &Log{
		Prefix:  prefix,
		Level:   level,
		Message: message,
		Time:    time,
	}
}

// levelPrefix returns the appropriate prefix for a log level.
func (l Logger) levelPrefix(logLevel int) string {
	var prefix string

	switch logLevel {
	case levelDebug:
		prefix = "DEBUG"
	case levelInfo:
		prefix = "INFO"
	case levelWarn:
		prefix = "WARN"
	case levelError:
		prefix = "ERROR"
	case levelCrit:
		prefix = "CRIT"
	default:
		// TODO: Implement support for custom log levels.
	}
	return prefix
}

// log writes a formatted log of level logLevel to a Logger's out.
func (l Logger) log(logLevel int, message string) {
	// Don't write to log file if the level of the log type
	// in question is less than the logLevel threshold.
	if l.threshold > logLevel {
		return
	}

	// Get the timestamp as soon as the log level is
	// found to pass the threshold.
	timeStamp := time.Now().String()

	// Get the appropriate prefix for the log message.
	prefix := l.levelPrefix(logLevel)

	log := newLog(prefix, logLevel, message, timeStamp)

	log.writeJSON(l.out)
}

func (l *Log) writeJSON(w io.Writer) {
	b, _ := json.Marshal(*l)
	w.Write(b)
}

// Debug writes a Debug level message to a Logger's out.
func (l Logger) Debug(message string) {
	if l.threshold > levelError {
		return
	}

	l.log(levelDebug, message)
}

// Info writes an Info level message to Logger's out.
func (l Logger) Info(message string) {
	if l.threshold > levelInfo {
		return
	}

	l.log(levelInfo, message)
}

// Warn writes a Warn level message to a Logger's out.
func (l Logger) Warn(message string) {
	if l.threshold > levelWarn {
		return
	}

	l.log(levelWarn, message)
}

// Error writes an Error level message to a Logger's out.
func (l Logger) Error(message string) {
	if l.threshold > levelError {
		return
	}

	l.log(levelError, message)
}

// Crit writes a Crit level message to a Logger's out.
func (l Logger) Crit(message string) {
	if l.threshold > levelCrit {
		return
	}

	l.log(levelCrit, message)
}

// Other writes a custom log level to a Logger's out.
func (l Logger) Other() {}

// addLogLevel adds a custom log level to a Logger's customLogs map, in the format
// [logLevel]logName.
// Also adds the custom log level to a Logger's reverseCustomLogs, in the format [logName]logLevel.

// Assign custom log level to unique position in the Logger's hash map.
// The logLevel is the key, so as to prevent duplicate log levels of the same severity
// and keep it interacting well with levelPrefix().
