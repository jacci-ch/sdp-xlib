// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package logx

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
)

// CallerPrettifyFuncForText
//
// The caller-prettify function for logrus.TextFormatter.
func CallerPrettifyFuncForText(frame *runtime.Frame) (string, string) {
	return fmt.Sprintf("%s()", filepath.Base(frame.Function)), ""
}

// CallerPrettifyFuncForJSON
//
// The caller-prettify function for logrus.JSONFormatter.
func CallerPrettifyFuncForJSON(frame *runtime.Frame) (string, string) {
	return filepath.Base(frame.Function) + "()", filepath.Base(frame.File)
}

// Logger
//
// Usage:
// logx.Logger.Info(...)
//
// TODO: Write a logrus hook to parse the caller frame so we can use logx as: logx.Info(...)
var Logger *logrus.Logger

// init
//
// Create logger object with default configuration.
// See DefCfg object for more.
func init() {
	if err := ApplyAllConfigs(DefCfg); err != nil {
		panic(err)
	}
}
