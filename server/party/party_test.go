package party

import (
	"context"
	"log"
	"time"
	"testing"

	"google.golang.org/grpc"
)

func TestServerStartup(t *testing.T) {
	go StartBackendServer()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:9960", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := NewPartyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := GetPartyMapRequest{}
	partyMap, err := client.GetPartyMap(ctx, &request)
	if err != nil {
		log.Fatalf("%v.GetPartyMap(_) = _, %v", client, err)
	}
	if len(partyMap.Grid) != 1 {
		log.Fatalf("returned grid size is not 1.")
	}
}
