package main

import (
	"fmt"
	"net/http"

	"github.com/Kwynto/gosession"
)

func (app *application) manage(w http.ResponseWriter, r *http.Request) {
	// Links := []MsmLink{
	// 	{Id: "1", Link: "https://github.com/Kwynto"},
	// 	{Id: "2", Link: "https://github.com/Kwynto"},
	// 	{Id: "3", Link: "https://github.com/Kwynto"},
	// 	{Id: "4", Link: "https://github.com/Kwynto"},
	// 	{Id: "5", Link: "https://github.com/Kwynto"},
	// 	{Id: "6", Link: "https://github.com/Kwynto"},
	// 	{Id: "7", Link: "https://github.com/Kwynto"},
	// 	{Id: "8", Link: "https://github.com/Kwynto"},
	// 	{Id: "9", Link: "https://github.com/Kwynto"},
	// 	{Id: "10", Link: "https://github.com/Kwynto"},
	// 	{Id: "11", Link: "https://github.com/Kwynto"},
	// 	{Id: "12", Link: "https://github.com/Kwynto"},
	// 	{Id: "13", Link: "https://github.com/Kwynto"},
	// 	{Id: "14", Link: "https://github.com/Kwynto"},
	// 	{Id: "15", Link: "https://github.com/Kwynto"},
	// 	{Id: "16", Link: "https://github.com/Kwynto"},
	// 	{Id: "17", Link: "https://github.com/Kwynto"},
	// 	{Id: "18", Link: "https://github.com/Kwynto"},
	// 	{Id: "19", Link: "https://github.com/Kwynto"},
	// 	{Id: "20", Link: "https://github.com/Kwynto"},
	// }
	tD := tmplManageData{User: "", Hash: "", Links: []MsmLink{
		{Id: "1", Link: "https://github.com/Kwynto"},
		{Id: "2", Link: "https://github.com/Kwynto"},
		{Id: "3", Link: "https://github.com/Kwynto"},
		{Id: "4", Link: "https://github.com/Kwynto"},
		{Id: "5", Link: "https://github.com/Kwynto"},
		{Id: "6", Link: "https://github.com/Kwynto"},
		{Id: "7", Link: "https://github.com/Kwynto"},
		{Id: "8", Link: "https://github.com/Kwynto"},
		{Id: "9", Link: "https://github.com/Kwynto"},
		{Id: "10", Link: "https://github.com/Kwynto"},
		{Id: "11", Link: "https://github.com/Kwynto"},
		{Id: "12", Link: "https://github.com/Kwynto"},
		{Id: "13", Link: "https://github.com/Kwynto"},
		{Id: "14", Link: "https://github.com/Kwynto"},
		{Id: "15", Link: "https://github.com/Kwynto"},
		{Id: "16", Link: "https://github.com/Kwynto"},
		{Id: "17", Link: "https://github.com/Kwynto"},
		{Id: "18", Link: "https://github.com/Kwynto"},
		{Id: "19", Link: "https://github.com/Kwynto"},
		{Id: "20", Link: "https://github.com/Kwynto"},
	}}

	id := gosession.StartSecure(&w, r)

	username := id.Get("username")
	if username != nil {
		tD.User = fmt.Sprint(username)
		app.render(w, r, "homeauth.html", tD) // app.render(w, r, "manage.html", tD)
	} else {
		app.render(w, r, "manage.html", tD) // http.Redirect(w, r, "/auth", http.StatusSeeOther)
	}
}

func (app *application) manageAuthPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("login")
	password := r.FormValue("password")

	id := gosession.StartSecure(&w, r)

	if username != "" && password != "" {
		pasHash := app.getMd5(password)
		id.Set("username", username)
		id.Set("hash", pasHash)
	}

	http.Redirect(w, r, "/manage", http.StatusSeeOther)
}

func (app *application) manageOutPage(w http.ResponseWriter, r *http.Request) {
	id := gosession.StartSecure(&w, r)
	id.Remove("username")
	id.Remove("hash")

	http.Redirect(w, r, "/manage", http.StatusSeeOther)
}
