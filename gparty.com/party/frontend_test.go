package party

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func postRequest(req proto.Message, protoRes proto.Message, path string) {
	bodyBytes, _ := protojson.Marshal(req)
	res, err := httpClient.Post("http://localhost:8060/_/"+path,
		"application/json", bytes.NewBuffer(bodyBytes))
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Read addNewUser resp failed, err: %v", err)
	}
	res.Body.Close()
	protojson.Unmarshal(resBody, protoRes)
}

func httpAddNewUser() *UserState {
	req := &AddNewUserRequest{
		UserName:   "Nicholas",
		ProfileImg: "https://wwww.google.com",
	}
	protoRes := &AddNewUserResponse{}
	postRequest(req, protoRes, "addNewUser")
	return protoRes.UserState
}

func httpMoveUser(userId string, x, y int64) *Position {
	req := &MoveUserRequest{
		UserId: userId,
		NewPos: &Position{X: x, Y: y},
	}
	protoRes := &MoveUserResponse{}
	postRequest(req, protoRes, "moveUser")
	return protoRes.Pos
}

func httpGetUserStates() []*UserState {
	res, err := httpClient.Get("http://localhost:8060/_/getUserStates")
	if err != nil {
		log.Fatalf("Failed to call getUserStates, erro %v", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	dataResp := GetUserStatesResponse{}
	protojson.Unmarshal(body, &dataResp)
	return dataResp.UserStates
}

func TestGetUserStatesHandler(t *testing.T) {
	Reset()
	httpAddNewUser()
	userStates := httpGetUserStates()
	if len(userStates) != 1 {
		log.Fatalf("Should return 1 UserState.")
	}
}

func TestAddNewUserHandler(t *testing.T) {
	Reset()
	httpAddNewUser()
}

func TestMoveUserHandler(t *testing.T) {
	Reset()
	userState := httpAddNewUser()
	pos := httpMoveUser(userState.UserId, 3, 4)
	if pos.X != 3 || pos.Y != 4 {
		log.Fatalf("Failed to move user %s.", userState.UserId)
	}
}
