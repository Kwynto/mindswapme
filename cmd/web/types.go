package main

import (
	"html/template"
	"log"
	"net/http"
)

// Creating an `application` structure to store the dependencies of the entire web application.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

// Structure for configuration
type Config struct {
	Addr      string
	StaticDir string
	Start     string
}

// Blocking direct access to the file system
type neuteredFileSystem struct {
	fs http.FileSystem
}

// Structure for the data template
type MsmLink struct {
	Id   string
	Link string
}

type tmplHomeData struct {
	isLinks bool
	Link1, Link2, Link3, Link4, Link5, Link6, Link7, Link8, Link9, Link10,
	Link11, Link12, Link13, Link14, Link15, Link16, Link17, Link18, Link19 MsmLink
}

type tmplManageData struct {
	User  string
	Hash  string
	Links []MsmLink
}
