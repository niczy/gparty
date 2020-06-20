package party

import (
	"context"
	"log"
	"testing"
	"time"
)

func addUser(userName string, profileImg string) *UserState {
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
	return response.UserState
}

func moveUser(userId string, newX, newY int64) *Position {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &MoveUserRequest{
		UserId: userId,
		NewPos: &Position{
			X: newX,
			Y: newY,
		},
	}
	response, err := client.MoveUser(ctx, req)
	if err != nil {
		log.Fatalf("fail to move user.")
	}
	return response.Pos
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

func TestMoveUser(t *testing.T) {
	userState := addUser("Nicholas", "https://www.google.com")
	pos := moveUser(userState.UserId, 2, 3)
	if pos.X != 2 || pos.Y != 3 {
		log.Fatalf("Fail to move user")
	}
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
