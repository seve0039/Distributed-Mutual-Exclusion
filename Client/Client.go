package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	gRPC "github.com/seve0039/Distributed-Mutual-Exclusion.git/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	gRPC.UnimplementedTokenRingServer
	//participants     map[string]
	participantMutex sync.RWMutex
	name             string
	port             string
}

var ID string = "0"
var clientsName = flag.String("name", ID, "Sender's name")
var clientPort = flag.String("server", "5400", "Client's port")

var client gRPC.TokenRingClient
var ClientConn *grpc.ClientConn

func main() {

	flag.Parse()
	// Connect to the clients
	launchConnection()
	defer ClientConn.Close()

	joinServer()

	sendConnectRequest()

	// Listen for connections from other clients

	// Listen for messages from other clients
	//go listenForBroadcast(stream)

}


func sendConnectRequest() {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(fmt.Sprintf(":%s", *clientPort), opts...)
	if err != nil {
		log.Fatalf("Fail to Dial : %v", err)
	}

	client = gRPC.NewTokenRingClient(conn)
	ClientConn = conn
}


func launchConnection() {
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *clientPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *clientPort, err)
	}

	clientConnection := grpc.NewServer()
	server := &Client{
		name: *clientsName,
		port: *clientPort,
	}

	gRPC.RegisterTokenRingServer(clientConnection, server)
	log.Printf("NEW SESSION: Server %s: Listening at %v\n", *clientsName, list.Addr())

	if err := clientConnection.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}


func joinServer() {
	_, err := client.Join(context.Background(), &gRPC.JoinRequest{Name: *clientsName})
	if err != nil {
		log.Fatalf("Failed to join server: %v", err)
	}
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


func readFromPortFile() {
	file, err := os.Open("Ports.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Get file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file size:", err)
		return
	}

	// Read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func writeToPortFile(port string){
	data := []byte(port + "\n")

	file, err := os.OpenFile("Ports.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Append the new data to the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("Data appended to file successfully.")
}
