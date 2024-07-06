package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler handles the "/form" route in
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the incoming request
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST REQUEST SUCCESSFUL")
	// Get the value of the "name" form field
	name := r.FormValue("name")
	// Get the value of the "address" form field
	address := r.FormValue("address")
	// Print the retrieved form values to the response writer
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// helloHandler handles the "/hello" route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested URL path is "/hello"
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found Broski", http.StatusNotFound)
		return
	}
	// Check if the HTTP method is GET
	if r.Method != "GET" {
		http.Error(w, "404 Not Found Broski", http.StatusNotFound)
		return
	}
	// Print "hello!" to the response writer
	fmt.Fprintf(w, "hello!")
}

func main() {
	// Create a file server to serve static files from the "./static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	// Register the file server to handle requests for the root path "/"
	http.Handle("/", fileServer)
	// Register the formHandler function to handle requests for the "/form" path
	http.HandleFunc("/form", formHandler)
	// Register the helloHandler function to handle requests for the "/hello" path
	http.HandleFunc("/hello", helloHandler)

	// Print a message indicating that the server is starting on port 8080
	fmt.Printf("Starting server at port 8080\n")
	// Start the web server and log any errors
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
