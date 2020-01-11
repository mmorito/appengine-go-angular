package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"server/src/api/v1/hello"
	"server/src/utilities/db"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	http.Handle("/", e)

	// database
	db := db.Init()
	defer db.Close()

	// routes
	g := e.Group("/api/v1")

	i := g.Group("/hello")
	i.GET("", hello.Get)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}
