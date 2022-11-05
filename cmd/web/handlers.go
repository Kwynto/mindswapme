package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Kwynto/gosession"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	tD := &templateData{User: "", Hash: "", Cart: []string{""}, Transitions: []string{""}}

	id := gosession.StartSecure(&w, r)

	transitions := id.Get("transitions")
	if transitions == nil {
		transitions = ""
	}
	transitions = fmt.Sprint(transitions, " ", r.RequestURI)
	id.Set("transitions", transitions)
	tStr := fmt.Sprintf("%v", transitions)
	tStrs := strings.Split(tStr, " ")
	tD.Transitions = tStrs

	cart := id.Get("cart")
	if cart == nil {
		cart = ""
		id.Set("cart", fmt.Sprint(cart))
		tD.Cart = []string{"There's nothing here yet."}
	} else {
		sCart := fmt.Sprint(cart)
		prods := strings.Split(sCart, " ")
		tD.Cart = prods
	}

	username := id.Get("username")
	if username != nil {
		tD.User = fmt.Sprint(username)
		app.render(w, r, "homeauth.page.html", tD)
	} else {
		app.render(w, r, "home.page.html", tD)
	}
}
