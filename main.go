package main

import (
	"fmt"
	backpack "github.com/idomath/CheetahWeb/backpack"
	"net/http"
)

func main() {
	backpack := backpack.Backpack{
		Port:   8080,
		Routes: make(map[string]backpack.Route),
	}

	backpack.Get("/", handleHome)
	//TODO: figure out how to make middleware work
	backpack.RegisterMiddleware("GET/", logMiddleware)

	backpack.Serve()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome home"))
}

func logMiddleware() {
	fmt.Println("logging")
}
