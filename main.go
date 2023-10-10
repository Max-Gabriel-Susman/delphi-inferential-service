package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	pb "github.com/Max-Gabriel-Susman/delphi-inferential-service/inference"
	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/handler"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
	defaultName       = "world"
)

var (
	// addr = flag.String("addr", "10.96.0.3:50052", "the address to connect to")
	// addr = flag.String("addr", "10.100.0.3:50052", "the address to connect to")
	addr = flag.String("addr", "localhost:50053", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	port = flag.Int("port", 50054, "The server port") // actual port dictation
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitCodeErr)
	}
}

func run(ctx context.Context, _ []string) error {
	// var cfg struct {
	// 	ServiceName string `env:"SERVICE_NAME" envDefault:"delphi-inferential-service"`
	// 	Env         string `env:"ENV" envDefault:"local"`
	// 	Database    struct {
	// 		User   string `env:"INFERENTIAL_DB_USER,required"`
	// 		Pass   string `env:"INFERENTIAL_DB_PASSWORD,required"`
	// 		Host   string `env:"INFERENTIAL_DB_HOST"`
	// 		Port   string `env:"INFERENTIAL_DB_PORT" envDefault:"3306"`
	// 		DBName string `env:"INFERENTIAL_DB_Name" envDefault:"identity"`
	// 		Params string `env:"INFERENTIAL_DB_Param_Overrides" envDefault:"parseTime=true"`
	// 	}
	// 	Datadog struct {
	// 		Disable bool `env:"DD_DISABLE"`
	// 	}
	// 	Migration struct {
	// 		Enable bool `env:"ENABLE_MIGRATE"`
	// 	}
	// }
	// if err := env.Parse(&cfg); err != nil {
	// 	return errors.Wrap(err, "parsing configuration")
	// }

	// create grpc connection
	// conn, err := grpc.Dial("localhost:9092")
	// if err != nil {
	// 	panic(err)
	// }

	// create text generation client
	// tgc := textgeneration.NewTextGenerationServiceClient(conn)
	// _ = textgeneration.NewTextGenerationServiceClient(conn)

	// h := handler.API(handler.Deps{}, tgc) // needs implementation l8r
	h := handler.API(handler.Deps{})

	// Start HTTP Service
	api := http.Server{
		Handler: h,
		// Addr:              "127.0.0.1:80",
		Addr:              "0.0.0.0:8082",
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Start GRPC Service
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Make a channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start listening for requests
	go func() {
		// log info about this
		serverErrors <- api.ListenAndServe()
	}()
	// Shutdown

	// logic for handling shutdown gracefully
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case <-ctx.Done():
		// log something

		// request a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
	// return &pb.HelloReply{Message: "Hello world"}, nil
}

// Decode implements textgeneration.GreeterServer
func (s *server) Decode(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	// // Set up a connection to the server.
	// conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// c := pb.NewGreeterClient(conn)

	// // Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetMessage())

	// return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
	return &pb.HelloReply{Message: "Hello world"}, nil
}
