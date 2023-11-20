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
	clientId = *clientPort

	go startServer()
	listenForOtherClient()
	defer serverConn.Close()

	go handleCommands()
	//client
	//requestCriticalSection(int32(*clientPort), stream)
	//server
	for {
	}

}

func handleCommands() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter request: ")
	for reader.Scan() {
		fmt.Println("Enter request: ")
		msg := reader.Text()

		fmt.Println(int32(clientId))

		if msg == "request-cs" {
			_, _ = server.SendRequestAccess(context.Background(), &gRPC.CriticalSectionRequest{
				NodeId: int32(clientId), Denied: false,
			})

		}

	}
}

// Client
func listenForOtherClient() {

	fmt.Println("Connecting to server...")
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
		fmt.Println("Client id:", clientId, "Send id:", criticalSectionRequest.NodeId)
		enterCriticalSection()
		return &emptypb.Empty{}, nil
	} else if criticalSectionRequest.NodeId == int32(clientId) && criticalSectionRequest.Denied {
		fmt.Println("Desværre du!")
		return &emptypb.Empty{}, nil
	}
	fmt.Printf("A request is recieved from sender: %v", criticalSectionRequest.NodeId)
	fmt.Println()

	//Checks if it is in critical section or denied and denies request
	if inCriticalSection {
		criticalSectionRequest.Denied = true
		_, _ = server.SendRequestAccess(context, criticalSectionRequest)
		fmt.Println("Access denied: ", criticalSectionRequest.NodeId)
		return &emptypb.Empty{}, nil
	} else {
		_, _ = server.SendRequestAccess(context, criticalSectionRequest)
		fmt.Println("Access granted: ", criticalSectionRequest.NodeId)
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

	grpcServer.Serve(list)

	fmt.Println("Server started!")

}

/*func checkMessageId(id int32, stream gRPC.TokenRing_RequestCriticalSectionClient) {

	if clientId != int(id) {
		requestCriticalSection(id, stream)
	}
}*/

func enterCriticalSection() {
	fmt.Println("Entered CriticalSection")
	inCriticalSection = true
	// A client will be in the critical section for 5 seconds before exiting
	go func() {
		time.Sleep(10 * time.Second) // Wait for 5 seconds
		mu.Lock()
		inCriticalSection = false
		fmt.Println("Exited CriticalSection")
		mu.Unlock()
	}()

}
