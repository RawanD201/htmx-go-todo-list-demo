package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Title string
}

func main() {
	fmt.Println("Go app...")

	// handler function #1 - returns the index.html template, with Todo data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		todo := map[string][]Todo{
			"Todos": {
				{Title: "Check email and respond to urgent messages"},
				{Title: "Review calendar and schedule for the day"},
				{Title: "Complete a short but important task to get momentum going"},
			},
		}
		tmpl.Execute(w, todo)
	}

	// handler function #2 - returns the template block with the newly added todos, as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		// htmlStr := fmt.Sprintf("<li>%s </li>", title)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "=todo-list", Todo{Title: title})
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-todo/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
