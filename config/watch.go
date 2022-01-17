package config

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/bytedance/sonic"
	"go-micro.dev/v4/config"
)

var (
	// md - save all config key/value(md5)
	md map[string]string = make(map[string]string)

	// mux -
	mux sync.Mutex
)

type watcher struct{}

// watch - watch remote config data
// handle something when data changes ex: restart service etc.
func watch(ctx context.Context, conf config.Config, r *root) {
	watcher, _ := conf.Watch()

	// new md map
	md = newMap(r.Services)

	// block watcher
	for {
		v, err := watcher.Next()
		if err != nil {
			continue
		}

		// get services config
		node, _ := sonic.Get(v.Bytes(), SERVICES)

		// convert to raw(string)
		raw, _ := node.Raw()

		// unmarshal
		m := make(map[string]interface{})
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			continue
		}

		// find diff config
		df := diff(md, newMap(m))
		if df == nil {
			continue
		}

		// reset md value
		md[df.Target] = string(df.Value)

		// get target value
		n := node.Get(df.Target)
		deepRaw, _ := n.Raw()

		// print
		log.Printf("Config change : %v service, value: %v\n", df.Target, deepRaw)

		// send to channel
		//r.Ch <- Values{
		//Target: df.Target,
		//Value:  []byte(deepRaw),
		//}
	}
}

// newMap -
func newMap(m interface{}) map[string]string {
	mux.Lock()
	defer mux.Unlock()

	mi, ok := m.(map[string]interface{})

	result := make(map[string]string)
	if !ok {
		return result
	}

	for k, v := range mi {
		result[k] = genMD5(v)
	}

	return result
}

// genMD5 -
func genMD5(in interface{}) string {
	Md5Inst := md5.New()

	d, _ := json.Marshal(in)
	Md5Inst.Write(d)

	return fmt.Sprintf("%x", Md5Inst.Sum([]byte("")))
}

// diff -
func diff(old, new map[string]string) *Values {
	for k, v := range old {
		m, ok := new[k]
		if !ok {
			continue
		}

		if v != m {
			return &Values{
				Target: k,
				Value:  []byte(m),
			}
		}
	}

	return nil
}

// allowObject -
func allowObject(target string, allow ...string) bool {
	if len(allow) == 0 {
		return true
	}

	for _, v := range allow {
		if v == target {
			return true
		}
	}
	return false

}
