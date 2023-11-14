package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	gRPC "github.com/seve0039/Distributed-Mutual-Exclusion.git/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	gRPC.UnimplementedTokenRingServer
	name string
	port string
}

var wg sync.WaitGroup
var max = "7"
var clientsName = flag.String("name", "default", "client's name") //TODO: Find a way to get ID
var clientPort = flag.String("server", "5400", "Client's port")
var prevPort = flag.String("prev", "5400", "Previous port")

var server gRPC.TokenRingClient
var serverConn *grpc.ClientConn

/*
func NewClient(id, nextPort string) *Client {
	return &Client{
		id:   id,
		port: "500" + id,
		nextPort: nextPort,
	}

}*/

func main() {
	flag.Parse()

	wg.Add(1)
	wg.Add(1)

	go func() {
		*clientPort = "540" + readFromPortFile()
		*prevPort = "540" + readFromPortFile()

	}()
	wg.Wait()

	writeToPortFile(*clientPort)

	launchConnection()

	sendConnectRequest()
	defer serverConn.Close()

	joinServer()

}

func sendConnectRequest() {

	fmt.Println("Connecting to server...")
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	intport, err := strconv.Atoi(*prevPort)
	if err != nil {
		log.Fatalf("Fail to Dial : %v", err)
	}

	intport--
	*prevPort = strconv.FormatInt(int64(intport), 10)
	fmt.Println(*prevPort)

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", *prevPort), opts...)
	if err != nil {
		log.Fatalf("Fail to Dial : %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to server!")

	// Move these lines outside of the error handling block
	server = gRPC.NewTokenRingClient(conn)
	serverConn = conn

	/*if *clientPort == "540"+max {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", "5400"), opts...)
	if err != nil {
		log.Fatalf("Fail to Dial : %v", err)

		client = gRPC.NewTokenRingClient(conn)
		ClientConn = conn

	}

	*/
}

func launchConnection() {
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *clientPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *clientPort, err)
	}

	grpcServer := grpc.NewServer()
	server := &Client{
		name: *clientsName,
		port: *clientPort,
	}

	gRPC.RegisterTokenRingServer(grpcServer, server)
	log.Printf("NEW SESSION: Server %s: Listening at %v\n", *clientsName, list.Addr())

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}

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

func requestCriticalSection() {
	fmt.Println("Requested CriticalSection")
}

func readFromPortFile() string {
	defer wg.Done()
	file, err := os.Open("Ports.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	// Get file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file size:", err)
		return ""
	}
	// Read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	b := bs[len(bs)-1]
	str := string(b)
	return str
}

func writeToPortFile(port string) {

	filePath := "Ports.txt"

	// Create or open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write content to the file
	newPort, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Error converting port to int:", err)
		return
	}
	newPort++
	strport := strconv.FormatInt(int64(newPort), 10)

	content := []byte(strport)
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
