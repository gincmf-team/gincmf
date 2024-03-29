package main

import (
	"flag"
	"fmt"
	"gincmf/common/bootstrap/data"
	"net/http"
	"strings"

	"gincmf/service/portal/api/internal/config"
	"gincmf/service/portal/api/internal/handler"
	"gincmf/service/portal/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/portal.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 初始化
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 获取请求头域名
			scheme := "http://"
			if r.Header.Get("Scheme") == "https" {
				scheme = "https://"
			}
			host := r.Host
			domain := scheme + host
			ctx.Config.App.Domain = domain
			data.SetDomain(domain)
			ctx.Request = r
			// 处理userId
			r.ParseForm()
			userId := strings.Join(r.Form["userId"], "")
			ctx.Set("userId", userId)
			next(w, r)
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
