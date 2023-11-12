package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	grpc "github.com/seve0039/Distributed-Mutual-Exclusion/tree/main/proto"
	"google.golang.org/grpc"
)

var client grpc.Client
var clientconn grpc.ClientConn

func main() {
	// Connect to the clients
	sendConnectRequest()

	// Listen for connections from other clients
	listenForConnection()

	stream, err := client.Broadcast(context.Background())
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}
	// Listen for messages from other clients
	go listenForBroadcast(stream)

}

func sendConnectRequest() {
	fmt.Println("Sending Connect Request")
}

func listenForConnection() {
	// Listen for connections from other clients
	fmt.Println("Listening for connections")
}

func EnterCriticalSection() {
	fmt.Println("Entered CriticalSection")
}

func listenForBroadcast() {
	fmt.Println("Listening for message")
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
