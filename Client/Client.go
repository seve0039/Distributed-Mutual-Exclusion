package main

import (
	"flag"
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

var clientsName = flag.String("name", "user", "Sender's name")
var serverPort = flag.String("server", "5400", "Tcp server")

type Client struct {
	gRPC.UnimplementedTokenRingServer
	participants     map[string]gRPC.TokenRing_RequestCriticalSectionServer
	participantMutex sync.RWMutex
	clientName       string
	clientPort       string
	lamportClock     int64
}

var client gRPC.TokenRingClient
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

	conn, err := grpc.Dial(fmt.Sprintf(":%s", *serverPort), opts...)
	if err != nil {
		log.Fatalf("Fail to Dial : %v", err)
	}

	server = gRPC.NewTokenRingClient(conn)
	ServerConn = conn
}

func launchConnection() {
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *port, err)
	}

	grpcServer := grpc.NewServer()
	client := &Client{
		clientName:   *clientsName,
		clientPort:   *clientPort,
		participants: make(map[string]gRPC.TokenRing_RequestCriticalSectionServer),
	}

	gRPC.RegisterChittyChatServer(grpcServer, client)
	log.Printf("NEW SESSION: Server %s: Listening at %v\n", *clientName, list.Addr())

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
