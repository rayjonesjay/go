package main

import (
	"fmt"
	"plugin"
	"time"
)

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	s, err := p.Lookup("Ray")
	if err != nil {
		panic(err)
	}

	sys := s.(func())

	fmt.Println("plugin will run in 10 sec")
	time.Sleep(10 * time.Second)

	sys()
}
