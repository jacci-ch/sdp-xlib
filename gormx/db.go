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

func GenGormDB(cfg *Config) (*gorm.DB, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gormx: %v", err)
	}

	return db, err
}

func init() {
	cfg, err := LoadConfigs()
	if err != nil {
		logx.Logger.Fatal(err)
		return
	}

	db, err := GenGormDB(cfg)
	if err != nil {
		logx.Logger.Fatal(err)
	}

	if cfg.Debug {
		db = db.Debug()
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&DB)), unsafe.Pointer(db))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&currCfg)), unsafe.Pointer(cfg))

	logx.Logger.Info("gormx: database dsn: ", cfg.Dsn)
}
