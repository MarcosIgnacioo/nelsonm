package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// we put the handler as a method to the application because is cool and a way to make dependencies available to our handlers just in the application structure which we can put there when we initialize it in main and all of this without resorting to global variables or closures

// Q: Why is writer passed as value and request as pointer?
// request is passed as a pointer mainly because it is a big struct, and passing as a pointer makes stuff just faster
// but the reason for using writer as a value still confuses me mainly because i dont really know how does go implement the interface stuff
// here is a link to literally an explanation on that :3
// READ!!!
// https://research.swtch.com/interfaces
// Because it's an interface.
// An interface is a "fat-pointer", it's a two-word value with one pointer to the data and other pointer to a method table.
// lets implement them in c
func (app *Application) healthcheck_handler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"status": "ok"}`))
}

func (app *Application) image(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	file, handler, err := request.FormFile("image_field")

	if err != nil {
		http.Error(writer, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	file_name := handler.Filename
	dst, err := createFile(file_name)
	if err != nil {
		http.Error(writer, "Error creating the empty file", http.StatusInternalServerError)
		return
	}

	defer dst.Close()
	if _, err := dst.ReadFrom(file); err != nil {
		http.Error(writer, "Error writing the content of file", http.StatusInternalServerError)
	}

	msg := fmt.Sprintf("{\"status\": \"ok\", \"file_name\":\"%s\" }", file_name)
	writer.Write([]byte(msg))
}

func createFile(filename string) (*os.File, error) {
	// Create an uploads directory if it doesn’t exist
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	// Build the file path and create it
	dst, err := os.Create(filepath.Join("uploads", filename))
	if err != nil {
		return nil, err
	}

	return dst, nil
}
