package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type User struct {
	Timetaken time.Duration
	Name      string
	Path      string
}

var users = []User{}

func main() {
	u1 := User{
		Name:      "John",
		Timetaken: 2,
		Path:      "/",
	}
	u2 := User{
		Name:      "Paul",
		Timetaken: 3,
		Path:      "/STAT",
	}
	u3 := User{
		Name:      "Petre",
		Timetaken: 4,
		Path:      "/",
	}
	u4 := User{
		Name:      "Gumo",
		Timetaken: 5,
		Path:      "//",
	}
	// 4 users will make the request to / if a user takes longer than 3 seconds his request gets cancelled

	users = append(users, u1, u2, u3, u4)
	wg := sync.WaitGroup{}
	for _, u := range users {
		wg.Add(1)
		go Server(u, &wg)
	}
	wg.Wait()
}

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%+v\n", r)
	time.Sleep(3 * time.Second)
	if r.URL.Path != "/" {
		http.Error(rw, "not found", 404)
		return
	}
	rw.WriteHeader(200)
	rw.Write([]byte("how are you"))
}

func Server(u User, wg *sync.WaitGroup) {
	defer wg.Done()
	mux := http.NewServeMux()
	mux.HandleFunc(u.Path, helloHandler)
	server := http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		WriteTimeout: time.Second * u.Timetaken,
	}
	err := server.ListenAndServe()
	fmt.Printf("%s requested %q %v\n", u.Name, u.Path, err)
}
