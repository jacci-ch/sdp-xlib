package etcdx

import (
	"context"
	etcd "go.etcd.io/etcd/client/v3"
)

func getWithContext(key string, opts ...etcd.OpOption) (*etcd.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), currCfg.ReadTimeoutDuration)
	defer cancel()

	return gClient.Get(ctx, key, opts...)
}

func GetBytes(key string) ([]byte, error) {
	rsp, err := getWithContext(key)
	if err == nil && rsp.Count > 0 && len(rsp.Kvs) != 0 {
		return rsp.Kvs[0].Value, nil
	}
	return nil, err
}

func Get(key string) (string, error) {
	bytes, err := GetBytes(key)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}

func GetAs(key string, i any) error {
	return nil
}

// GetKeys
// Request for only keys with given prefix.
func GetKeys(prefix string) ([]string, error) {
	rsp, err := getWithContext(prefix, etcd.WithPrefix(), etcd.WithKeysOnly())
	if err != nil {
		return nil, err
	}

	result := make([]string, len(rsp.Kvs))
	for cc, kv := range rsp.Kvs {
		result[cc] = string(kv.Key)
	}

	return result, nil
}

func GetByPrefix(prefix string) ([]*KV, error) {
	rsp, err := getWithContext(prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}

	result := make([]*KV, len(rsp.Kvs))
	for cc, item := range rsp.Kvs {
		result[cc] = &KV{Key: string(item.Key), Value: string(item.Value)}
	}

	return result, nil
}

func GetByPrefixAsMap(prefix string) (map[string]string, error) {
	rsp, err := getWithContext(prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, item := range rsp.Kvs {
		result[string(item.Key)] = string(item.Value)
	}

	return result, nil
}
