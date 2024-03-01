package logger

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "runtime"
)

type LogLevel int

const (
    // All levels
    All LogLevel = iota
    // Fine for fine-grained information
    Fine
    // Debug for diagnostic information
    Debug
    // Info for normal application behavior
    Info
    // Warn for potentially harmful situations
    Warn
    // Error for not so severe errors
    Error
    // Fatal for very severe errors
    Fatal
    // Off to turn off logging
    Off
)

// MyLogger wraps log.Logger and adds support for logging levels
type MyLogger struct {
    level     LogLevel
    stdLogger *log.Logger
}

// New creates a new logger
func New(name string) *MyLogger {
    logger := new(MyLogger)
    logger.level = Info
    logger.stdLogger = log.New(os.Stdout, name + " ", 0)
    return logger
}

// SetLevel to change logging level
func (logger *MyLogger) SetLevel(level LogLevel) {
    logger.level = level
}

func (logger MyLogger) log(prefix string, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    if _, file, line, ok := runtime.Caller(2); ok {
        logger.stdLogger.Print(fmt.Sprintf("%-5s %s:%d %s", prefix, filepath.Base(file), line, message))
    } else {
        logger.stdLogger.Print(fmt.Sprintf("%-5s %s", prefix, message))
    }
}

// Fine for fine-grained information
func (logger MyLogger) Fine(format string, args ...interface{}) {
    if Fine >= logger.level {
        logger.log("FINE", format, args...)
    }
}

// Debug for diagnostic information
func (logger MyLogger) Debug(format string, args ...interface{}) {
    if Debug >= logger.level {
        logger.log("DEBUG", format, args...)
    }
}

// Info for normal application behavior
func (logger MyLogger) Info(format string, args ...interface{}) {
    if Info >= logger.level {
        logger.log("INFO", format, args...)
    }
}

// Warn for potentially harmful situations
func (logger MyLogger) Warn(format string, args ...interface{}) {
    if Warn >= logger.level {
        logger.log("WARN", format, args...)
    }
}

// Error for not so severe errors
func (logger MyLogger) Error(format string, args ...interface{}) {
    if Error >= logger.level {
        logger.log("ERROR", format, args...)
    }
}

// Fatal for very severe errors
func (logger MyLogger) Fatal(format string, args ...interface{}) {
    if Fatal >= logger.level {
        logger.log("FATAL", format, args...)
    }
}
