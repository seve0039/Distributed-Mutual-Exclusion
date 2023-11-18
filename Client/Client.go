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

	go startServer()

	sendConnectRequest()

	defer serverConn.Close()

	joinServer()

	/*stream, err := server.RCS(context.Background())
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}*/

	for {
	}
}

func sendConnectRequest() {

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

func joinServer() {
	_, err := server.Join(context.Background(), &gRPC.JoinRequest{NodeId: *clientsName})
	if err != nil {
		log.Fatalf("Failed to join server: %v", err)
	}
}

func (c *Client) Join(ctx context.Context, joinReq *gRPC.JoinRequest) (*gRPC.JoinAck, error) {
	ack := &gRPC.JoinAck{Message: fmt.Sprintf("Welcome to Chitty-Chat, %s!", joinReq.NodeId)}
	return ack, nil
}

func EnterCriticalSection() {
	fmt.Println("Entered CriticalSection")
}

/*func listenForBroadcast(stream gRPC.TokenRingClient) {
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

}*/

func requestCriticalSection(ClientId int, stream gRPC.TokenRing_RCSClient) {

	fmt.Println("Requested CriticalSection")
}

func checkAndChangePort() {
	if *clientPort > 5400+max {
		*clientPort = 5400
	} else {
		*clientPort = 5400
	}
}
