package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"io"
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

func main() {
	f, _ := os.Create("Example.json")
	w := io.Writer(f)

	log := NewLogger(levelDebug, w)
	for i := 0; i < 10000; i++ {
		log.Warn(fmt.Sprintf("Log %d", i))
	}

	log.FlushAll()
}

// A Log represents a message log and it's properties.
type Log struct {
	Prefix  string `json:"prefix"`
	Level   int    `json:"level"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

// A Logger writes log messages to a log destination in json format.
type Logger struct {
	threshold int // threshold represents the lower bound for log levels that will be logged.
	out       io.Writer
	// customLevels map[string]int
	logs map[int]Log
}

// NewLogger returns a new Logger.
func NewLogger(threshold int, out io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		out:       out,
		logs:      make(map[int]Log),
	}
}

// NewLog returns a new Log.
func newLog(prefix string, level int, message string) Log {
	return Log{
		Prefix:  prefix,
		Level:   level,
		Message: message,
		Time:    time.Now().String(),
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
func (l *Logger) log(logLevel int, message string) {
	// Get the appropriate prefix for the log message.
	prefix := l.levelPrefix(logLevel)

	log := newLog(prefix, logLevel, message)
	l.logs[len(l.logs)] = log
}

// writeJSON encodes a log in json format and writes it to out.
func (l *Log) writeJSON(w io.Writer, b json.Encoder) {
	b.Encode(l)
}

// FlushAll encodes all of Logger's stored logs to JSON at the target out.
func (l Logger) FlushAll() {
	b := json.NewEncoder(l.out)
	var wg sync.WaitGroup

	for _, log := range l.logs {
		wg.Add(1)
		go func(currentLog Log) {
			defer wg.Done()
			currentLog.writeJSON(l.out, *b)
		}(log)
	}

	wg.Wait()
	fmt.Println("All logged!")
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
