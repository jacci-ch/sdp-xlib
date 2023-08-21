// Copyright 2023 to now() The SDP Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logx

import (
	"github.com/jacci-ch/sdp-xlib/cfgx"
	"github.com/sirupsen/logrus"
	"time"
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
	KeyFileRotationCompress   = "logger.output.file.rotation.compress.enable"
)

var (
	Cfg    *Config
	DefCfg = &Config{
		Level:                        logrus.InfoLevel,
		WithCaller:                   false,
		FormatterType:                "text",
		FormatterTimestampEnable:     true,
		FormatterTimestampFormat:     time.DateTime + ".000",
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
)

type Config struct {
	Level      logrus.Level
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

// parseLevel - translate level string (e.g "info") to logrus.Level
// (e.g logrus.InfoLevel) value. This function returns logrus.InfoLevel
// if the str value is not a valid level string.
func parseLevel(str string) logrus.Level {
	if level, err := logrus.ParseLevel(str); err == nil {
		return level
	}
	return logrus.InfoLevel
}

// loadConfigs - loads and parses all configuration items from cfgx.
// This function will ignore errors.
func loadConfigs() *Config {
	cfg, value := &Config{}, ""

	_ = cfgx.AsStr(KeyLevel, &value, DefCfg.Level.String())
	cfg.Level = parseLevel(value)

	_ = cfgx.AsBool(KeyWithCaller, &cfg.WithCaller, DefCfg.WithCaller)
	_ = cfgx.AsStr(KeyFormatterType, &cfg.FormatterType, DefCfg.FormatterType)
	_ = cfgx.AsBool(KeyTimestampEnable, &cfg.FormatterTimestampEnable, DefCfg.FormatterTimestampEnable)
	_ = cfgx.AsStr(KeyTimestampFormat, &cfg.FormatterTimestampFormat, DefCfg.FormatterTimestampFormat)
	_ = cfgx.AsBool(KeyPadLevelText, &cfg.FormatterPadLevelText, DefCfg.FormatterPadLevelText)

	_ = cfgx.AsBool(KeyStdoutEnable, &cfg.OutputStdoutEnable, DefCfg.OutputStdoutEnable)
	_ = cfgx.AsBool(KeyFileEnable, &cfg.OutputFileEnable, DefCfg.OutputFileEnable)
	_ = cfgx.AsStr(KeyFileDir, &cfg.OutputFileDir, DefCfg.OutputFileDir)
	_ = cfgx.AsStr(KeyFileName, &cfg.OutputFileName, DefCfg.OutputFileName)

	_ = cfgx.AsBool(KeyFileRotationEnable, &cfg.OutputFileRotationEnable, DefCfg.OutputFileRotationEnable)
	_ = cfgx.AsInt(KeyFileRotationMaxSize, &cfg.OutputFileRotationMaxSize, DefCfg.OutputFileRotationMaxSize)
	_ = cfgx.AsInt(KeyFileRotationMaxBackups, &cfg.OutputFileRotationMaxBackups, DefCfg.OutputFileRotationMaxBackups)
	_ = cfgx.AsInt(KeyFileRotationMaxAge, &cfg.OutputFileRotationMaxAge, DefCfg.OutputFileRotationMaxAge)
	_ = cfgx.AsBool(KeyFileRotationCompress, &cfg.OutputFileRotationCompress, DefCfg.OutputFileRotationCompress)

	return cfg
}
