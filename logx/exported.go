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

import "github.com/sirupsen/logrus"

var (
	gLogger *logrus.Logger
)

func Debugf(format string, args ...any) {
	gLogger.Debugf(format, args...)
}

func Infof(format string, args ...any) {
	gLogger.Infof(format, args...)
}

func Printf(format string, args ...any) {
	gLogger.Printf(format, args...)
}

func Warnf(format string, args ...any) {
	gLogger.Warnf(format, args...)
}

func Warningf(format string, args ...any) {
	gLogger.Warningf(format, args...)
}

func Errorf(format string, args ...any) {
	gLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...any) {
	gLogger.Fatalf(format, args...)
}

func Panicf(format string, args ...any) {
	gLogger.Panicf(format, args...)
}

func Debug(args ...any) {
	gLogger.Debug(args...)
}

func Info(args ...any) {
	gLogger.Info(args...)
}

func Print(args ...any) {
	gLogger.Print(args...)
}

func Warn(args ...any) {
	gLogger.Warn(args...)
}

func Warning(args ...any) {
	gLogger.Warning(args...)
}

func Error(args ...any) {
	gLogger.Error(args...)
}

func Fatal(args ...any) {
	gLogger.Fatal(args...)
}

func FatalErr(err error) error {
	gLogger.Fatal(err)
	return err
}

func Panic(args ...any) {
	gLogger.Panic(args...)
}

func Debugln(args ...any) {
	gLogger.Debugln(args...)
}

func Infoln(args ...any) {
	gLogger.Infoln(args...)
}

func Println(args ...any) {
	gLogger.Println(args...)
}

func Warnln(args ...any) {
	gLogger.Warnln(args...)
}

func Warningln(args ...any) {
	gLogger.Warningln(args...)
}

func Errorln(args ...any) {
	gLogger.Errorln(args...)
}

func Fatalln(args ...any) {
	gLogger.Fatalln(args...)
}

func Panicln(args ...any) {
	gLogger.Panicln(args...)
}
