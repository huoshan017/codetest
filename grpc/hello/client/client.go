package main

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"codetest/grpc/hello"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("connect service err: ", err)
		return
	}

	defer conn.Close()

	c := hello.NewGreeterClient(conn)

	// SayHello
	r, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "horika"})
	if err != nil {
		fmt.Println("call service SayHello err: ", err)
		return
	}

	fmt.Println("call success: ", r.Message)

	// SayHello2
	s, err := c.SayHello2(context.Background())
	if err != nil {
		fmt.Println("call service SayHello2 err: ", err)
		return
	}

	name := "hello"
	for i := 1; i <= 100; i++ {
		err = s.Send(&hello.HelloRequest{Name: name + strconv.Itoa(i)})
		if err != nil {
			fmt.Println("SayHello2 failed to send, err: ", err)
			break
		}
	}

	reply, err := s.CloseAndRecv()
	if err != nil {
		fmt.Println("SayHello2 failed to CloseAndRecv, err: ", err)
	}

	fmt.Println("SayHello2 get reply: ", reply.Message)

	// SayHello3
	c3, err := c.SayHello3(context.Background(), &hello.HelloRequest{})
	if err != nil {
		fmt.Println("SayHello3 call err: ", err)
		return
	}

	for {
		r3, err := c3.Recv()
		if err == io.EOF {
			fmt.Println("SayHello3 closed")
			break
		}
		if err != nil {
			fmt.Println("SayHello3 recv err: ", err)
		}
		fmt.Println("SayHello3 recv: ", r3.Message)
	}

	// SayHello4
	c4, err := c.SayHello4(context.Background())
	if err != nil {
		fmt.Println("SayHello4 call err: ", err)
		return
	}

	for i := 1; i <= 100; i++ {
		err = c4.Send(&hello.HelloRequest{Name: name + strconv.Itoa(i)})
		if err != nil {
			fmt.Println("SayHello4 call err: ", err)
			break
		}
	}

	for i := 1; i <= 100; i++ {
		var r *hello.HelloReply
		r, err = c4.Recv()
		if err == io.EOF {
			fmt.Println("SayHello4 closed")
			break
		}
		if err != nil {
			fmt.Println("SayHello4 recv err: ", err)
			break
		}
		fmt.Println("SayHello4 recv: ", r.Message)
	}
	c4.CloseSend()
}
