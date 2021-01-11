package main

import (
	"encoding/json"
	"fmt"
	"github.com/koyachi/go-nude"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)


type Image struct {
	Name string `json:"Name"`
	Size int64 `json:"size"`
	Nudity bool `json:"isNude"`
}

// let's declare a global Images array
// that we can then populate in our main function
// to simulate a database
var Images []Image

func Greeting(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "{'message': 'Welcome to is nude API', 'version': 1.0}")
}

func main() {

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/", Greeting)
	http.HandleFunc("/upload", uploadHandler)
	log.Fatal(s.ListenAndServe())
}


// Compile templates on start of the application
var templates = template.Must(template.ParseFiles("public/upload.html"))

// Display the named template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	imagePath := handler.Filename

	isNude, err := nude.IsNude(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("isNude: %+v\n", isNude)
	image := Image{handler.Filename, handler.Size, isNude}

	js, err := json.Marshal(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "upload", nil)
	case "POST":
		uploadFile(w, r)
	}
}