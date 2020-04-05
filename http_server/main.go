package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"../sum"
	"net/http"
	"os"
)

const (
	SumEndpoint = "localhost:9090"
)

var err error
var bearer string

func run() error {
	log.Info("Starting...")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux, err := NewMux(ctx)
	if err != nil {
		return err
	}
	log.Info()
	server := NewServer(mux)
	return server.ListenAndServe()
}

func NewServer(mux *runtime.ServeMux) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: logRequest(mux),
	}
}

func NewMux(ctx context.Context) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			addAuthHeadersInterceptor,
			addHeadersInterceptor,
		)),
	}

	err = sum.RegisterSumComputerHandlerFromEndpoint(ctx, mux, SumEndpoint, opts)
	log.Info("Trying to register sum")
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func getEnvs() {
	bearer = os.Getenv("BEARER")
	if bearer == "" {
		log.Error("Can not find BEARER env.")
	}
}

func main() {
	getEnvs()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func addAuthHeadersInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "Bearer", bearer)
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func addHeadersInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "service", "serviceHeaderValue")
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
