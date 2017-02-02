package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
)

var hitCounter uint
var signatures []Signature

// Signature represents a signature!
type Signature struct {
	Name      string
	Timestamp time.Time
}

func guestbookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle a new signature being posted

		r.ParseForm()
		name := r.Form.Get("name")
		name = strings.TrimSpace(name)

		if name != "" {
			signatures = append(signatures, Signature{
				Name:      name,
				Timestamp: time.Now(),
			})
		}

	}

	// Load the template
	t, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		http.Error(w, "Error loading template", 500)
		log.Printf("Error loading template: %s", err)
		return
	}

	// Increment the hit counter
	hitCounter = hitCounter + 1

	data := struct {
		Hits       uint
		Signatures []Signature
	}{hitCounter, signatures}

	// Render the page
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error generating page", 500)
		log.Printf("Error generating page: %s", err)
	}
}

func getMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", guestbookHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux
}

func main() {
	http.Handle("/", handlers.LoggingHandler(os.Stdout, getMux()))
	panic(http.ListenAndServe(":8080", nil))
}
