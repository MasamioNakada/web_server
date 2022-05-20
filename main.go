package main

import "fmt"

func main() {
	fmt.Println("Server is running")
	server := NewServer(":3001")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()

}
