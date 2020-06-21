package party

import (
	"io/ioutil"
	"log"
	"bytes"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

func httpAddNewUser() {
	req := &AddNewUserRequest{
		UserName: "Nicholas",
		ProfileImg: "https://wwww.google.com",
	}
	bodyBytes, _ := protojson.Marshal(req)
	res, err := httpClient.Post("http://localhost:8060/_/addNewUser",
		"application/json", bytes.NewBuffer(bodyBytes))
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Read addNewUser resp failed, err: %v", err)
	}
	res.Body.Close()
	dataResp := &AddNewUserResponse{}
	protojson.Unmarshal(resBody, dataResp)
	if dataResp.UserState == nil {
		log.Fatalf("AddNewUser Failed.")
	}
}

func TestGetUserStatesHandler(t *testing.T) {
	Reset()
	httpAddNewUser()
	res, err := httpClient.Get("http://localhost:8060/_/getUserStates")
	if err != nil {
		log.Fatalf("Failed to call getUserStates, erro %v", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	dataResp := GetUserStatesResponse{}
	protojson.Unmarshal(body, &dataResp)
	if len(dataResp.UserStates) != 1 {
		log.Fatalf("Should return 1 UserState.")
	}
}

func TestAddNewUserHandler(t *testing.T) {
	Reset()
        httpAddNewUser()
}

func TestMoveUserHandler(t *testing.T) {
}
