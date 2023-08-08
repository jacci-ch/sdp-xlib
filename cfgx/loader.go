package cfgx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/cfgx/cfgv"
	"github.com/jacci-ch/sdp-xlib/logx"
	"github.com/jacci-ch/sdp-xlib/osx"
	"gopkg.in/ini.v1"
)

var (
	CfgFilePaths = []string{
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

func probeCfgFile() string {
	for _, name := range CfgFilePaths {
		if osx.Exist(name) {
			return name
		}
	}

	return ""
}

func loadCfgFromFile() error {
	name := probeCfgFile()
	if len(name) == 0 {
		return ErrNotFound
	}

	CfgFile = name // for debug
	iniFile, err := ini.Load(name)
	if err != nil {
		return errors.New("cfgx: " + err.Error())
	}

	valueKeeper := newValueKeeper()
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

func init() {
	if err := loadCfgFromFile(); err == nil {
		logx.ApplyConfigs(gValueKeeper)

		logx.Logger.Infof("cfgx: use configuration file: %v", CfgFile)
	} else if err == ErrNotFound {
		return // All use default configurations.
	} else {
		panic(err)
	}
}
