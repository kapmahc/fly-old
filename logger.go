package fly

import "log"

const (
	// DEBUG level
	DEBUG = iota
	// INFO level
	INFO = iota
	// WARN level
	WARN = iota
	// ERROR level
	ERROR = iota
)

// Logger logger
type Logger struct {
	level int
}

// Level set level
func (p *Logger) Level(l int) {
	p.level = l
}

// Debug print debug message
func (p *Logger) Debug(args ...interface{}) {
	p.print(DEBUG, args...)
}

// Info print info message
func (p *Logger) Info(args ...interface{}) {
	p.print(INFO, args...)
}

// Warn print warn message
func (p *Logger) Warn(args ...interface{}) {
	p.print(WARN, args...)
}

// Error print error message
func (p *Logger) Error(args ...interface{}) {
	p.print(ERROR, args...)
}

func (p *Logger) print(level int, args ...interface{}) {
	if level < p.level {
		return
	}

	var lvl string
	switch level {
	case INFO:
		lvl = "\033[1;34mINFO\033[0m "
	case WARN:
		lvl = "\033[1;33mWARN\033[0m "
	case ERROR:
		lvl = "\033[1;35mERROR\033[0m"
	default:
		lvl = "DEBUG"
	}
	val := []interface{}{lvl}
	val = append(val, args...)
	log.Println(val...)
}
