package user

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"woah/config"

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
func Apply(ctx context.Context, cmd string, IC config.IConfig) {
	// run service if check is true
	if !strings.Contains(cmd, appName) {
		return
	}
	//spew.Dump(string(IC.Get()))

	//d := IC.Get("services", "user")
	//spew.Dump(string(d))

	ts := strings.Split(cmd, ":")

	if len(ts) == 2 {
		port = fmt.Sprintf(":%s", ts[1])
	}

	router := gin.Default()

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	//
	gracehttp.Serve(srv)

}
