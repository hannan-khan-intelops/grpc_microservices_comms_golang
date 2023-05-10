// __author__ = "Hannan Khan"
// __credits__ = ["Hannan Khan"]
// __version__ = "0.0.1"
// __email__ = "hkhan@intelops.dev"
// __status__ = "Development"

// This script is the client portion of the grpc microservice. The client will initiate streaming, and the server will
// stream its data (multiple responses) to the client.
// Tutorial src: https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/

// package has to be main in order for this to run as a standalone script.
package main

// import required modules, including the microservice module we have created using the proto generated code.
import (
	pb "example.com/microservice"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

// add a custom server struct that we can use for later. This will also implement an empty unimplmented method, which is
// required to work.
type server struct {
	pb.UnimplementedStreamServiceServer
}

// add a constant for the server address.
const (
	serverAddress = "localhost:50005"
)

// FetchResponse we implement a FetchResponse function for our server. This will take the request from the client, and do what needs
// to be accomplished within the for loop (as this is a streaming service).
func (s server) FetchResponse(in *pb.Request, srv pb.StreamService_FetchResponseServer) error {

	log.Printf("Fetch response for response.id: %d", in.Id)

	// we will be using a waitGroup in order to allow the process to be concurrent, and will wait for multiple
	// goroutines to finish.
	var wg sync.WaitGroup
	// We stream five objects.
	for i := 0; i < 5; i++ {
		// we add this particular stream to the waitGroup.
		wg.Add(1)
		// we begin the goroutine, with one variable, count (same as the index) of seconds to sleep.
		go func(count int64) {
			// we defer the finishing of this waitGroup until we are done with this subroutine.
			defer wg.Done()

			// this sleep function will simulate server process time.
			// HERE IS WHERE YOU WOULD IMPLEMENT ANOTHER FUNCTION TO CARRY OUT BASED ON USE CASE.
			time.Sleep(time.Duration(count) * time.Second)
			// we generate a response with a result.
			resp := pb.Response{Result: fmt.Sprintf("Request #%d for Id: %d", count, in.Id)}
			// we send the response and check for errors at the same time.
			if err := srv.Send(&resp); err != nil {
				log.Printf("Send error %v", err)
			}
			log.Printf("Finishing request number: %d", count)
		}(int64(i)) // we pass in the current index i, as the count for this goroutine.
	}
	// we wait for all the waitgroups to finish.
	wg.Wait()
	// we finish the function.
	return nil
}

// the main entry point for this script. Also required to run as a standalone script.
func main() {
	// create a listener to listen for the request.
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// create the actual gRPC server
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, server{})

	log.Printf("Starting server at: %v\n", serverAddress)
	// start listening via the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
