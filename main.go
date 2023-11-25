package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/clients/openai"
	tg "github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/textgeneration"
	pb "github.com/Max-Gabriel-Susman/delphi-inferential-service/textgeneration"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

var port = flag.Int("port", 50054, "The server port") // actual port dictation

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
	apiKey := os.Getenv("api-key") // we'll want to get from SSM later
	organization := os.Getenv("api-org")
	openaiClient := openai.NewClient(apiKey, organization)

	// Start GRPC Service
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tgs := tg.NewTextGenerationServer(openaiClient)
	pb.RegisterGreeterServer(s, &tgs.Server)
	log.Printf("server listening at %v", lis.Addr())

	// register the reflection service which allows clients to determine the methods for this gRPC service
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return nil
}
