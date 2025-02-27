package helpers

import (
	"context"
	"log"
	"strings"
	"time"

	dgo "github.com/dgraph-io/dgo/v210"
	api "github.com/dgraph-io/dgo/v210/protos/api"
)              
  
type CancelFunc func()  
func GetDgraphClient() (*dgo.Dgraph, CancelFunc) {
	conn, err := dgo.DialCloud("https://blue-surf-600096.us-east-1.aws.cloud.dgraph.io/graphql", "OGM0Y2IxYmM5OTJhYmI4ZjY0YzJhODY2NzgwMzhlZTY=")
	if err != nil {
		log.Fatal("Connection error", err)
	}

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	ctx := context.Background()

	for {
		err = dg.Login(ctx, "groot", "password")
		if err == nil || !strings.Contains(err.Error(), "Please retry") {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatalf("While trying to login %v", err.Error())
	}
	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}