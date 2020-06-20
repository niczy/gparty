package party

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

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

func addUser(userName string, profileImg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := AddNewUserRequest{
		UserName:   "Nicholas Zhao",
		ProfileImg: "https://www.google.com",
	}
	response, err := client.AddNewUser(ctx, &request)
	if err != nil {
		log.Fatalf("AddNewUser error %v", err)
	}
	if response.UserState == nil {
		log.Fatalf("returned add new user response fail.")
	}
}

func getUserStates() []*UserState {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := GetUserStatesRequest{}
	response, err := client.GetUserStates(ctx, &request)
	if err != nil {
		log.Fatalf("%v.GetPartyMap(_) = _, %v", client, err)
	}
	return response.UserStates
}

func TestAddNewUser(t *testing.T) {
	addUser("Nicholas Zhao", "https://www.gogle.co")
	addUser("bla bal", "https://www.gogle.co")
}

func TestGetUserStates(t *testing.T) {
	Reset()
	userStates := getUserStates()
	if len(userStates) != 0 {
		log.Fatalf("User states len should be 0. actual userStates %v",
			userStates)
	}
	addUser("Nicholas Zhao", "https://www.google.com")
	addUser("blabla", "https://www.fb.com")
	userStates = getUserStates()
	if len(userStates) != 2 {
		log.Fatalf("User states len should be 2. Actual length %d, actual userStates %v", len(userStates), userStates)
	}
}
