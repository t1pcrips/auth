package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	desc "github.com/t1pcrips/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addres = "localhost:50051"
	ID     = 1223
)

func main() {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf(color.RedString("failed to connect to server - %v", addres)))
	}

	c := desc.NewUserV1Client(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: ID})

	if err != nil {
		log.Fatalf(fmt.Sprintf(color.RedString("failed to create Request to server: %v", err)))
	} else {
		log.Print(fmt.Sprintf(color.HiMagentaString("Request Created!")))
	}

	log.Printf(fmt.Sprintf(color.RedString("Note Info:\n") + color.GreenString("%+v", r.GetUser())))
}
