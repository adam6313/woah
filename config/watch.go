package config

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/sergi/go-diff/diffmatchpatch"
	"go-micro.dev/v4/config"
)

// watch - watch remote config data
// handle something when data changes ex: restart service etc.
func watch(ctx context.Context, conf config.Config) {
	watcher, _ := conf.Watch()

	// block watcher
	for {
		v, err := watcher.Next()
		if err != nil {
			return
		}

		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(string(conf.Bytes()), string(v.Bytes()), false)
		spew.Dump(diffs)

		//log.Infof("[loadConfigFile] file change? %s", string(v.Bytes()))
	}
}
