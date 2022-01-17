package user

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"woah/config"
	"woah/pkg/broadcast"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

const (
	// AppName -
	appName = "user"
)

var (
	port string = ""
)

// Apply -
func Apply(ctx context.Context, cmd string, IC config.IConfig, b broadcast.Broadcast) {
	// run service if check is true
	if !strings.Contains(cmd, appName) {
		return
	}

	ts := strings.Split(cmd, ":")

	if len(ts) == 2 {
		port = fmt.Sprintf(":%s", ts[1])
	}

	// sub broadcast
	b.Sub(fmt.Sprintf("conf-%s", appName), func(d []byte) {
		fmt.Println(string(d))
	})

	router := gin.Default()

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	//
	gracehttp.Serve(srv)

}
