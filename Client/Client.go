package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	gRPC "github.com/seve0039/Distributed-Mutual-Exclusion.git/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientId = 0
var max = 3
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

	serverStream, err := server.RCS(context.Background())
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}

	clientStream, err := serverConn.RCS(context.Background())
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}
	//client
	requestCriticalSection(int64(*clientPort), serverStream)
	//server
	go listenForMessage(clientStream)

	for {
	}
}

//Client
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

func requestCriticalSection(ClientId int64, stream gRPC.TokenRing_RCSClient) {

	msg := &gRPC.CriticalSectionRequest{NodeId: ClientId}
	stream.Send(msg)

}


//Server
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



func (s *Client) RCS(stream gRPC.TokenRing_RCSServer) {

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Println("Failed to receive message: ", err)
			return
		}

		fmt.Println(msg.GetNodeId())
	}

}
func checkMessageId(id int64, stream gRPC.TokenRing_RCSClient) {

	if clientId != int(id) {
		requestCriticalSection(id, stream)
	}
}

func EnterCriticalSection() {
	fmt.Println("Entered CriticalSection")
}
