package main

import (
	"context"
	"log"
	"net"

	identity "github.com/spotify/backstage/backend/proto/identity/v1"
	pb "github.com/spotify/backstage/backend/proto/scaffolder/v1"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterScaffolderServer(grpcServer, &server{})

	log.Println("Serving Scaffolder Service")
	grpcServer.Serve(lis)
}

// LocalTemplate json representation
type LocalTemplate struct {
	ID   string
	Name string
}

type server struct{}

func (s *server) GetAllTemplates(ctx context.Context, req *pb.Empty) (*pb.GetAllTemplatesResponse, error) {
	template := &pb.Template{
		Id:   "react-ssr-template",
		Name: "React SSR Template",
		User: &identity.User{
			Id:   "spotify",
			Name: "Spotify",
		},
	}

	templates := []*pb.Template{template}

	return &pb.GetAllTemplatesResponse{
		Templates: templates,
	}, nil
}
