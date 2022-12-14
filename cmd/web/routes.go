package main

import (
	"net/http"
	"path/filepath"
)

// Blocking direct access to the file system
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, _ := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

func (app *application) routes() *http.ServeMux {
	// Routes
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)

	mux.HandleFunc("/manage", app.manage)
	mux.HandleFunc("/auth", app.manageAuthPage)
	mux.HandleFunc("/logout", app.manageOutPage)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
