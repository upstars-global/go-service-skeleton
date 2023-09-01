package logger

import (
	"fmt"
	"os"

	"github.com/upstars-global/go-service-skeleton/pkg/argumentsresolver"
	"github.com/upstars-global/go-service-skeleton/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Interface interface {
	With(args ...interface{}) Interface
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

type Logger struct {
	zap *zap.SugaredLogger
}

var tsFormat = "2006-01-02 15:04:05.000000"

func New(
	args argumentsresolver.ArgumentsInterface,
	logCfg config.LoggerConfigProvider,
) (logger Interface, err error) {
	zapCfg := zap.NewProductionConfig()
	zapCfg.OutputPaths = []string{"stdout"}

	verbose, err := args.GetBool(argumentsresolver.ArgumentConfigVerbose)
	if err != nil {
		panic(err)
	}

	availableLogLevels := map[string]bool{
		"fatal":  true,
		"panic":  true,
		"dpanic": true,
		"error":  true,
		"warn":   true,
		"info":   true,
		"debug":  true,
		"trace":  true,
	}

	level := zap.NewAtomicLevel()
	switch {
	case verbose:
		zapCfg.Development = true
		level.SetLevel(zapcore.DebugLevel)
	case logCfg.GetLoggerLevel() == "":
		level.SetLevel(zapcore.InfoLevel)
	default:
		if _, ok := availableLogLevels[logCfg.GetLoggerLevel()]; ok {
			if err := level.UnmarshalText([]byte(logCfg.GetLoggerLevel())); err != nil {
				panic(err)
			}
		} else {
			fmt.Printf(`Bad log level "%s" passed! (available formats are: fatal, panic, dpanic, error, warn, info, debug, trace)`, logCfg.GetLoggerLevel())
			os.Exit(1)
		}
	}
	zapCfg.Level = level

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.TimeEncoderOfLayout(tsFormat)
	zapCfg.EncoderConfig = encoder

	availableLogFormats := map[string]bool{
		"console": true,
		"json":    true,
	}

	if logCfg.GetLoggerFormat() == "" {
		zapCfg.Encoding = "console"
	} else {
		if _, ok := availableLogFormats[logCfg.GetLoggerFormat()]; ok {
			zapCfg.Encoding = logCfg.GetLoggerFormat()
		} else {
			fmt.Printf(`Bad log format "%s" passed! (available formats are: console, json)`, logCfg.GetLoggerFormat())
			os.Exit(1)
		}
	}

	core, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}

	core = core.WithOptions(zap.AddCallerSkip(1))

	logger = &Logger{core.Sugar()}
	return
}

func (l *Logger) With(args ...interface{}) Interface {
	clone := *l
	clone.zap = clone.zap.With(args...)
	return &clone
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *Logger) Debug(args ...interface{}) {
	l.zap.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *Logger) Info(args ...interface{}) {
	l.zap.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *Logger) Warn(args ...interface{}) {
	l.zap.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *Logger) Error(args ...interface{}) {
	l.zap.Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *Logger) Panic(args ...interface{}) {
	l.zap.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Logger) Fatal(args ...interface{}) {
	l.zap.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.zap.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.zap.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.zap.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.zap.Errorf(template, args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.zap.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.zap.Fatalf(template, args...)
}
