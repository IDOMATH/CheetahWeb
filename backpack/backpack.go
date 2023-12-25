package backpack

import (
	"fmt"
	"net/http"
)

type Backpack struct {
	Port      int
	Routes    map[string]Route
	Handle405 http.HandlerFunc
}

type Route struct {
	Method     string
	Url        string
	Middleware []func()
	Handler    http.HandlerFunc
}

func (b *Backpack) Serve() {
	fmt.Printf("Running on port: %d", b.Port)
	for _, route := range b.Routes {
		for _, middleware := range route.Middleware {
			middleware()
		}
		http.HandleFunc(route.Url, route.Handler)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", b.Port), nil)
}

func (b *Backpack) Get(url string, handler http.HandlerFunc) {
	acceptedMethod := "GET"
	hf := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == acceptedMethod {
			handler(w, r)
		} else {
			if b.Handle405 != nil {
				b.Handle405(w, r)
			} else {
				DefaultHandle405(w, r)
			}
		}
	}
	route := Route{
		Method:  acceptedMethod,
		Url:     url,
		Handler: hf,
	}
	b.Routes[acceptedMethod+url] = route
}

func (b *Backpack) Post(url string, handler http.HandlerFunc) {
	acceptedMethod := "POST"
	hf := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == acceptedMethod {
			handler(w, r)
		} else {
			if b.Handle405 != nil {
				b.Handle405(w, r)
			} else {
				DefaultHandle405(w, r)
			}
		}
	}
	route := Route{
		Method:  acceptedMethod,
		Url:     url,
		Handler: hf,
	}
	b.Routes[acceptedMethod+url] = route
}

func (b *Backpack) Put(url string, handler http.HandlerFunc) {
	acceptedMethod := "PUT"
	hf := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == acceptedMethod {
			handler(w, r)
		} else {
			if b.Handle405 != nil {
				b.Handle405(w, r)
			} else {
				DefaultHandle405(w, r)
			}
		}
	}
	route := Route{
		Method:  acceptedMethod,
		Url:     url,
		Handler: hf,
	}
	b.Routes[acceptedMethod+url] = route
}

func (b *Backpack) Delete(url string, handler http.HandlerFunc) {
	acceptedMethod := "DELETE"
	hf := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == acceptedMethod {
			handler(w, r)
		} else {
			if b.Handle405 != nil {
				b.Handle405(w, r)
			} else {
				DefaultHandle405(w, r)
			}
		}
	}
	route := Route{
		Method:  acceptedMethod,
		Url:     url,
		Handler: hf,
	}
	b.Routes[acceptedMethod+url] = route
}

func (b *Backpack) RegisterMiddleware(route string, middleware func()) {
	registerTo := b.Routes[route]
	registerTo.Middleware = append(registerTo.Middleware, middleware)
}

func DefaultHandle405(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method not supported"))
}
