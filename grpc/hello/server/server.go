package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"

	"codetest/grpc/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "hello " + in.Name}, nil
}

func (s *server) SayHello2(g hello.Greeter_SayHello2Server) error {
	var l string
	for {
		r, err := g.Recv()
		if err == io.EOF {
			fmt.Println("SayHello2 closed")
			break
		}
		if err != nil {
			fmt.Println("SayHello2 recv err: ", err)
			break
		}
		l += r.Name
	}
	g.SendAndClose(&hello.HelloReply{Message: l})
	return nil
}

func (s *server) SayHello3(in *hello.HelloRequest, g hello.Greeter_SayHello3Server) error {
	for i := 0; i < 100; i++ {
		err := g.Send(&hello.HelloReply{Message: "name" + in.Name + strconv.Itoa(i)})
		if err != nil {
			fmt.Println("SayHello3 send err: ", err)
			break
		}
	}
	return nil
}

func (s *server) SayHello4(g hello.Greeter_SayHello4Server) error {
	for {
		r, err := g.Recv()
		if err == io.EOF {
			fmt.Println("SayHello4 closed")
			break
		}
		if err != nil {
			fmt.Println("SayHello4 recv err: ", err)
			break
		}
		err = g.Send(&hello.HelloReply{Message: "name" + r.Name})
		if err != nil {
			fmt.Println("SayHello4 send err: ", err)
			break
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("listen port err: %v", err)
		return
	}

	s := grpc.NewServer()

	hello.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)

	if err != nil {
		fmt.Printf("open service err: %v", err)
		return
	}
}
