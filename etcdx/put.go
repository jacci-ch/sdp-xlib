package etcdx

import (
	"context"
	etcd "go.etcd.io/etcd/client/v3"
)

func putWithContext(key, value string, opts ...etcd.OpOption) (*etcd.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), currCfg.WriteTimeoutDuration)
	defer cancel()

	return gClient.Put(ctx, key, value, opts...)
}

func Put(key, value string) error {
	_, err := putWithContext(key, value)
	return err
}

func PutWithLease(key, value string, lease int64) error {
	_, err := putWithContext(key, value, etcd.WithLease(etcd.LeaseID(lease)))
	return err
}
