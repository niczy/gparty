package party

import (
	"context"
	"net"
	"log"

	guuid "github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	users = make(map[string]*UserState)
)

type partyServer struct {
	UnimplementedPartyServer }

func (s *partyServer) GetUserStates(ctx context.Context, req *GetUserStatesRequest) (*GetUserStatesResponse, error) {
	userStates := make([]*UserState, 0, len(users))
	for _, userState := range users {
		userStates = append(userStates, userState)
	}
	return &GetUserStatesResponse{UserStates: userStates}, nil
}

func (s *partyServer) AddNewUser(ctx context.Context, req *AddNewUserRequest) (*AddNewUserResponse, error) {
	userId := guuid.New().String()
	position := &Position{X: 1, Y: 2}
	userState := UserState{
		UserId: userId,
		ProfileImg: req.ProfileImg,
		UserName: req.UserName,
		Pos: position,
	}
	users[userId] = &userState
	return &AddNewUserResponse{UserState: &userState}, nil
}

func newServer() *partyServer {
	s := &partyServer{}
	return s
}

func Reset() {
	users = make(map[string]*UserState)
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



