package grpc

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"

	"infant/comm/etc"
)

func NewClientEtc(key string) *grpc.ClientConn {
	hosts := strings.Split(etc.String("host", key), ",")
	port := etc.Int("service", key)
	addrs := make([]string, 0, len(hosts))
	for _, host := range hosts {
		addrs = append(addrs, fmt.Sprintf("%s:%d", host, port))
	}
	return NewClient(addrs...)
}

func NewClient(addrs ...string) *grpc.ClientConn {
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(time.Second),
		grpc.WithBalancer(grpc.RoundRobin(NewResolver(addrs...))),
	}
	conn, err := grpc.Dial("", dialOpts...)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func NewServer() *grpc.Server {
	return grpc.NewServer()
}

func ServeEtc(s *grpc.Server, key string) {
	port := etc.Int("service", key)
	addr := fmt.Sprintf(":%d", port)
	log.Printf("service %s listening on %s", key, addr)
	Serve(s, addr)
}

func Serve(s *grpc.Server, addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	s.Serve(lis)
}
