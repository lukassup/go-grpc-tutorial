package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/lukassup/go-grpc-tutorial/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DEFAULT_SOCK_PROTO = "tcp"
	DEFAULT_SOCK_ADDR  = "127.0.0.1:9991"
	DEFAULT_MESSAGE    = "ping"
	LOG_PREFIX_CLIENT  = "[client] "
	LOG_PREFIX_SERVER  = "[server] "
)

type echoServer struct {
	pb.EchoServiceServer
}

func (s *echoServer) Echo(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("received request: data=%q\n", req.Data)

	resp := &pb.Response{
		Status: pb.Response_OK,
		Data:   req.Data,
	}
	log.Printf("sending response: data=%q, status=%s\n", resp.Data, resp.Status)
	return resp, nil
}

func runServer(sockAddr string) {
	listner, err := net.Listen(DEFAULT_SOCK_PROTO, sockAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("listening on %s://%s\n", DEFAULT_SOCK_PROTO, sockAddr)

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, &echoServer{})
	if err := server.Serve(listner); err != nil {
		log.Fatalf("failed to serve, error: %v\n", err)
	}
}

func clientConnect(sockAddr string, data string) {
	log.Printf("connecting to %s://%s\n", DEFAULT_SOCK_PROTO, sockAddr)
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	connection, err := grpc.Dial(sockAddr, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	client := pb.NewEchoServiceClient(connection)
	request := &pb.Request{Data: data}
	log.Printf("sending request: data=%q\n", request.Data)

	response, err := client.Echo(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("received response: data=%q, status=%s\n", response.Data, response.Status)
}

func main() {
	// set logger flags to move prefix before message
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)

	serverFlag := flag.Bool("server", false, "run server")
	sockAddrFlag := flag.String("addr", DEFAULT_SOCK_ADDR, "socket `address` for server and client")
	flag.Parse()

	if *serverFlag {
		log.SetPrefix(LOG_PREFIX_SERVER)
		runServer(*sockAddrFlag)
	} else {
		log.SetPrefix(LOG_PREFIX_CLIENT)
		// use default message unless one was provided as first arg
		data := DEFAULT_MESSAGE
		if flag.NArg() == 1 {
			data = flag.Arg(0)
		}
		clientConnect(*sockAddrFlag, data)
	}
}
