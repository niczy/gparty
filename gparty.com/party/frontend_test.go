package party

import (
	"io/ioutil"
	"log"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestGetUserStatesHandler(t *testing.T) {
	res, err := httpClient.Get("http://localhost:8060/_/getUserStates")
	if err != nil {
		log.Fatalf("Failed to call getUserStates, erro %v", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	// bodyString := string(body)
	// log.Printf("Read body %v", bodyString)
	res.Body.Close()
	dataResp := GetUserStatesRequest{}
	// TODO: Fix unmarshal
	_ = protojson.Unmarshal(body, &dataResp)
	// log.Printf("resp is %v", dataResp)
}

func TestAddNewUserHandler(t *testing.T) {
}

func TestMoveUserHandler(t *testing.T) {
}
