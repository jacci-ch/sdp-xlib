package logx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/cfgx/cfgv"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"unsafe"
)

const (
	KeyDefault = cfgv.Default
)

var (
	ErrNoOutput = errors.New("logx: no output enabled")
)

func GetLevel(str string) logrus.Level {
	if level, err := logrus.ParseLevel(str); err == nil {
		return level
	}
	return logrus.InfoLevel
}

func GenFileWriter(cfg *Config) (io.Writer, error) {
	if err := os.MkdirAll(cfg.OutputFileDir, 0766); err != nil {
		return nil, errors.New("logx:" + err.Error())
	}

	var writer io.Writer
	filename := filepath.Join(cfg.OutputFileDir, cfg.OutputFileName)

	if cfg.OutputFileRotationEnable {
		writer = &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    cfg.OutputFileRotationMaxSize,
			MaxBackups: cfg.OutputFileRotationMaxBackups,
			MaxAge:     cfg.OutputFileRotationMaxAge,
			LocalTime:  true,
			Compress:   cfg.OutputFileRotationCompress,
		}
	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return nil, errors.New("logx:" + err.Error())
		}

		writer = file
	}

	return writer, nil
}

func GenLoggerWithConfig(cfg *Config) (*logrus.Logger, error) {
	logger := logrus.New()

	logger.SetLevel(GetLevel(cfg.Level))
	logger.SetReportCaller(cfg.WithCaller)

	switch strings.ToLower(cfg.FormatterType) {
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:    cfg.FormatterTimestampEnable,
			TimestampFormat:  cfg.FormatterTimestampFormat,
			CallerPrettyfier: CallerPrettifyFuncForText,
			PadLevelText:     cfg.FormatterPadLevelText,
		})
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp: !cfg.FormatterTimestampEnable,
			TimestampFormat:  cfg.FormatterTimestampFormat,
			CallerPrettyfier: CallerPrettifyFuncForJSON,
		})
	}

	var writer io.Writer
	if cfg.OutputStdoutEnable {
		writer = os.Stdout
	}

	if cfg.OutputFileEnable {
		if fileWriter, err := GenFileWriter(cfg); err == nil {
			if writer == nil {
				writer = fileWriter
			} else {
				writer = io.MultiWriter(writer, fileWriter)
			}
		} else {
			return nil, errors.New("logx:" + err.Error())
		}
	}

	logger.SetOutput(writer)
	return logger, nil
}

func ApplyAllConfigs(cfg *Config) error {
	if !cfg.OutputStdoutEnable && !cfg.OutputFileEnable {
		return ErrNoOutput
	}

	logger, err := GenLoggerWithConfig(cfg)
	if err != nil {
		return err
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&Logger)), unsafe.Pointer(logger))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&currCfg)), unsafe.Pointer(cfg))

	return nil
}

// ApplyConfigs
// This function MUST be called after init() function in this package.
func ApplyConfigs(cfgx cfgv.ValueGetter) {
	cfg := Config{}

	_ = cfgx.ToStr(KeyDefault, KeyLevel, &cfg.Level, defCfg.Level)
	_ = cfgx.ToBool(KeyDefault, KeyWithCaller, &cfg.WithCaller, defCfg.WithCaller)
	_ = cfgx.ToStr(KeyDefault, KeyFormatterType, &cfg.FormatterType, defCfg.FormatterType)
	_ = cfgx.ToBool(KeyDefault, KeyTimestampEnable, &cfg.FormatterTimestampEnable, defCfg.FormatterTimestampEnable)
	_ = cfgx.ToStr(KeyDefault, KeyTimestampFormat, &cfg.FormatterTimestampFormat, defCfg.FormatterTimestampFormat)
	_ = cfgx.ToBool(KeyDefault, KeyPadLevelText, &cfg.FormatterPadLevelText, defCfg.FormatterPadLevelText)

	_ = cfgx.ToBool(KeyDefault, KeyStdoutEnable, &cfg.OutputStdoutEnable, defCfg.OutputStdoutEnable)
	_ = cfgx.ToBool(KeyDefault, KeyFileEnable, &cfg.OutputFileEnable, defCfg.OutputFileEnable)
	_ = cfgx.ToStr(KeyDefault, KeyFileDir, &cfg.OutputFileDir, defCfg.OutputFileDir)
	_ = cfgx.ToStr(KeyDefault, KeyFileName, &cfg.OutputFileName, defCfg.OutputFileName)

	_ = cfgx.ToBool(KeyDefault, KeyFileRotationEnable, &cfg.OutputFileRotationEnable, defCfg.OutputFileRotationEnable)
	_ = cfgx.ToInt(KeyDefault, KeyFileRotationMaxSize, &cfg.OutputFileRotationMaxSize, defCfg.OutputFileRotationMaxSize)
	_ = cfgx.ToInt(KeyDefault, KeyFileRotationMaxBackups, &cfg.OutputFileRotationMaxBackups, defCfg.OutputFileRotationMaxBackups)
	_ = cfgx.ToInt(KeyDefault, KeyFileRotationMaxAge, &cfg.OutputFileRotationMaxAge, defCfg.OutputFileRotationMaxAge)
	_ = cfgx.ToBool(KeyDefault, KeyFileRotationCompress, &cfg.OutputFileRotationCompress, defCfg.OutputFileRotationCompress)

	if err := ApplyAllConfigs(&cfg); err != nil {
		panic(err)
	}

	if cfg.OutputFileEnable {
		Logger.Infof("logx: output log message to %v", cfg.OutputFileDir)
	}
}
