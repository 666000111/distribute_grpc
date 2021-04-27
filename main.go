package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	"grpc_server/proto_gen/grpc_server/book"
	"log"
	"net"
	"grpc_server/server"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	book.RegisterBookServer(grpcServer,newServer())
	grpcServer.Serve(lis)
}

func newServer() *server.BookServer {
	s := &server.BookServer{}
	return s
}