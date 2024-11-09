package log

import (
	"sync"

	"github.com/charmbracelet/log"
)

const (
	Prefix = "[playground] "
)

const (
	FormatInvalidLogLevel = "invalid log level: %s"
)

var once sync.Once

func init() {
	once.Do(func() {
		log.SetPrefix(Prefix)
	})
}

func SetLevel(level string) {
	init := func(level log.Level) {
		switch level {
		case log.DebugLevel:
			log.SetCallerOffset(2)
		case log.InfoLevel:
		case log.WarnLevel:
			log.SetCallerOffset(2)
		case log.ErrorLevel:
			log.SetCallerOffset(2)
		case log.FatalLevel:
			log.SetCallerOffset(2)
		}
	}
	if level, err := log.ParseLevel(level); err == nil {
		log.SetLevel(level)
		defer init(level)
	} else {
		log.Warnf(FormatInvalidLogLevel, level)
		return
	}
}

var (
	Debug = log.Debug
	Info  = log.Info
	Warn  = log.Warn
	Error = log.Error
	Fatal = log.Fatal

	Debugf = log.Debugf
	Infof  = log.Infof
	Warnf  = log.Warnf
	Errorf = log.Errorf
	Fatalf = log.Fatalf
)
