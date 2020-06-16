package party

import (
	"context"
	"net"
	"log"

	"google.golang.org/grpc"
)

type partyServer struct {
	UnimplementedPartyServer
}

func (s *partyServer) GetPartyMap(ctx context.Context, req *GetPartyMapRequest) (*PartyMap, error) {
	return &PartyMap{}, nil
}

func newServer() *partyServer {
	s := &partyServer{}
	return s
}

func StartServer() {
	lis, err := net.Listen("tcp", "localhost:9960")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	RegisterPartyServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

