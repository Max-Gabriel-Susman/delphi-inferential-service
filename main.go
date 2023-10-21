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
	// h := handler.API(handler.Deps{})

	// Start HTTP Service
	// api := http.Server{
	// 	Handler: h,
	// 	// Addr:              "127.0.0.1:80",
	// 	Addr:              "0.0.0.0:8082",
	// 	ReadHeaderTimeout: 2 * time.Second,
	// }

	// Start GRPC Service
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tgs := tg.NewTextGenerationServer()
	pb.RegisterGreeterServer(s, &tgs.Server)
	// pb.RegisterGreeterServer(s, tg.Server{})
	log.Printf("server listening at %v", lis.Addr())

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Make a channel to listen for errors coming from the listener
	// serverErrors := make(chan error, 1)

	// // Start listening for requests
	// go func() {
	// 	// log info about this
	// 	serverErrors <- api.ListenAndServe()
	// }()
	// // Shutdown

	// // logic for handling shutdown gracefully
	// select {
	// case err := <-serverErrors:
	// 	return errors.Wrap(err, "server error")

	// case <-ctx.Done():
	// 	// log something

	// 	// request a deadline for completion
	// 	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// 	defer cancel()

	// 	if err := api.Shutdown(ctx); err != nil {
	// 		api.Close()
	// 		return errors.Wrap(err, "could not stop server gracefully")
	// 	}
	// }

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return nil
}
