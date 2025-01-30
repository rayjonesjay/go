package main

import (
	"encoding/json"
	"fmt"
	"log"

	"todo/admin"

	supabase "github.com/supabase-community/supabase-go"
)

func main() {
	Client, err := supabase.NewClient(admin.API_URL, admin.API_KEY, &supabase.ClientOptions{})
	if err != nil {
		log.Println("error creating client", err)
	}

	// data, count, err := Client.From("todos").Select("*", "exact", false).Execute()
	// if err != nil {
	// 	log.Println("error fetching", err)
	// }
	// fmt.Println(string(data))
	// fmt.Println(string(count))

	task := "Be the best"

	marsh, _ := json.Marshal(task)

	data, count, err := Client.From("todos").Update(marsh, "", "").Single().Execute()
	if err != nil {
		log.Println("error fetching", err)
	}
	fmt.Println(string(data))
	fmt.Println(string(count))
}
