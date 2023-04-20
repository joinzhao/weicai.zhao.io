package etcdx

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

const schema = "etcd"

// Resolver 实现grpc的grpc.resolve.Builder接口的Build与Scheme方法
// resolver is the implementaion of grpc.resolve.Builder
type Resolver struct {
	endpoints []string
	service   string
	cli       *clientv3.Client
	cc        resolver.ClientConn
}

// NewResolver return resolver builder
// endpoints example: http://127.0.0.1:2379 http://127.0.0.1:12379 http://127.0.0.1:22379"
// service is service name
func NewResolver(endpoints []string, service string) resolver.Builder {
	return &Resolver{endpoints: endpoints, service: service}
}

// Scheme return etcd schema
func (r *Resolver) Scheme() string {
	// 最好用这种，因为grpc resolver.Register(r)在注册时，会取scheme，如果一个系统有多个grpc发现，就会覆盖之前注册的
	return schema + "-" + r.service
}

// ResolveNow
func (r *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {
}

// Close
func (r *Resolver) Close() {
}

// Build to resolver.Resolver
// 实现grpc.resolve.Builder接口的方法
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error

	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints: r.endpoints,
	})
	if err != nil {
		return nil, fmt.Errorf("grpclb: create clientv3 client failed: %v", err)
	}

	r.cc = cc

	go r.watch(r.service)

	return r, nil
}

func (r *Resolver) watch(prefix string) {
	addrDict := make(map[string]resolver.Address)

	update := func() {
		addrList := make([]resolver.Address, 0, len(addrDict))
		for _, v := range addrDict {
			addrList = append(addrList, v)
		}
		_ = r.cc.UpdateState(resolver.State{Addresses: addrList})
	}
	resp, err := r.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err == nil {
		for i, kv := range resp.Kvs {
			info := &ServiceInfo{}
			_ = json.Unmarshal(kv.Value, info)
			addrDict[string(resp.Kvs[i].Value)] = resolver.Address{Addr: info.IP}
		}
	}

	update()

	rch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for n := range rch {
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				info := &ServiceInfo{}
				_ = json.Unmarshal(ev.Kv.Value, info)
				addrDict[string(ev.Kv.Key)] = resolver.Address{Addr: info.IP}
			case mvccpb.DELETE:
				delete(addrDict, string(ev.PrevKv.Key))
			}
		}
		update()
	}
}
