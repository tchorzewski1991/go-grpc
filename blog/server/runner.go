package main

import (
	"log"
	"google.golang.org/grpc"
	"net"
	"github.com/tchorzewski1991/go-grpc/blog/blogpb"
	"os"
	"os/signal"
	"golang.org/x/net/context"
	"github.com/tchorzewski1991/go-grpc/blog/helpers"
	"github.com/tchorzewski1991/go-grpc/blog/driver"
	blogRepo "github.com/tchorzewski1991/go-grpc/blog/repository/blog"
	grpcService "github.com/tchorzewski1991/go-grpc/blog/service/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
	"fmt"
	"time"
)

var dbClient *driver.DB

func main() {
	log.SetFlags(log.LstdFlags)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = setProperty(ctx, "MONGO_DB_HOST", "localhost")
	ctx = setProperty(ctx, "MONGO_DB_PORT", "27017")
	ctx = setProperty(ctx, "GRPC_SERVER_PORT", "50051")
	ctx = setProperty(ctx, "HTTP_SERVER_PORT", "8080")

	dbClient = driver.ConnectMongoDB(ctx)

	go runGrpc(ctx)
	go runHTTP(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<- stop
}

func runHTTP(ctx context.Context) {
	proxyPort := ctx.Value("HTTP_SERVER_PORT").(string)
	grpcPort  := ctx.Value("GRPC_SERVER_PORT").(string)

	proxyAddr := fmt.Sprintf("0.0.0.0:%v", proxyPort)
	grpcAddr  := fmt.Sprintf("0.0.0.0:%v", grpcPort)

	log.Printf("Starting HTTP proxy server on port: %v", proxyPort)

	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := blogpb.RegisterBlogServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts)
	helpers.LogIfError(err, "Failed to connect with gRPC server: %v")

	err = http.ListenAndServe(proxyAddr, withLogging(mux))
	helpers.LogIfError(err, "Failed to run proxy server: %v")

	defer func() {
		log.Println("Stopping HTTP Proxy server...")
	}()
}

func runGrpc(ctx context.Context) {
	port := ctx.Value("GRPC_SERVER_PORT").(string)
	addr := fmt.Sprintf("0.0.0.0:%v", port)

	log.Printf("Starting gRPC server on port: %v", port)

	listener, err := net.Listen("tcp", addr)
	helpers.LogIfError(err, "Failed to connect tcp: %v")

	server  := grpc.NewServer()
	repo    := blogRepo.NewMongoDbBlogRepo(dbClient.Mongo)
	service := grpcService.NewBlogService(repo)

	blogpb.RegisterBlogServiceServer(server, service)

	err = server.Serve(listener)
	helpers.LogIfError(err, "Failed to serve: %v")

	defer func() {
		log.Println("Stopping gRPC server...")
		server.Stop()
		listener.Close()
		dbClient.Mongo.Client().Disconnect(ctx)
	}()
}

func setProperty(ctx context.Context, lookupKey string, defaultValue string) context.Context {
	value, ok := os.LookupEnv(lookupKey)
	if !ok { value = defaultValue }
	return context.WithValue(ctx, lookupKey, value)
}

func withLogging(inner *runtime.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("HTTP | Method: %s\t%s\t%s", r.Method, r.RequestURI, time.Since(start))
	}
}
