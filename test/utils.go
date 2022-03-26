// +acceptance
package test

import (
	"log"

	"github.com/goethesum/protos/rocket/v1"
	"google.golang.org/grpc"
)

func GetClient() rocket.RocketServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	rocketClinet := rocket.NewRocketServiceClient(conn)
	return rocketClinet
}
