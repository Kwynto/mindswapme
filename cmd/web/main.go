package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var app *application
var srv *http.Server

func main() {
	// Reading flags from the application launch bar.
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.Start, "start", "no", "Do I need to start the server")
	flag.Parse()

	// Creating loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize the template cache.
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize the structure with the application dependencies.
	app = &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}

	if strings.ToLower(cfg.Start) == "yes" {
		// Starting the server
		srv = &http.Server{
			Addr:     cfg.Addr,
			ErrorLog: errorLog,
			Handler:  app.routes(),
		}
		go app.serverStart(srv)
	}

	time.Sleep(time.Second)
	fmt.Println("The server is running. You can execute commands.")
	fmt.Println("!!! Use the 'help' command to get a list of commands.")
	fmt.Print("> ")

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		command := sc.Text()
		switch command {
		case "help":
			data, err := os.ReadFile("./cmd/web/help.txt")
			if err != nil {
				app.errorLog.Fatal(err)
			}
			os.Stdout.Write(data)
		case "overload":
			templateCache, err = newTemplateCache("./ui/html/")
			if err != nil {
				app.errorLog.Fatal(err)
			}
			app.templateCache = templateCache
			app.infoLog.Println("Templates have been reloaded.")
		case "stop":
			// Stoping the server
			srv.Shutdown(context.Background())
			app.infoLog.Println("HTTP server Shutdown.")
		case "start":
			// Starting the server
			srv = &http.Server{
				Addr:     cfg.Addr,
				ErrorLog: errorLog,
				Handler:  app.routes(),
			}
			go app.serverStart(srv)
		case "restart":
			// ReStarting the server
			srv.Shutdown(context.Background())
			app.infoLog.Println("HTTP server Shutdown.")
			srv = &http.Server{
				Addr:     cfg.Addr,
				ErrorLog: errorLog,
				Handler:  app.routes(),
			}
			go app.serverStart(srv)
		default:
			fmt.Println("Unidentified command.")
		}
		time.Sleep(time.Second)
		fmt.Print("> ")
	}
}
