package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/luci/go-render/render"
	"google.golang.org/grpc"
	"grpc_server/proto_gen/grpc_server/book"
	//"google.golang.org/grpc/credentials"
	"log"
)
var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)



func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts,grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := book.NewBookClient(conn)
	c := context.Background()
	client.AddBook(c,&book.AddBookReq{
		Name:    "hhhhh",
		Content: "rpcSuuccess",
	})
	client.DeleteBook(c,&book.DeleteBookReq{Id: 3})
	client.ModifyBook(c,&book.ModifyBookReq{Id: 1,Content: "sdasdad",Name: "dsadsdad"})
	resp,_ := client.QueryBook(c,&book.QueryBookReq{})
	for _,v := range resp.Books{
		fmt.Println(render.Render(v))
	}

}
