package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
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
				NodeId: int32(clientId),
			})
			/*
				serverStream, err := server.RequestCriticalSection(context.Background())
				if err != nil {
					log.Println("Failed to send message:", err)
					continue
				}

				fmt.Println("line 70")

				requestCriticalSection(int32(*clientPort), serverStream)
			*/

		}
		//fmt.Print("Enter request: ")
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

func requestCriticalSection(ClientId int32, stream gRPC.TokenRing_RequestCriticalSectionClient) {
	mu.Lock()
	defer mu.Unlock()

	if inCriticalSection {
		fmt.Println("Already in critical section")
		return
	}

	//Send request to the next client
	msg := &gRPC.CriticalSectionRequest{NodeId: int32(ClientId)}
	_, err := RequestCSHelper(msg)
	if err != nil {
		log.Println("Failed to request critical section: ", err)
		return
	}

	server.SendRequestAccess(context.Background(), msg)
}

func (s *Client) SendRequestAccess(context context.Context, criticalSectionRequest *gRPC.CriticalSectionRequest) (*emptypb.Empty, error) {

	fmt.Printf("Sender: %v", criticalSectionRequest.NodeId)

	return &emptypb.Empty{}, nil
}

func RequestCSHelper(req *gRPC.CriticalSectionRequest) (*gRPC.CriticalSectionRequest, error) {
	mu.Lock()
	defer mu.Unlock()

	if inCriticalSection || int(req.NodeId) == clientId {
		if int(req.NodeId) == clientId {
			enterCriticalSection()
			fmt.Println("Entered critical section for 5 seconds")
		}
		return &gRPC.CriticalSectionRequest{}, nil
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
		time.Sleep(5 * time.Second) // Wait for 5 seconds
		mu.Lock()
		inCriticalSection = false
		fmt.Println("Exited CriticalSection")
		mu.Unlock()
	}()

}

func (s *Client) RequestCriticalSection(stream gRPC.TokenRing_RequestCriticalSectionServer) error {

	fmt.Println("Waiting for message...")
	msg, err := stream.Recv()
	fmt.Println("Receive is not blocking")
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Println("Received message: ", msg.GetNodeId())

	stream.Send(msg)
	return nil
}
