package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// var wg sync.WaitGroup

func serverStart(app *application, srv *http.Server, cfg *Config) {
	app.infoLog.Printf("Start server: %s", cfg.Addr)
	err := srv.ListenAndServe()
	app.errorLog.Println(err)
	// wg.Done()
}

func main() {
	// Reading flags from the application launch bar.
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
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
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}

	// Server structure with address, logger and routing
	// srv := &http.Server{
	// 	Addr:     cfg.Addr,
	// 	ErrorLog: errorLog,
	// 	Handler:  app.routes(),
	// }
	var srv *http.Server

	time.Sleep(time.Second)
	fmt.Println("The server is running. You can execute commands.")
	fmt.Print("> ")
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		command := sc.Text()
		switch command {
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
			go serverStart(app, srv, cfg)
		case "restart":
			// ReStarting the server
			srv.Shutdown(context.Background())
			app.infoLog.Println("HTTP server Shutdown.")
			srv = &http.Server{
				Addr:     cfg.Addr,
				ErrorLog: errorLog,
				Handler:  app.routes(),
			}
			go serverStart(app, srv, cfg)
		case "help":
			fmt.Println("This is helping.")
		default:
			fmt.Println("Unidentified command.")
		}
		time.Sleep(time.Second)
		fmt.Print("> ")
	}

	// wg.Wait()
}
