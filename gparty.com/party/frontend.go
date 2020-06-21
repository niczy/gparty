package party

import (
	"context"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(), 10*time.Second)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func dispatchRequest(w http.ResponseWriter,
			r *http.Request,
			f func(context.Context) (proto.Message, error)) {
		ctx, cancel := newContext()
		defer cancel()
		response, err := f(ctx)
		if err != nil {
			log.Fatalf("Failed to call backend")
		}
		js, err := protojson.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	return
}

func getUserStatesHandler(client PartyClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dispatchRequest(w, r, func(ctx context.Context) (proto.Message, error) {
			request := &GetUserStatesRequest{}
			response, err := client.GetUserStates(ctx, request)
			return response, err
		})
	})
}

func moveUserHandler(client PartyClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dispatchRequest(w, r, func(ctx context.Context) (proto.Message, error) {
			request := &MoveUserRequest{}
			response, err := client.MoveUser(ctx, request)
			return response, err
		})
	})
}

func addNewUserHandler(client PartyClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dispatchRequest(w, r, func(ctx context.Context) (proto.Message, error) {
			body, err := ioutil.ReadAll(r.Body)
			request := &AddNewUserRequest{}
			err = protojson.Unmarshal(body, request)
			response, err := client.AddNewUser(ctx, request)
			return response, err
		})
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
