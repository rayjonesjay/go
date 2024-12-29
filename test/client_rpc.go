package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type QuotRem struct {
	Q, R int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage", os.Args[0], "server")
		os.Exit(1)
	}

	// where is the server
	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":3000")
	if err != nil {
		log.Fatal("dialing", err)
	}
	//synchronous call
	args := Args{12, 5}
	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith err", err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quotRem QuotRem
	err = client.Call("Arith.Divide", args, &quotRem)
	if err != nil {
		log.Fatal("divide, err:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quotRem.Q, quotRem.R)

	fmt.Println("testing evenity")
	var b *bool
	// var wg sync.WaitGroup
	// wg.Add(1)
	// defer wg.Done()
	client.Call("Arith.IsEven", 4, &b)
	fmt.Println(*b)
	// wg.Wait()
}
