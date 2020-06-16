package party
 import (
	 "context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
 )

func newContext() (context.Context, context.CancelFunc) {
        return context.WithTimeout(
		context.Background(), 10*time.Second)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func getPartyMap(client PartyClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := newContext()
		defer cancel()
		request := GetPartyMapRequest{}
		partyMap, err := client.GetPartyMap(ctx, &request)
		if err != nil {
			log.Fatalf("%v.GetPartyMap(_) = _, %v", client, err)
		}

		fmt.Fprintf(w, "GetPartyResponse: %v", partyMap)
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
	mux.Handle("/_/getPartyMap", getPartyMap(client))
	log.Fatal(http.ListenAndServe(":8060", mux))
}
