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
	"strings"
	"sync"

	gRPC "github.com/seve0039/Distributed-Mutual-Exclusion.git/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientId = 0
var max = 2
var nexClient string
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

	stream, err := server.RCS(context.Background())
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}
	go listenForMessage(stream)
	//client
	//requestCriticalSection(int32(*clientPort), stream)
	//server

	for {
	}
}

func handleCommands() {
	serverStream, err := server.RCS(context.Background())
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter request: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		if cmd == "request-cs" {
			requestCriticalSection(int32(*clientPort), serverStream)
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

func requestCriticalSection(ClientId int32, stream gRPC.TokenRing_RCSClient) {
	mu.Lock()
	defer mu.Unlock()

	if inCriticalSection {
		fmt.Println("Already in critical section")
		return
	}

	//Send request to the next client
	msg := &gRPC.CriticalSectionRequest{NodeId: int32(ClientId)}
	_, err := RequestCSHelper(context.Background(), msg)
	if err != nil {
		log.Println("Failed to request critical section: ", err)
		return
	}
	fmt.Println("Requested critical section access")
	stream.Send(msg)

}


func (s *Client) RequestCSHelper(ctx context.Context, req *gRPC.CriticalSectionRequest) (*gRPC.CriticalSectionRequest, error) {
	mu.Lock()
	defer mu.Unlock()

	if inCriticalSection || int(req.NodeId) == clientId {
		if int(req.NodeId) == clientId {
			inCriticalSection = true
			fmt.Println("Entered critical section")
		}
		return &gRPC.CriticalSectionRequest{}, nil
	}

	_, err := s.RequestCSHelper(ctx, req)
	if err != nil {
		log.Println("Failed to forward critical section request: ", err)
	}
	return &gRPC.CriticalSectionRequest{}, nil
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

func listenForMessage(stream gRPC.TokenRing_RCSClient) {

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Println("Failed to receive message: ", err)
			return
		}

		fmt.Println(msg.GetNodeId())
	}

}

/*func checkMessageId(id int32, stream gRPC.TokenRing_RCSClient) {

	if clientId != int(id) {
		requestCriticalSection(id, stream)
	}
}*/

func EnterCriticalSection() {
	fmt.Println("Entered CriticalSection")
}
