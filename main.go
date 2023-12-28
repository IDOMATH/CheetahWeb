package main

import (
	"fmt"
	backpack "github.com/idomath/CheetahWeb/backpack"
	"net/http"
)

func main() {
	backpack := backpack.NewBackpack(8080)

	backpack.Get("/", logMiddleware(handleHome))
	//TODO: figure out how to make middleware work
	//backpack.RegisterMiddleware("GET/", logMiddleware)

	backpack.Serve()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome home"))
}

func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logging")
		handler(w, r)
	}
}
