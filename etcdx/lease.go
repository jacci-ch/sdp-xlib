package etcdx

import (
	"context"
	etcd "go.etcd.io/etcd/client/v3"
)

// Grant
// Ask etcd server for a lease to use.
func Grant(ttl int64) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), currCfg.WriteTimeoutDuration)
	defer cancel()

	rsp, err := gClient.Grant(ctx, ttl)
	if err != nil {
		return 0, err
	}

	return int64(rsp.ID), nil
}

// KeepaliveOnce
// Send keepalive (heartbeat) message to etcd server only once.
func KeepaliveOnce(lease int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), currCfg.WriteTimeoutDuration)
	defer cancel()

	// ignore the result.
	_, err := gClient.KeepAliveOnce(ctx, etcd.LeaseID(lease))
	return err
}
