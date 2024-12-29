package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Values struct {
	A, B int
}

// any division operation has a quotient and a reminder
type QuotRem struct {
	Q, R int
}

type Arith int

func (t *Arith) IsEven(num *Arith, ans *bool) error {
	if *num%2 == 0 {
		*ans = true
	} else {
		*ans = false//means its odd
	}
	return nil
}

func (t *Arith) Multiply(args *Values, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Values, q *QuotRem) error {
	if args.B == 0 {
		return errors.New("error division by zero")
	}
	q.Q = args.A / args.B
	q.R = args.B % args.B
	return nil
}

func Register() {
	fmt.Println("registering...")
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	Register()
}
