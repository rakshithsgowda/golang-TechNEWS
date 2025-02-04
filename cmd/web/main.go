package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
	session *scs.SessionManager
}

type server struct {
	port string
	host string
	url  string
}

func main() {
	server := server{
		host: "localhost",
		port: "8009",
		url:  "http://localhost:8009",
	}
	app := &application{
		server:  server,
		appName: "hackernews",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Llongfile),
	}

	// JET template initalization
	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"))
	}

	// Sessions initialization
	// app.session = scs.New()
	// app.session.Lifetime = 24 * time.Hour
	// app.session.Cookie.Persist = true
	// app.session.Cookie.Name = app.appName
	// app.session.Cookie.Domain = app.server.host
	// app.session.Cookie.SameSite = http.SameSiteStrictMode

	// serve the App
	if err := app.listenAndServe(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello world!")
}
