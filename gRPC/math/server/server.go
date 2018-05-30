package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	pb "math/math"
)

type server struct{}

var port = ":8080"

func (s *server) Div(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	n, d := in.Dividend, in.Divisor
	if d == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "division by zero")
	}
	return &pb.Response{
		Quotient:  int64(n / d),
		Remainder: int64(n % d),
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterMathServer(s, &server{})
	s.Serve(lis)
}
