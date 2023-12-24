package pprofmiddleware

import (
	"context"
	"os"
	"runtime/pprof"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a new unary server interceptors that performs per-request logging.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		f, _ := os.Create(uuid.New().String() + "heap.out")
		resp, err = handler(ctx, req)
		pprof.Lookup("heap").WriteTo(f, 0)
		f.Close()
		return resp, err
	}
}
