package backpack

import (
	"fmt"
	"html/template"
	"path/filepath"
)

type Renderer struct {
	TemplateLocation string
	LayoutLocation   string
	TemplateCache    map[string]*template.Template
	UseCache         bool
}

func NewRenderer(templateLocation, layoutLocation string, useCache bool) (*Renderer, error) {
	cache := map[string]*template.Template{}

	renderer := &Renderer{
		TemplateLocation: templateLocation,
		LayoutLocation:   layoutLocation,
		TemplateCache:    cache,
		UseCache:         useCache,
	}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.go.html", templateLocation))
	if err != nil {
		return renderer, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		template, err := template.New(name).ParseFiles(page)
		if err != nil {
			return renderer, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.go.html", layoutLocation))
		if err != nil {
			return renderer, err
		}

		if len(matches) > 0 {
			template, err = template.ParseGlob(fmt.Sprintf("%s/*.go.html"))
			if err != nil {
				return renderer, err
			}
		}
		cache[name] = template
	}
	return renderer, nil
}

// TODO: Add the actual render function
