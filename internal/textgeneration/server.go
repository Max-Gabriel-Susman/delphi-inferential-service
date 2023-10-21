package textgeneration

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	pb "github.com/Max-Gabriel-Susman/delphi-inferential-service/inference"
	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/clients/openai"
)

const defaultName = "world"

var (
	// we need to parameterize and resolve these addr redundancies l8r
	// addr = flag.String("addr", "10.96.0.3:50052", "the address to connect to")
	// addr = flag.String("addr", "10.100.0.3:50052", "the address to connect to")
	addr = flag.String("addr", "localhost:50053", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type Server interface {
	SayHello(context.Context, *pb.HelloRequest) (*pb.HelloReply, error)
	Decode(context.Context, *pb.HelloRequest) (*pb.HelloReply, error)
}

type TextGenerationServer struct {
	Server server
}

type server struct {
	pb.UnimplementedGreeterServer
}

func NewTextGenerationServer() *TextGenerationServer {
	return &TextGenerationServer{
		Server: server{},
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
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

	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")

	client := openai.NewClient(apiKey, organization)

	r := openai.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.Message{
			{
				Role:    "user",
				Content: "Say this is a test!",
			},
		},
		Temperature: 0.7,
	}

	completions, err := client.CreateCompletions(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(completions)

	return &pb.HelloReply{Message: "Hello world"}, nil
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
