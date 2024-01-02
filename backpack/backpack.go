package backpack

import (
	"fmt"
	"net/http"
)

type Backpack struct {
	Port        int
	Routes      map[string]Route
	Middlewares Middleware
	Renderer    *Renderer
	Handle405   http.HandlerFunc
}

func NewBackpack(port int) *Backpack {
	return &Backpack{
		Port:   port,
		Routes: make(map[string]Route),
	}
}

type Route struct {
	Method     string
	Url        string
	Middleware func()
	Handler    http.HandlerFunc
}

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func (b *Backpack) Serve() {
	fmt.Printf("Running on port: %d", b.Port)
	for _, route := range b.Routes {
		//route.Middleware()
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

func (route *Route) RegisterMiddleware(middleware Middleware) {
	route.Handler = func(w http.ResponseWriter, r *http.Request) {
		middleware(route.Handler)
	}
}

func DefaultHandle405(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method not supported"))
}
