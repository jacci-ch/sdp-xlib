// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

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
	"time"
	"unsafe"
)

const (
	KeyLevel                  = "logger.level"
	KeyWithCaller             = "logger.with.caller"
	KeyFormatterType          = "logger.formatter.type"
	KeyTimestampEnable        = "logger.formatter.timestamp.enable"
	KeyTimestampFormat        = "logger.formatter.timestamp.format"
	KeyPadLevelText           = "logger.formatter.pad.level.text"
	KeyStdoutEnable           = "logger.output.stdout.enable"
	KeyFileEnable             = "logger.output.file.enable"
	KeyFileDir                = "logger.output.file.dir"
	KeyFileName               = "logger.output.file.name"
	KeyFileRotationEnable     = "logger.output.file.rotation.enable"
	KeyFileRotationMaxSize    = "logger.output.file.rotation.max.size"
	KeyFileRotationMaxBackups = "logger.output.file.rotation.max.backups"
	KeyFileRotationMaxAge     = "logger.output.file.rotation.max.age"
	KeyFileRotationCompress   = "logger.output.file.rotation.compress"

	DatetimeFormat = time.DateTime + ".000"
)

var (
	DefCfg = &Config{
		Level:                        "info",
		WithCaller:                   false,
		FormatterType:                "text",
		FormatterTimestampEnable:     true,
		FormatterTimestampFormat:     DatetimeFormat,
		FormatterPadLevelText:        true,
		OutputStdoutEnable:           true,
		OutputFileEnable:             false,
		OutputFileDir:                "/var/log/sdp",
		OutputFileName:               "sdp.log",
		OutputFileRotationEnable:     false,
		OutputFileRotationMaxSize:    100,
		OutputFileRotationMaxBackups: 3,
		OutputFileRotationMaxAge:     30,
		OutputFileRotationCompress:   false,
	}

	Cfg *Config
)

// Config
//
// A struct holds all configurations.
type Config struct {
	Level      string
	WithCaller bool

	// Text formatter configuration
	FormatterType            string
	FormatterTimestampEnable bool
	FormatterTimestampFormat string
	FormatterPadLevelText    bool

	// stdout output configuration.
	OutputStdoutEnable bool

	// file output configuration.
	OutputFileEnable bool
	OutputFileDir    string
	OutputFileName   string

	// file rotation configuration
	OutputFileRotationEnable     bool
	OutputFileRotationMaxSize    int
	OutputFileRotationMaxBackups int
	OutputFileRotationMaxAge     int
	OutputFileRotationCompress   bool
}

// GetLevel
//
// Translate level string (e.g "info") to logrus.Level (e.g logrus.InfoLevel) value.
// This function returns logrus.InfoLevel if the str value is not
// a level string.
func GetLevel(str string) logrus.Level {
	if level, err := logrus.ParseLevel(str); err == nil {
		return level
	}
	return logrus.InfoLevel
}

// NewFileWriter
//
// Generate a new file output writer with given configuration. A disk file
// will be opened with os.O_APPEND flag.
//
// The file writer will be wrapped by lumberjack.Logger if the rotation
// feature is enabled.
func NewFileWriter(cfg *Config) (io.Writer, error) {
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

// NewLogger
//
// Generate a new logrus.Logger object with given configuration. This function
// returns nil object and error if any errors occur.
func NewLogger(cfg *Config) (*logrus.Logger, error) {
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
	default:
		return nil, errors.New("logx: invalid configuration")
	}

	var writer io.Writer
	if cfg.OutputStdoutEnable {
		writer = os.Stdout
	}

	// Combine stdout writer and file writer into a multi-writer
	// if both stdout and file-output are enabled.
	if cfg.OutputFileEnable {
		if fileWriter, err := NewFileWriter(cfg); err == nil {
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

// ApplyAllConfigs
//
// Generate default logrus.Logger object with given configurations and
// store the object to global value Logger.
//
// The cfg will be stored into global value Cfg too.
func ApplyAllConfigs(cfg *Config) error {
	if !cfg.OutputStdoutEnable && !cfg.OutputFileEnable {
		return errors.New("logx: no output enabled")
	}

	logger, err := NewLogger(cfg)
	if err != nil {
		return err
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&Logger)), unsafe.Pointer(logger))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&Cfg)), unsafe.Pointer(cfg))

	return nil
}

// ApplyConfigs
//
// Load configurations from a cfgv.ValueGetter and generate the default
// logrus.Logger object with the loaded configurations.
//
// This function MUST be called after init() function in this package.
// We need use logx module to output info messages in cfgx module, and we
// can't import cfgx module in logx modules to avoid 'import cycle' error.
//
// So this function will be called in cfgx module.
func ApplyConfigs(cfgx cfgv.DefaultValueGetter) {
	cfg := Config{}

	_ = cfgx.ToStr(KeyLevel, &cfg.Level, DefCfg.Level)
	_ = cfgx.ToBool(KeyWithCaller, &cfg.WithCaller, DefCfg.WithCaller)
	_ = cfgx.ToStr(KeyFormatterType, &cfg.FormatterType, DefCfg.FormatterType)
	_ = cfgx.ToBool(KeyTimestampEnable, &cfg.FormatterTimestampEnable, DefCfg.FormatterTimestampEnable)
	_ = cfgx.ToStr(KeyTimestampFormat, &cfg.FormatterTimestampFormat, DefCfg.FormatterTimestampFormat)
	_ = cfgx.ToBool(KeyPadLevelText, &cfg.FormatterPadLevelText, DefCfg.FormatterPadLevelText)

	_ = cfgx.ToBool(KeyStdoutEnable, &cfg.OutputStdoutEnable, DefCfg.OutputStdoutEnable)
	_ = cfgx.ToBool(KeyFileEnable, &cfg.OutputFileEnable, DefCfg.OutputFileEnable)
	_ = cfgx.ToStr(KeyFileDir, &cfg.OutputFileDir, DefCfg.OutputFileDir)
	_ = cfgx.ToStr(KeyFileName, &cfg.OutputFileName, DefCfg.OutputFileName)

	_ = cfgx.ToBool(KeyFileRotationEnable, &cfg.OutputFileRotationEnable, DefCfg.OutputFileRotationEnable)
	_ = cfgx.ToInt(KeyFileRotationMaxSize, &cfg.OutputFileRotationMaxSize, DefCfg.OutputFileRotationMaxSize)
	_ = cfgx.ToInt(KeyFileRotationMaxBackups, &cfg.OutputFileRotationMaxBackups, DefCfg.OutputFileRotationMaxBackups)
	_ = cfgx.ToInt(KeyFileRotationMaxAge, &cfg.OutputFileRotationMaxAge, DefCfg.OutputFileRotationMaxAge)
	_ = cfgx.ToBool(KeyFileRotationCompress, &cfg.OutputFileRotationCompress, DefCfg.OutputFileRotationCompress)

	if err := ApplyAllConfigs(&cfg); err != nil {
		panic(err)
	}

	if cfg.OutputFileEnable {
		Logger.Infof("logx: write log message to file %v", cfg.OutputFileDir)
	}
}
