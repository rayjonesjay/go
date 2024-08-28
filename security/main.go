// package main

// import( 
// 	"fmt"
// 	// "log"
// 	"net/http"
// 	"os"
// )

// func getProgName(name string) string {
// 	var prog string 
// 	for i := len(name)-1; i >= 0; i--{
// 		if name[i] == '/'{
// 			break
// 		}
// 		prog += string(name[i])
// 	}
// 	slice := []rune(prog)
// 	for i , j := 0, len(slice)-1; i < j ; i , j = i+1, j-1{
// 		slice[i] , slice[j] = slice[j] , slice[i]
// 	}
// 	return string(slice)
// }

// func printUsage(programName string) {
// 	programName = getProgName(programName)
// 	fmt.Println(programName + ` <listenUrl> <directory> `, "\n" ,`example:` , "\n" ,programName+` localhost:8080 .`, "\n",programName+` 0.0.0.0:9999 /home/ubuntu`)
// }

// func checkArgs() (string,string) {
// 	args := os.Args
// 	if len(args) != 3 {
// 		printUsage(args[0])	
// 		os.Exit(1)
// 	}
// 	return args[1], args[2]
// }

// func main() {
// 	fmt.Println("hello world the server is back up")
// 	listenerUrl, directoryPath := checkArgs()
// 	err := http.ListenAndServe(listenerUrl, http.FileServer(http.Dir(directoryPath)))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error serving file\n")
// 		os.Exit(0)
// 	}
// }



package main
import (
 "fmt"
 "net/http"
 "log"
)
func indexHandler(writer http.ResponseWriter, request *http.Request) {
 // Write the contents of the response body to the writer interface
 // Request object contains information about and from the client
 fmt.Fprintf(writer, "You requested: " + request.URL.Path)
}
func main() {
 http.HandleFunc("/", indexHandler)
 err := http.ListenAndServe("localhost:8080", nil)
 if err != nil {
 log.Fatal("Error creating server. ", err)
 }
 fmt.Println("hellow")
}