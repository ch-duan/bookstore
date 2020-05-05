package main

import "bookstore/server"

func main() {
	r := server.NewRouter()
	r.Run(":8080")
}
