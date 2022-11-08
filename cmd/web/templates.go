package main

import (
	"html/template"
	"path/filepath"
)

// Structure for the data template
type tmplHomeData struct {
	isLinks bool
	Link1, Link2, Link3, Link4, Link5, Link6, Link7, Link8, Link9, Link10,
	Link11, Link12, Link13, Link14, Link15, Link16, Link17, Link18, Link19 msmLink
}

type tmplManageData struct {
	User  string
	Hash  string
	Links []msmLink
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// We use the filepath.Glob function to get a slice of all file paths with the extension '.page.tmpl'.
	pages, err := filepath.Glob(filepath.Join(dir, "*.html"))
	if err != nil {
		return nil, err
	}

	// We iterate through the template file from each page.
	for _, page := range pages {
		// Extracting the final file name
		name := filepath.Base(page)

		// Processing the iterated template file.
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Adding the resulting set of templates to the cache using the page name
		cache[name] = ts
	}

	// We return the received map.
	return cache, nil
}
