package party

import (
	"context"
	"net"
	"log"

	"google.golang.org/grpc"
)

var (
	users = make(map[string]UserState)
)

type partyServer struct {
	UnimplementedPartyServer
}

func (s *partyServer) GetUserStates(ctx context.Context, req *GetUserStatesRequest) (*GetUserStatesResponse, error) {
	userState := UserState{ProfileImg: "/img/avatar-1.png"}
	return &GetUserStatesResponse{UserStates: []*UserState{&userState}}, nil
}

func (s *partyServer) AddNewUser(ctx context.Context, req *AddNewUserRequest) (*AddNewUserResponse, error) {
	userId := "uid"
	position := &Position{X: 1, Y: 2}
	userState := UserState{
		UserId: userId,
		ProfileImg: req.ProfileImg,
		UserName: req.UserName,
		Pos: position,
	}
	users[userId] = userState
	return &AddNewUserResponse{UserState: &userState}, nil
}

func newServer() *partyServer {
	s := &partyServer{}
	return s
}

func StartBackendServer() {
	lis, err := net.Listen("tcp", "localhost:9960")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	RegisterPartyServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}



