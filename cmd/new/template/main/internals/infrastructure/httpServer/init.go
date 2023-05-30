package httpServer

import (
	"context"
	"net/http"

	"{{ service_name }}/internals/config"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	Config  config.Configuration
	Server  *runtime.ServeMux
	HttpMux *http.ServeMux
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {
}

func NewServer(
	config config.Configuration,
	rmux *runtime.ServeMux,
	httpMux *http.ServeMux,
) *Server {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s := &Server{
		Config:  config,
		Server:  rmux,
		HttpMux: httpMux,
	}

	s.Configure(context.Background(), opts)
	return s
}
