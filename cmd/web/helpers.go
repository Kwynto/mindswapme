package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) getMd5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// We extract the corresponding set of templates from the cache, depending on the page name
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("template %s not exist", name))
		return
	}

	// Render template files by passing dynamic data from the `td` variable.
	err := ts.Execute(w, td)
	if err != nil {
		app.serverError(w, err)
	}
}
