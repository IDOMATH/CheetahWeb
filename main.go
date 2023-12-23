package main

import (
	backpack "github.com/idomath/CheetahWeb/backpack"
	"net/http"
)

func main() {
	backpack := backpack.Backpack{
		Port:      8080,
		Routes:    []backpack.Route{},
		Handle405: handleWrongMethod,
	}

	backpack.Get("/", handleHome)

	backpack.Serve()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome home"))
}

func handleWrongMethod(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method not supported"))
}
