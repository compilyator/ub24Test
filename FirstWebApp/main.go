package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Create a simple form template for the GET request
var formTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Simple Go WebApp</title>
</head>
<body>
    <h1>Submit Your Name</h1>
    <form action="/submit" method="post">
        <label for="name">Name:</label>
        <input type="text" id="name" name="name" required>
        <input type="submit" value="Submit">
    </form>
</body>
</html>
`

// Handler for the GET request
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the form to the user
	fmt.Fprintf(w, formTemplate)
}

// Handler for the POST request
func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Get the name from the form
	name := r.FormValue("name")

	// Display a response with the submitted name
	fmt.Fprintf(w, "<h1>Hello, %s!</h1>", template.HTMLEscapeString(name))
}

func main() {
	// Handle the home route (GET)
	http.HandleFunc("/", homeHandler)

	// Handle the form submission (POST)
	http.HandleFunc("/submit", submitHandler)

	// Start the server on port 8080
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
