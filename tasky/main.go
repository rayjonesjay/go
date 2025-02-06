package main

import (
	"context"
	"fmt"
	"log"
	"tasky/db"
	"tasky/models"
)

func main() {
	db.ConnectDB()

	defer db.Conn.Close(context.Background())

	// // test user functions
	// err := models.CreateUser("test@example.org", "nemesis123")
	// if err != nil {
	// 	log.Fatalf("error creating user: %v", err)
	// }

	user, err := models.GetUserByEmail("test@example.org")
	if err != nil {
		panic(err)
	}

	fmt.Printf("got %+v\n", user)

	err = models.CreateTask(user.Id, "my first task", "this is a test", "work")

	if err != nil {
		log.Println(err)
	}

	tasks, err := models.GetTasksByUser(user.Id)
	if err != nil {
		log.Fatalf("error retrieving tasks %v", err)
	}

	fmt.Printf("retrieved tasks: %+v\n", tasks)
}
