package main

import "fmt"
gRPC ""

func main() {
	// Connect to the clients
	Connect()
	// Start the client
	Start()
}

func Connect() {
	// Connect to the other clients
}

func Start() {
	// broadcasts to other clients
	// listens to other clients
}

func EnterCriticalSection() {
	fmt.Println("Entered CriticalSection")
}

func requestCriticalSection() {
	fmt.Println("Requested CriticalSection")
}

func listen() {
	fmt.Println("Listening")
}
