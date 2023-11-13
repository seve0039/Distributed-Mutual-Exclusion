package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	gPRC "github.com/seve0039/Distributed-Mutual-Exclusion/proto"
	"google.golang.org/grpc"
)

var client gPRC.TokenRingClient
var clientconn grpc.ClientConn

func main() {
	// Connect to the clients
	sendConnectRequest()

	// Listen for connections from other clients
	launchConnection()

	stream, err := client.(context.Background())
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}
	// Listen for messages from other clients
	go listenForBroadcast(stream)

}

func sendConnectRequest() {
	var err error
	clientconn, err = grpc.Dial("server_address:port", nsecure.NewCredentials())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	fmt.Println("Sending Connect Request")
}

func listenForConnection() {
	// Listen for connections from other clients
	fmt.Println("Listening for connections")
}

func EnterCriticalSection() {
	fmt.Println("Entered CriticalSection")
}

func listenForBroadcast(stream gRPC.TokenRingClient) {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println("Failed to receive broadcast: ", err)
				return
			}
	
			fmt.Println(msg.GetMessage())
	}
	
}

func requestCriticalSection() {
	fmt.Println("Requested CriticalSection")
}

func createLogFile() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}
