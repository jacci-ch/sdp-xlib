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
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// CallerPrettify - Prettify the caller.
func CallerPrettify(frame *runtime.Frame) (string, string) {
	return fmt.Sprintf("%s()", filepath.Base(frame.Function)), ""
}

// newFileWriter - Generate a new file output writer with given configuration.
// A disk file will be opened with os.O_APPEND flag.
//
// The file writer will be wrapped by lumberjack.Logger if the rotation
// feature is enabled.
func newFileWriter(cfg *Config) (io.Writer, error) {
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

// NewLogger - Create a new logger, ignore errors.
func NewLogger(cfg *Config) *logrus.Logger {
	logger := logrus.New()

	logger.AddHook(&CallerHook{})
	logger.SetLevel(cfg.Level)
	logger.SetReportCaller(cfg.WithCaller)

	switch strings.ToLower(cfg.FormatterType) {
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:    cfg.FormatterTimestampEnable,
			TimestampFormat:  cfg.FormatterTimestampFormat,
			CallerPrettyfier: CallerPrettify,
			PadLevelText:     cfg.FormatterPadLevelText,
		})
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp: !cfg.FormatterTimestampEnable,
			TimestampFormat:  cfg.FormatterTimestampFormat,
			CallerPrettyfier: CallerPrettify,
		})
	}

	var writer io.Writer
	if cfg.OutputStdoutEnable {
		writer = os.Stdout
	}

	// Combine stdout writer and file writer into a multi-writer
	// if both stdout and file-output are enabled.
	if cfg.OutputFileEnable {
		if fileWriter, err := newFileWriter(cfg); err == nil {
			if writer == nil {
				writer = fileWriter
			} else {
				writer = io.MultiWriter(writer, fileWriter)
			}
		}
	}

	logger.SetOutput(writer)
	return logger
}
