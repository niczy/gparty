package party

import (
	"context"
	"log"
	"time"
	"testing"
	"os"

	"google.golang.org/grpc"
)

var (
	client PartyClient
)
func TestMain(m *testing.M) {
	go StartBackendServer()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:9960", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client = NewPartyClient(conn)
	log.Println("start running tests.")
	os.Exit(m.Run())
}

func TestAddNewUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := AddNewUserRequest{
		UserName: "Nicholas Zhao",
		ProfileImg: "https://www.google.com",
	}
	response, err := client.AddNewUser(ctx, &request)
	if (err != nil) {
		log.Fatalf("AddNewUser error %v", err)
	}
	if (response.UserState.UserId != "uid") {
		log.Fatalf("returned add new user response fail.")
	}

}

func TestGetUserStates(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := GetUserStatesRequest{}
	response, err := client.GetUserStates(ctx, &request)
	if err != nil {
		log.Fatalf("%v.GetPartyMap(_) = _, %v", client, err)
	}
	if len(response.UserStates) != 1 {
		log.Fatalf("returned grid size is not 1.")
	}
}
