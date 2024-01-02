package main

import (
	"fmt"
	backpack "github.com/idomath/CheetahWeb/backpack"
	"net/http"
)

var bp *backpack.Backpack

func main() {
	//renderer, err := backpack.NewRenderer("./templates", "./layouts", true)
	//if err != nil {
	//	log.Fatal("failed to create renderer")
	//}
	bp = backpack.NewBackpack(8080)
	//bp.Renderer = renderer

	//TODO: what to do about passing renderer to handlers to actually render the templates

	bp.Get("/", handleHome)
	homeRoute := bp.Routes["GET/"]
	homeRoute.RegisterMiddleware(logMiddleware)
	bp.Routes["GET/"] = homeRoute
	//TODO: figure out how to make middleware work
	//backpack.RegisterMiddleware("GET/", logMiddleware)

	bp.Serve()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	//err := bp.Renderer.Template(w, r, "home.go.html", &backpack.TemplateData{
	//	PageTitle: "Home",
	//})
	//w.Write([]byte(err.Error()))
	w.Write([]byte("Welcome HOme"))
}

func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logging")
		handler(w, r)
	}
}
