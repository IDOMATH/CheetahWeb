package backpack

import (
	"fmt"
	"net/http"
)

type Backpack struct {
	Port   int
	Routes []Route
}

type Route struct {
	Method  string
	Url     string
	Handler *http.HandlerFunc
}

func (b *Backpack) Serve() {
	fmt.Printf("Running on port: %d", b.Port)
	for _, route := range b.Routes {
		http.HandleFunc(route.Url, *route.Handler)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", b.Port), nil)
}

func (b *Backpack) Get(url string, handler http.HandlerFunc) {
	route := Route{
		Method:  "GET",
		Url:     url,
		Handler: &handler,
	}
	b.Routes = append(b.Routes, route)
}

func (b *Backpack) Post(url string, handler http.HandlerFunc) {
	route := Route{
		Method:  "POST",
		Url:     url,
		Handler: &handler,
	}
	b.Routes = append(b.Routes, route)
}

func (b *Backpack) Put(url string, handler http.HandlerFunc) {
	route := Route{
		Method:  "PUT",
		Url:     url,
		Handler: &handler,
	}
	b.Routes = append(b.Routes, route)
}

func (b *Backpack) Delete(url string, handler http.HandlerFunc) {
	route := Route{
		Method:  "DELETE",
		Url:     url,
		Handler: &handler,
	}
	b.Routes = append(b.Routes, route)
}
