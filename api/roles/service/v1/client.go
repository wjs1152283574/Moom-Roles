package v1

import (
	context "context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	g "google.golang.org/grpc"
)

func NewClient() (client RolesClient, err error) {
	// env := os.Getenv("ENV")
	var conn *g.ClientConn
	ctx := context.Background()
	conn, err = grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("127.0.0.1:9966"),
		grpc.WithMiddleware(
			//jwt
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte("serviceTestKey"), nil
			}),
			//日志中间件
			logging.Client(log.DefaultLogger),
		),
	)
	if err != nil {
		return
	}
	client = NewRolesClient(conn)
	return
}
