package etcdx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/logx"
	etcd "go.etcd.io/etcd/client/v3"
	"strings"
	"sync/atomic"
	"unsafe"
)

var (
	gClient *etcd.Client
)

// GenEtcdClient
// Generate a etcd client with given configuration.
func GenEtcdClient(cfg *Config) (*etcd.Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	etcdCfg := etcd.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeoutDuration,
	}

	client, err := etcd.New(etcdCfg)
	if err != nil {
		return nil, errors.New("etcdx: " + err.Error())
	}

	return client, nil
}

func init() {
	cfg, err := LoadConfigs()
	if err != nil {
		logx.Logger.Error(err)
		return
	}

	client, err := GenEtcdClient(cfg)
	if err != nil {
		logx.Logger.Error(err)
		panic(err)
	}

	// Atomic set the etcd client pointer
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&gClient)), unsafe.Pointer(client))
	currCfg = cfg

	logx.Logger.Infof("etcdx: endpoints is %v", strings.Join(cfg.Endpoints, ","))
}
