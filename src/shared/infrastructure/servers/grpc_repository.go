package servers

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServerGrpcApp() {
	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	if err != nil {
		log.Fatalf("Error creating repository: %s", err.Error())
	}

	s := grpc.NewServer()

	reflection.Register(s)
	fmt.Println("Server is running on port 5060")

	if err := s.Serve(list); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	} else {
		fmt.Println("Server is running on port 5060")
	}

	// v := fmt.Sprintln("Server is running on port %s", ":5060")
	fmt.Println("Server is running on port 5060")
}
