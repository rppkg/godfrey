package etcd

import (
	"time"

	etcdv3 "go.etcd.io/etcd/client/v3"
)

type Options struct {
	Endpoints   []string      `json:"endpoints"`
	DialTimeout time.Duration `json:"dial-timeout"`
	Username    string        `json:"username"`
	Password    string        `json:"password"`
}

func Init(opts *Options) (*etcdv3.Client, error) {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints:            opts.Endpoints,
		DialTimeout:          opts.DialTimeout,
		Username:             opts.Username,
		Password:             opts.Password,
	})
	if err != nil {
		return nil, err
	}

	return cli, nil
}
