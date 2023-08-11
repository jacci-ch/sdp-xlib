// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package gormx

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync/atomic"
	"unsafe"
)

var (
	DB *gorm.DB
)

// NewDB
//
// Open and generate a new gorm.DB with given configuration.
func NewDB(cfg *Config) (*gorm.DB, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gormx: %v", err)
	}

	return db, err
}

// init
//
// Load configurations and generates a new gorm.DB object. This
// function store the new gorm.DB object into DB global value.
func init() {
	cfg, err := LoadConfigs()
	if err != nil {
		logx.Logger.Fatal(err)
		return
	}

	db, err := NewDB(cfg)
	if err != nil {
		logx.Logger.Fatal(err)
	}

	if cfg.Debug {
		db = db.Debug()
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&DB)), unsafe.Pointer(db))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&Cfg)), unsafe.Pointer(cfg))

	logx.Logger.Info("gormx: database dsn: ", cfg.Dsn)
}
