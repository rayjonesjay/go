package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

/*
	first you define a go struct that looks like the data thats coming in from the api
*/

type User struct {
	Name  string `json:"name"` // these tags represent the keys that the json data type will have and the type(int,string) represent the type of value that key will be associated with.
	Age   int    `json:"age,omitempty"`
	Email string `json:"email"`
}

func main() {
	go startApi()
	url := "http://localhost:9000/user" // represent a real api

	fmt.Println("getting data from", url)
	// wait for 2 seconds
	time.Sleep(time.Duration(2) * time.Second)
	getFromApi(url)

	fmt.Println("*****************")
	// create an instance of Usr
	usr := Usr{
		ID: 1, Name: "BOBO", Contact: Contact{
			Email: "bobo@email.com",
			Phone: "1234",
		},
	}
	convertNestedToJson(usr)
}

func getFromApi(url string) {
	// make http Get request to the endpoint of our api
	req, err := http.NewRequest("GET", url, nil)

	checkErr(err)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error fetching data: ", err)
		return // dont proceed
	}
	defer resp.Body.Close()

	if !resp.Close {
		fmt.Println("not closed yet value is", resp.Close)
	}

	// read the body that the api gave us after sending the get request
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response", err)
		return
	}
	fmt.Println(">>>>", string(body))

	var user User
	fmt.Println("unmarshalling...")
	// when you get data from an api you need to unmarshall,unpack,decode it
	err = json.Unmarshal(body, &user)
	checkErr(err)
	fmt.Println("done...")

	// print our struct to see how the data was unmarshalled into our struct
	fmt.Println("user name: ", user.Name)
	fmt.Println("user age: ", user.Age)
	fmt.Println("user email: ", user.Email)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error occurred", err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		fmt.Println(r.URL.Path)
		w.Write([]byte("wrong path"))
		return
	}

	if r.Method != http.MethodGet {
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	// now we need to simulate sending json to user
	users := []User{
		{Name: "alice", Age: 20, Email: "alice@email.com"},
		{Name: "john", Age: 22, Email: "john@email.com"},
	}

	// to convert a go struct to a json you marshall it or encode or pack it
	jsonData, er := json.Marshal(users)
	checkErr(er)
	fmt.Println("marshalled json", string(jsonData))

	fmt.Println("received GET request from....", r.RemoteAddr)

	u := User{
		Name:  "ray",
		Age: 80,
		Email: "ray@gmail.com",
	}

	userJson, err := json.Marshal(u)
	checkErr(err)

	fmt.Println("sending such data to user", string(userJson))

	w.Header().Add("Connection", "close")
	w.WriteHeader(202)
	w.Header().Add("Content-Type", "application/json")
	w.Write(userJson)
}

func startApi() {
	fmt.Println("starting server... at http://localhost:9000")
	mux := http.NewServeMux()
	mux.Handle("GET /user", http.HandlerFunc(userHandler))
	err := http.ListenAndServe("localhost:9000", mux)
	checkErr(err)

}

// the Contact has two fields, email and phone
type Contact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Define main struct which includes the nested struct
type Usr struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Contact Contact `json:"contact"` // nested struct
}

func convertNestedToJson(usr Usr) {
	jsonData, err := json.MarshalIndent(usr,"", " ")
	checkErr(err)
	fmt.Println(string(jsonData))
}
