package textgeneration

import (
	"context"
	"flag"

	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/clients/openai"
	pb "github.com/Max-Gabriel-Susman/delphi-inferential-service/textgeneration"
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

type server struct {
	pb.UnimplementedGreeterServer
	OpenAIClient *openai.Client
}

type TextGenerationServer struct {
	Server server
}

func NewTextGenerationServer(openaiClient *openai.Client) *TextGenerationServer {
	return &TextGenerationServer{
		Server: server{
			OpenAIClient: openaiClient,
		},
	}
}
