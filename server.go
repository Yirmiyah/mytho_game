package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func erreur(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Handling 500 errors
	_, err := template.ParseFiles("./nui/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

}

func main() {
	fileServer := http.FileServer(http.Dir("./nui"))

	http.Handle("/", fileServer)
	http.HandleFunc("/error", erreur)
	http.HandleFunc("/game", Game)

	fmt.Printf("Starting server at port:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func Game(w http.ResponseWriter, r *http.Request) {

}
