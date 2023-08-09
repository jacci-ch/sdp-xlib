package logx

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
)

func CallerPrettifyFuncForText(frame *runtime.Frame) (string, string) {
	return fmt.Sprintf("[%v:%v:%s()]", filepath.Base(frame.File), frame.Line, filepath.Base(frame.Function)), ""
}

func CallerPrettifyFuncForJSON(frame *runtime.Frame) (string, string) {
	return filepath.Base(frame.Function) + "()", filepath.Base(frame.File)
}

// Logger
// TODO: Write a logrus hook to parse the caller frame so we can use logx as: logx.Info(...)
var Logger *logrus.Logger

func init() {
	if err := ApplyAllConfigs(defCfg); err != nil {
		panic(err)
	}
}
