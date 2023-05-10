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
	"context"
	pb "example.com/microservice"
	"google.golang.org/grpc"
	"io"
	"log"
)

// add a constant for the server address.
const (
	serverAddress = "localhost:50005"
)

// the main entry point for this script. Also required to run as a standalone script.
func main() {
	// dail the server, and create a connection object.
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect with server %v", err)
	}

	// create the client object
	client := pb.NewStreamServiceClient(conn)
	// create the request object. This request will be giving us the data from the stream (i.e. telling the stream to
	// start)
	in := &pb.Request{Id: 1}
	// create the stream by calling the FetchResponse function. This will start the streaming, as soon as the server
	// responds.
	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("Open stream error %v", err)
	}

	// note that this channel is not like make(chan bool, 1) because this is an unbuffered channel. This means that the
	// reads and writes to/from this channel are blocking. This is extremely helpful when dealing with routines. This is
	// why we will use this type of channel to create a done signal to indicate to the routines that we are done.
	done := make(chan bool)

	// We create a go routine to run functions synchronously.
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
