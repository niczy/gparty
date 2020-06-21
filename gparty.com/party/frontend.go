package party

import (
	"context"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(), 10*time.Second)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func getUserStatesHandler(client PartyClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := newContext()
		defer cancel()
		request := GetUserStatesRequest{}
		response, err := client.GetUserStates(ctx, &request)
		if err != nil {
			log.Fatalf("%v.GetUserStates(_) = _, %v", client, err)
		}
		js, err := protojson.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func moveUserHandler(client PartyClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := newContext()
		defer cancel()
		request := &MoveUserRequest{}
		response, err := client.MoveUser(ctx, request)
		if err != nil {
			log.Fatalf("%v.MoveUser(_) = _, %v", client, err)
		}

		fmt.Fprintf(w, "MoveUserResponse: %v", response)

	})
}

func addNewUserHandler(client PartyClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		request := &AddNewUserRequest{}
		err = protojson.Unmarshal(body, request)
		ctx, cancel := newContext()
		defer cancel()
		response, err := client.AddNewUser(ctx, request)
		if err != nil {
			log.Fatalf("%v.AddNewUser(_) = _, %v", client, err)
		}
		js, err := protojson.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func StartFrontendServer() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:9960", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := NewPartyClient(conn)

	fs := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.Handle("/_/moveUser", moveUserHandler(client))
	mux.Handle("/_/addNewUser", addNewUserHandler(client))
	mux.Handle("/_/getUserStates", getUserStatesHandler(client))
	log.Fatal(http.ListenAndServe(":8060", mux))
}
