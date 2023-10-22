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

/*
	TODOs:
		META:
			* start a documentation direcory
			* start implementing testing coverage
			* work more on readme
			* abstract what we can to delphi-go-kit (e.g. logging, tracing, etc.)
			* determine what logging tracing solutions I want to use long term(probably just something within aws honestly)
			* refactor rootlevel protobuf/grpc logic into corresponding
				internal directories
			* refactor main.go to cmd/delphi-x-service/main.go
			* clean up Make targets and keep them up to date
			* abstract build logic execution into submodule delphi build-utils
			* we may want to drop the db directory, not sure if it really belongs in the
				current iteration of this service anymore
		MESA:
*/

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
	// awsCfg, err := aws.NewConfig(ctx)
	// if err != nil {
	// 	return errors.Wrap(err, "could not create aws sdk config")
	// }

	// var cfg struct {
	// 	OpenAI struct {
	// 		APIKey string `env:""`
	// 		APIOrg string ``
	// 	}
	// }
	// if err := env.Parse(&cfg); err != nil {
	// 	return errors.Wrap(err, "parsing configuration")
	// }
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

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return nil
}

// func Sessions() (*session.Session, error) {
// 	sess, err := session.NewSession()
// 	svc := session.Must(sess, err)
// 	return svc, err
// }

// func NewSSMClient() *SSM {
// 	// Create AWS Session
// 	sess, err := Sessions()
// 	if err != nil {
// 		log.Println(err)
// 		return nil
// 	}
// 	ssmsvc := &SSM{ssm.New(sess)}
// 	// Return SSM client
// 	return ssmsvc
// }

// type Param struct {
// 	Name           string
// 	WithDecryption bool
// 	ssmsvc         *SSM
// }

// //Param creates the struct for querying the param store
// func (s *SSM) Param(name string, decryption bool) *Param {
// 	return &Param{
// 		Name:           name,
// 		WithDecryption: decryption,
// 		ssmsvc:         s,
// 	}
// }

// func (p *Param) GetValue() (string, error) {
// 	ssmsvc := p.ssmsvc.client
// 	parameter, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
// 		Name:           &p.Name,
// 		WithDecryption: &p.WithDecryption,
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	value := *parameter.Parameter.Value
// 	return value, nil
// }

// // ssmsvc := NewSSMClient()
// // apiKey, err := ssmsvc.Param("myparam", true).GetValue()
// // if err != nil {
// // 	log.Println(err)
// // }
// // log.Println(apiKey)
// // organization, err := ssmsvc.Param("myparam", true).GetValue()
// // if err != nil {
// // 	log.Println(err)
// // }
// // log.Println(organization)
