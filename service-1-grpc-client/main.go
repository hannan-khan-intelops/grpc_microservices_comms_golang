// __author__ = "Hannan Khan"
// __credits__ = ["Hannan Khan"]
// __version__ = "0.0.1"
// __email__ = "hkhan@intelops.dev"
// __status__ = "Development"

// This script is the client portion of the grpc microservice. The client will initiate streaming, and the server will
// stream its data (multiple responses) to the client.
// Tutorial src: https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/

package service_1_grpc_client

import (
	pb "example.com/microservice"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	// dail the server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect with server %v", err)
	}

	// create the stream
	client := pb.NewStreamServiceClient(conn)
	in := &pb.Request{id: 1}
	stream, err := client.FetchResponse()
	if err != nil {
		log.Fatalf("Open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true // this means that the stream is finished
				return
			}
			if err != nil {
				log.Fatalf("Cannot recieve %v", err)
			}
			log.Printf("Response received: %s", resp.Result)
		}
	}()
	<-done // we wait until the response is received
	log.Printf("Finished.")
}
