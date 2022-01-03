// Package provides ...
package config

import (
	"context"
	"log"

	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/asim/go-micro/plugins/config/source/consul/v4"
	mconfig "go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
)

// New - new config and Load
// load source and watch config implement
func New(ctx context.Context) IConfig {
	// yaml encoder
	e := yaml.NewEncoder()

	// new consul source
	consulSource := consul.NewSource(
		consul.WithPrefix(appName),
		consul.StripPrefix(true),
		source.WithEncoder(e),
	)

	// new config
	conf, err := mconfig.NewConfig(
		mconfig.WithReader(
			json.NewReader( // json reader for internal config merge
				reader.WithEncoder(e),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	//load source
	if err := conf.Load(consulSource); err != nil {
		log.Fatal(err)
	}

	// scan data
	r := &root{c: conf}
	if err := conf.Scan(&r); err != nil {
		log.Fatal(err)
	}

	//watch config change
	go watch(ctx, conf)

	return r
}
