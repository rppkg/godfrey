package mongo

import (
	"context"
	"time"

	gomongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Options struct {
	URL        string        `json:"url"`
	Database   string        `json:"database"`
	Collection string        `json:"collection"`
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	Timeout    time.Duration `json:"timeout"`
}

func Init(o *Options) (*gomongo.Client, error) {
	opts := options.Client().ApplyURI(o.URL).SetReadPreference(readpref.Primary())
	if o.Timeout > 0 {
		opts.SetConnectTimeout(o.Timeout).SetSocketTimeout(o.Timeout).SetServerSelectionTimeout(o.Timeout)
	}

	if o.Username != "" || o.Password != "" {
		opts.SetAuth(options.Credential{
			AuthSource: o.Database,
			Username:   o.Username,
			Password:   o.Password,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	client, err := gomongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}