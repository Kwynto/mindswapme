package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	tD := tmplHomeData{
		Links: make([]msmLink, 20),
	}
	tD.Links = append(tD.Links,
		msmLink{Id: "1", Link: "https://github.com/Kwynto"},
		msmLink{Id: "2", Link: "https://github.com/Kwynto"},
		msmLink{Id: "3", Link: "https://github.com/Kwynto"},
		msmLink{Id: "4", Link: "https://github.com/Kwynto"},
		msmLink{Id: "5", Link: "https://github.com/Kwynto"},
		msmLink{Id: "6", Link: "https://github.com/Kwynto"},
		msmLink{Id: "7", Link: "https://github.com/Kwynto"},
		msmLink{Id: "8", Link: "https://github.com/Kwynto"},
		msmLink{Id: "9", Link: "https://github.com/Kwynto"},
		msmLink{Id: "10", Link: "https://github.com/Kwynto"},
		msmLink{Id: "11", Link: "https://github.com/Kwynto"},
		msmLink{Id: "12", Link: "https://github.com/Kwynto"},
		msmLink{Id: "13", Link: "https://github.com/Kwynto"},
		msmLink{Id: "14", Link: "https://github.com/Kwynto"},
		msmLink{Id: "15", Link: "https://github.com/Kwynto"},
		msmLink{Id: "16", Link: "https://github.com/Kwynto"},
		msmLink{Id: "17", Link: "https://github.com/Kwynto"},
		msmLink{Id: "18", Link: "https://github.com/Kwynto"},
		msmLink{Id: "19", Link: "https://github.com/Kwynto"},
	)

	// id := gosession.StartSecure(&w, r)

	app.render(w, r, "home.html", tD)
}
