package main

import (
	asciiartweb "asciiartweb/functions"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Title  string
	Output string
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound) // 404 Not Found (generic for the path)
		return
	}

	if r.Method == http.MethodPost {
		// Handle form submission (POST request)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest) // 400 Bad Request
			return
		}

		// Get the banner
		banner := "standard"
		banner = r.FormValue("banner") // Access form data by name attribute

		// Attempt to read the banner file
		content, err := os.ReadFile("./banners/" + banner + ".txt")
		if err != nil {
			http.Error(w, "Banner not found", http.StatusNotFound) // 404 Not Found (specific for a file)
			return
		}

		// Split the file content into individual lines
		lines := strings.Split(string(content), "\n")

		input := r.FormValue("input") // Access form data by name attribute
		output := asciiartweb.PrintArt(input, lines)

		data := PageData{
			Title:  "ASCII Art Web",
			Output: output,
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError) // 500 Internal Server Error
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError) // 500 Internal Server Error
			return
		}
	} else if r.Method == http.MethodGet {
		// Handle initial GET request (show form)
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError) // 500 Internal Server Error
			return
		}
		data := PageData{
			Title:  "ASCII Art Web",
			Output: "", // Set an empty output initially
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError) // 500 Internal Server Error
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // 405 Method Not Allowed
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil) // nil uses the default ServeMux
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
