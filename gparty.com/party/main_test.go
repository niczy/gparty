package party

import (
	"log"
	"os"
	"testing"
	"time"
	"net/http"

	"google.golang.org/grpc"
)

var (
	client PartyClient

	httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
)

func TestMain(m *testing.M) {
	go StartBackendServer()
	go StartFrontendServer()
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


