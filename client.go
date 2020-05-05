package main

import (
	"github.com/ichtrojan/grpc-server/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	chatService := chat.NewChatServiceClient(conn)

	response, err := chatService.SayHello(context.Background(), &chat.Message{
		Body: "Trojan ⚠️",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Message: %s", response.Body)
}
