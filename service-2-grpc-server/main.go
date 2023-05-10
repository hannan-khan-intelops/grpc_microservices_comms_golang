// __author__ = "Hannan Khan"
// __credits__ = ["Hannan Khan"]
// __version__ = "0.0.1"
// __email__ = "hkhan@intelops.dev"
// __status__ = "Development"

// This script is the client portion of the grpc microservice. The client will initiate streaming, and the server will
// stream its data (multiple responses) to the client.
// Tutorial src: https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/

package main

import (
	pb "example.com/microservice"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

type server struct {
	pb.UnimplementedStreamServiceServer
}

func (s server) FetchResponse(in *pb.Request, srv pb.StreamService_FetchResponseServer) error {

	log.Printf("Fetch response for response.id: %d", in.Id)

	// we will be using a waitGroup in order to allow the process to be concurrent.
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()

			// this sleep function will simulate server process time.
			time.Sleep(time.Duration(count) * time.Second)
			resp := pb.Response{Result: fmt.Sprintf("Request #%d for Id: %d", count, in.Id)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("Send error %v", err)
			}
			log.Printf("Finishing request number: %d", count)
		}(int64(i))
	}
	wg.Wait()
	return nil
}

//func (s server) mustEmbedUnimplementedStreamServiceServer() {}

func main() {
	// create a listener
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// create the actual gRPC server
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, server{})

	log.Println("Starting server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
