// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package cfgx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/cfgx/cfgv"
	"github.com/jacci-ch/sdp-xlib/logx"
	"github.com/jacci-ch/sdp-xlib/osx"
	"gopkg.in/ini.v1"
)

var (
	CfgFiles = []string{
		"./sdp.conf",
		"./conf/sdp.conf",
		"./conf/sdp/sdp.conf",
		"./cfg/sdp.conf",
		"./cfg/sdp/sdp.conf",
		"./etc/sdp.conf",
		"./etc/sdp/sdp.conf",
		"/etc/sdp.conf",
		"/etc/sdp/sdp.conf",
		"/user/local/etc/sdp.conf",
		"/user/local/etc/sdp/sdp.conf",
	}

	CfgFile     = ""
	ErrNotFound = errors.New("cfgx: file not found")
)

// ProbeCfgFile
//
// Detect the available configuration file in the hard-coded list.
func ProbeCfgFile() string {
	for _, name := range CfgFiles {
		if osx.Exist(name) {
			return name
		}
	}

	return ""
}

// loadCfgFromFile
//
// Probe the configuration file and load all configuration
// items into a value keeper object.
func loadCfgFromFile() error {
	name := ProbeCfgFile()
	if len(name) == 0 {
		return ErrNotFound
	}

	CfgFile = name // for debug
	iniFile, err := ini.Load(name)
	if err != nil {
		return errors.New("cfgx: " + err.Error())
	}

	// Read all kv-pairs of all sections in the configuration
	// file into a new value keeper.
	valueKeeper := NewValueKeeper()
	for _, section := range iniFile.Sections() {
		values := make(map[string]*cfgv.Value)
		for _, key := range section.Keys() {
			values[key.Name()] = cfgv.Value(key.Value()).Addr()
		}
		valueKeeper[section.Name()] = values
	}

	gValueKeeper = valueKeeper
	return nil
}

// init
//
// Load configurations and apply logger configurations to logx.
// To avoid the 'import
func init() {
	err := loadCfgFromFile()
	if err != nil {
		logx.Logger.Fatal(err)
		return
	}

	logx.ApplyConfigs(Def)
	logx.Logger.Infof("cfgx: use configuration file: %v", CfgFile)
}