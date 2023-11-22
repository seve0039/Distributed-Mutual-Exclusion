package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	gRPC "github.com/seve0039/Distributed-Mutual-Exclusion.git/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var isSending = false
var clientId = 0
var max = 2
var nextClient string
var inCriticalSection bool
var mu sync.Mutex
var server gRPC.TokenRingClient
var serverConn *grpc.ClientConn
var clientsName = flag.String("name", "default", "Client's name")
var clientPort = flag.Int("server", 5400, "Tcp server")

type Client struct {
	gRPC.UnimplementedTokenRingServer
	name string
	port int
}

func main() {
	flag.Parse()
	createLogFile()
	clientId = *clientPort

	go startServer()
	listenForOtherClient()
	defer serverConn.Close()

	go handleCommands()

	for {}

}

func handleCommands() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter request: ")
	for reader.Scan() {
		fmt.Println("Enter request: ")
		msg := reader.Text()

		if msg == "request-cs" {
			isSending = true
			_, _ = server.SendRequestAccess(context.Background(), &gRPC.CriticalSectionRequest{
				NodeId: int32(clientId), Denied: false,
			})
		}

	}

}

// Client
func listenForOtherClient() {

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	nextPort := *clientPort

	if nextPort >= 5400+max {
		nextPort = 5400
	} else {
		nextPort++
	}

	sPort := strconv.FormatInt(int64(nextPort), 10)

	fmt.Println("Connect request to port:", sPort)

	conn, err := grpc.Dial(fmt.Sprintf(":%s", sPort), opts...)
	fmt.Println("Connected to port:", nextPort)
	if err != nil {
		log.Fatalf("Fail to Dial : %v", err)
	}

	// Move these lines outside of the error handling block
	server = gRPC.NewTokenRingClient(conn)
	serverConn = conn

}

func (s *Client) SendRequestAccess(context context.Context, criticalSectionRequest *gRPC.CriticalSectionRequest) (*emptypb.Empty, error) {
	//The client recieves its own request to enter critical section
	if criticalSectionRequest.NodeId == int32(clientId) && !criticalSectionRequest.Denied {
		enterCriticalSection()
		return &emptypb.Empty{}, nil
	} else if criticalSectionRequest.NodeId == int32(clientId) && criticalSectionRequest.Denied {
		fmt.Println("Access denied by client antoher client. Try again later")
		isSending = false
		return &emptypb.Empty{}, nil
	}
	fmt.Printf("A request is recieved from sender: %v", criticalSectionRequest.NodeId)
	fmt.Println()

	//Checks if it is in critical section or denied and denies request
	if inCriticalSection || (isSending && criticalSectionRequest.NodeId < int32(clientId)) {
		criticalSectionRequest.Denied = true
		_, _ = server.SendRequestAccess(context, criticalSectionRequest)
		log.Println(clientId, ": Denied access to: ", criticalSectionRequest.NodeId) //Log
		fmt.Println("Denied access to: ", criticalSectionRequest.NodeId)
		return &emptypb.Empty{}, nil
	} else {
		_, _ = server.SendRequestAccess(context, criticalSectionRequest)
		log.Println(clientId, ": Granted acces to", criticalSectionRequest.NodeId) //Log
		fmt.Println("Granted acces to", criticalSectionRequest.NodeId)
		return &emptypb.Empty{}, nil
	}
}

// Server
func startServer() {
	prevPort := *clientPort
	sPort := strconv.FormatInt(int64(prevPort), 10)

	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", sPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", sPort, err)
	}

	grpcServer := grpc.NewServer()
	server := &Client{
		name: *clientsName,
		port: *clientPort,
	}

	gRPC.RegisterTokenRingServer(grpcServer, server)
	log.Printf("NEW SESSION: Server %s: Listening at %v\n", *clientsName, list.Addr())
	log.Println(clientId, ": has joined the session")

	grpcServer.Serve(list)

	log.Println("Server started!")

}


func enterCriticalSection() {
	log.Println(clientId, ": Entered CriticalSection")
	fmt.Println("Entering critical section")
	isSending = false
	inCriticalSection = true
	// A client will be in the critical section for 5 seconds before exiting
	go func() {
		time.Sleep(10 * time.Second) // Wait for 5 seconds
		mu.Lock()
		inCriticalSection = false
		log.Println(clientId, ": Left the CriticalSection")
		fmt.Println("Left the critical section")
		mu.Unlock()

	}()

}
func createLogFile() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}
