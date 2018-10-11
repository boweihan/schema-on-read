package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: available"))
}

type Event struct {
	Type string
	Body string
}

func (e *Event) String() string {
	return fmt.Sprintf("Event read successfully. { Type: %v, Body: %s }",
		e.Type, e.Body)
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// parse the form
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() error: %v", err)
			return
		}

		// decode the body into an Event since it's a POST
		decoder := json.NewDecoder(r.Body)
		var event Event
		err := decoder.Decode(&event)
		if err != nil {
			panic(err)
		}

		// log the body as a response
		w.Write([]byte(event.String()))

	default:
		fmt.Fprintf(w, "Sorry, only POST is supported.")
	}
}

func main() {
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/event", eventHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
