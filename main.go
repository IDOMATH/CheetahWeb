package main

import (
	"fmt"
	backpack "github.com/idomath/CheetahWeb/backpack"
	"log"
	"net/http"
)

func main() {
	renderer, err := backpack.NewRenderer("./", "./", true)
	if err != nil {
		log.Fatal("failed to create renderer")
	}
	bp := backpack.NewBackpack(8080)
	bp.Renderer = renderer

	//TODO: what to do about passing renderer to handlers to actually render the templates

	bp.Get("/", logMiddleware(handleHome))
	//TODO: figure out how to make middleware work
	//backpack.RegisterMiddleware("GET/", logMiddleware)

	bp.Serve()
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
