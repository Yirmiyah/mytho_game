package main

import (
	"fmt"
	"log"
	database "mytho/database"
	"net/http"
	"text/template"
)

func erreur(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Handling 500 errors
	_, err := template.ParseFiles("templates/index.html")
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

	database.CreateDataBase()
	defer database.Db.Close()
	http.Handle("/nui/", http.StripPrefix("/nui/", http.FileServer(http.Dir("./nui"))))

	http.HandleFunc("/", Game)
	http.HandleFunc("/error", erreur)
	fmt.Printf("Starting server at port:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func Game(w http.ResponseWriter, r *http.Request) {

	database.GeneratePlayers()
	fmt.Printf("database.GetPlayers(): %v\n", database.GetPlayers())

	t := template.New("game")
	t = template.Must(t.ParseFiles("nui/game.html"))
	err := t.ExecuteTemplate(w, "game", nil)
	if err != nil {
		log.Fatal(err)
	}

}
