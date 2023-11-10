package textgeneration

import (
	"context"
	"fmt"
	"log"

	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/clients/openai"
	pb "github.com/Max-Gabriel-Susman/delphi-inferential-service/textgeneration"
)

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("guys it's me")
	log.Printf("Received: %v", in.GetName())
	r := openai.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.Message{
			{
				Role:    "user",
				Content: in.GetName(),
			},
		},
		Temperature: 0.7,
	}

	completions, err := s.OpenAIClient.CreateCompletions(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(completions)

	reply := &pb.HelloReply{Message: completions.Choices[0].Message.Content}

	return reply, nil
}
