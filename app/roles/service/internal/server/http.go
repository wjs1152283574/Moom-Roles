package server

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/app/roles/service/internal/service"
	"github.com/it-moom/moom-roles/pkg/errors"
	"github.com/it-moom/moom-roles/pkg/tool"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, s *service.RolesService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			AuthMiddleware,
			logging.Server(logger), // 添加全局日志中间件
			// logging.Client(logger),
			// ratelimit.Server(), // 启用过载保护（默认一个时间窗口 100 pass）

		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	// 服务内跨域处理
	// TODO: 引入网关后在网关处理跨域时，需要删除以下处理跨域的代码
	opts = append(opts, http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"DNT", "X-Mx-ReqToken", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Authorization", "udid", "appkey", "version", "authenticated", "cookie", "token"}),
		handlers.ExposedHeaders([]string{"DNT", "X-Mx-ReqToken", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Authorization", "udid", "appkey", "version", "authenticated", "cookie", "token"}),
		handlers.OptionStatusCode(204),
	)))

	// // 指定为json编码格式
	// opts = append(opts, http.ResponseEncoder(resencoder.ResponeJsonDeco()))
	srv := http.NewServer(opts...)
	v1.RegisterRolesHTTPServer(srv, s)

	return srv
}

// AuthMiddleware 网关服务会将userid添加到查询参数打到本服务。次中间件将userid 添加到上下文中
func AuthMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			ht, _ := tr.(*http.Transport)
			if !tool.InSlice(NotneedAuth, ht.Request().URL.Path) {
				tokenStr := strings.Split(ht.RequestHeader().Get("Authorization"), " ")
				if len(tokenStr) > 1 { // 存在token，解析
					tcliam, errs := tool.NewJWT(conf.GB.Global.TokenScrect).ParseToken(tokenStr[1])
					if errs != nil {
						return nil, errors.ErrAuthFail()
					}
					etxs := context.WithValue(ctx, "userid", tcliam.Subject)
					reply, err = handler(etxs, req)
				} else {
					return nil, errors.ErrAuthFail()
				}

			} else {
				reply, err = handler(ctx, req)
			}
		}

		return
	}
}
