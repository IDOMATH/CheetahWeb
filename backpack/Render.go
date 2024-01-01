package backpack

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Renderer struct {
	TemplateLocation string
	LayoutLocation   string
	TemplateCache    map[string]*template.Template
	UseCache         bool
}

type TemplateData struct {
	PageTitle string
	BoolMap   map[string]bool
	StringMap map[string]string
	IntMap    map[string]string
	ObjectMap map[string]interface{}
}

func NewRenderer(templateLocation, layoutLocation string, useCache bool) (*Renderer, error) {
	var cache map[string]*template.Template

	renderer := &Renderer{
		TemplateLocation: templateLocation,
		LayoutLocation:   layoutLocation,
		TemplateCache:    cache,
		UseCache:         useCache,
	}
	cache, err := renderer.createTemplateCache()
	if err != nil {
		return renderer, err
	}

	return renderer, nil
}

func (ren *Renderer) Template(w http.ResponseWriter, r *http.Request, tmpl string, data *TemplateData) error {
	var tc map[string]*template.Template
	var err error

	fmt.Println("Rendering template...")

	if ren.UseCache && ren.TemplateCache != nil {
		tc = ren.TemplateCache
	} else {
		tc, err = ren.createTemplateCache()
	}
	if err != nil {
		return err
	}
	fmt.Println("TC created of length", len(tc))

	for temp := range tc {
		fmt.Println(temp)
	}
	fmt.Println("Attempting to get template: ", tmpl)
	t, ok := tc[tmpl]
	if !ok {
		return errors.New("can't get template from cache")
	}
	fmt.Println("Got template from cache")
	buf := new(bytes.Buffer)

	err = t.Execute(buf, data)
	if err != nil {
		//TODO: probably shouldn't explode when a template fails to execute.
		log.Fatal(err)
	}
	fmt.Println("Executed  template")
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return err
	}
	fmt.Println("Wrote to w")

	return nil
}

func (ren *Renderer) createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.go.html", ren.TemplateLocation))
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		template, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.go.html", ren.LayoutLocation))
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			template, err = template.ParseGlob(fmt.Sprintf("%s/*.go.html"))
			if err != nil {
				return cache, err
			}
		}
		cache[name] = template
	}
	return cache, nil
}
