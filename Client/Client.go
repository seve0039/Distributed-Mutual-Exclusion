package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	gRPC "github.com/seve0039/Distributed-Mutual-Exclusion.git/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	gRPC.TokenRingServer
	//participants     map[string]
	participantMutex sync.RWMutex
	name             string
	port             string
	lamportClock     int64
}

var client gPRC.tokenRingClient
var clientconn *grpc.ClientConn

func main() {
	// Connect to the clients
	sendConnectRequest()

	// Listen for connections from other clients
	launchConnection()

	// Listen for messages from other clients
	go listenForBroadcast(stream)

}

func sendConnectRequest() {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	clientconn, err = gRPC.Dial("server_address:port", nsecure.NewCredentials())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	fmt.Println("Sending Connect Request")
}

func launchConnection() {
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *port, err)
	}

	grpcServer := grpc.NewServer()
	server := &Server{
		name:         *serverName,
		port:         *port,
		participants: make(map[string]gRPC.ChittyChat_BroadcastServer),
	}

	gRPC.RegisterChittyChatServer(grpcServer, server)
	log.Printf("NEW SESSION: Server %s: Listening at %v\n", *serverName, list.Addr())

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
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
