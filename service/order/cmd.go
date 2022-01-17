package order

import (
	"context"
	"fmt"
	"strings"
	"woah/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

const (
	// AppName -
	appName = "order"
)

var (
	port string = ""
)

// Apply -
func Apply(ctx context.Context, cmd string, IC config.IConfig) {

	// run service if check is true
	if !strings.Contains(cmd, appName) {
		return
	}

	ts := strings.Split(cmd, ":")

	if len(ts) == 2 {
		port = fmt.Sprintf(":%s", ts[1])
	}

	go func() {
		for {
			select {
			case v := <-IC.Watcher(context.Background(), appName):
				spew.Dump(v.Target, "apply")
			}
		}
	}()

	app := gin.New()

	app.Run(port)
}
