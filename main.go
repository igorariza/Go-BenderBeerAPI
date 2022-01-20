package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	svcbeerv1alpha1 "github.com/igorariza/gogrpc-apibeer/cmd/beer"
	beerv1alpha1 "github.com/igorariza/gogrpc-apibeer/gen/go"
	v1 "github.com/igorariza/gogrpc-apibeer/transport/http/v1"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "50052"
	}
	// instance service
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	// instance grpc and http
	grpcBeerService := svcbeerv1alpha1.NewService(logger)

	endpoints := v1.MakeEndpoints(grpcBeerService)
	httpHandler := v1.NewHTTPHandler(endpoints, nil, nil, logger)

	var opts []grpc.ServerOption

	opts = append(opts, grpc.MaxSendMsgSize(1000*1024*1024))   // set max response size body ... by default is 4mb .... increased to 1000mb
	opts = append(opts, grpc.MaxRecvMsgSize(1000*1024*1024))   // set max receiver size body ... by default is 4mb .... increased to 1000mb
	opts = append(opts, grpc.MaxConcurrentStreams(4294967290)) // increase max streams active

	// up servers
	var g group.Group
	{
		grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", ":50051")
			// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
			// the here demonstrated zipkin tracing middleware.
			//baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			baseServer := grpc.NewServer(opts...)

			beerv1alpha1.RegisterBeerAPIServiceServer(baseServer, &grpcBeerService)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", ":3000")
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", ":3000")
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
	/*
		listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

		fmt.Println("API V1 CERVEZAS")
		fmt.Println("CON MICROSERVICIOS INCLUIDOS ")

		if err != nil {
			fmt.Println("Error al iniciar el servidor")
		}

		DB := data.New()
		if err := DB.DB.Ping(); err != nil {
			fmt.Println("Error al conectar a la base de datos")
		}

		if err != nil {
			fmt.Println("Error cannot create tcp connection ", err)
		}
		fmt.Println("Connection established running on port ", port)

		ser := grpc.NewServer(opts...)
		beerv1alpha1.RegisterBeerAPIServiceServer(ser, &grpcBeerService)

		if err := ser.Serve(listen); err != nil {
			fmt.Println("Error cannot create tcp connection ", err)
		}

		defer data.Close() */
}
