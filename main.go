package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"worker/pb"
)

type server struct {
	pb.UnimplementedMathServer
	address net.Addr
}

func (s *server) GetFactors(ctx context.Context, in *pb.IntValue) (*pb.IntList, error) {
	factors := make([]int64, 0)
	n := in.Value
	var i int64
	for i = 2; i<<2 <= n; i++ {
		for ; n%i == 0; n /= i {
			factors = append(factors, i)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	log.Printf("GetFactors(%d) by %v\n", in.Value, s.address)
	return &pb.IntList{Values: factors}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Printf("Cannot connect to \"0.0.0.0:8080\": %v\n", err)
		panic(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterMathServer(s, &server{address: lis.Addr()})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
