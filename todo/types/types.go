package types

type Todo struct {
	Text string // task to be done for example: "dont stop working hard"
	Done bool   // indicate whether the task is done or not
}

// a collection of todos
var Todos []Todo
var LoggedInUser string
