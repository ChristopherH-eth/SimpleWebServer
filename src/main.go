package main

/**
 * @file main.go
 * @author Original author: Free Code Camp
 *		   Changes made by 0xChristopher for learning purposes
 *
 * This application demonstrates simple web server functionality by handling multiple routes and printing
 * messages to the screen based on user form submission.
 */

import (
	"fmt"
	"log"
	"net/http"
)

// The formHandler() function takes information submitted in a form by the user and prints it to the screen.
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s", address)
}

// The helloHandler() function handles the "/hello" route and prints "Hello!" to the screen.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Don't allow data requests
	if r.Method != "GET" {
		http.Error(w, "Method is Not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	// Routes to be handled
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Establish web server
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
