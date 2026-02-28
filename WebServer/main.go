package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

// UserData holds the data from the form
type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	// In-memory storage for user data
	users []UserData
	// Mutex to handle concurrent writes to the file
	mutex = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/", serveForm)
	http.HandleFunc("/submit", handleForm)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func serveForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the form values
	name := r.FormValue("name")
	email := r.FormValue("email")

	// Create a new UserData instance
	newUser := UserData{Name: name, Email: email}

	// Lock the mutex to ensure safe concurrent file writes
	mutex.Lock()
	defer mutex.Unlock()

	// Append data to JSON file
	if err := appendToJSON("data.json", newUser); err != nil {
		http.Error(w, "Error saving data to JSON", http.StatusInternalServerError)
		return
	}

	// Append data to CSV file
	if err := appendToCSV("data.csv", newUser); err != nil {
		http.Error(w, "Error saving data to CSV", http.StatusInternalServerError)
		return
	}

	// Add the new user to the in-memory slice
	users = append(users, newUser)

	// Send a success response
	fmt.Fprintf(w, "<h1>Data stored successfully!</h1><a href=\"/\">Go back</a>")
}

func appendToJSON(filename string, data UserData) error {
	// Read existing data from the file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var existingData []UserData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&existingData); err != nil && err.Error() != "EOF" {
		// Ignore EOF error for empty file
	}

	// Append new data
	existingData = append(existingData, data)

	// Go back to the beginning of the file to overwrite
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(existingData)
}

func appendToCSV(filename string, data UserData) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Check if the file is new to write the header
	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.Size() == 0 {
		if err := writer.Write([]string{"Name", "Email"}); err != nil {
			return err
		}
	}

	return writer.Write([]string{data.Name, data.Email})
}