package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mytho_game/database"
	"mytho_game/structure"
	"net/http"
	"strconv"
	"text/template"
)

func erreur(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Handling 500 errors
	_, err := template.ParseFiles("./nui/home.html")
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

	database.CreateJsonQuizz()

	fileServer := http.FileServer(http.Dir("./nui"))

	// http.Handle("/nui/", http.StripPrefix("/nui/", fileServer))
	http.Handle("/", fileServer)
	http.HandleFunc("/error", erreur)
	http.HandleFunc("/game", Game)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/room1", Room1)
	http.HandleFunc("/room2", Room2)
	http.HandleFunc("/manches", Manche)

	fmt.Printf("Starting server at port:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

var team2 structure.Team2
var team1 structure.Team1
var manche structure.Manche

func Home(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	t, err := template.ParseFiles("nui/home.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}
	t.ExecuteTemplate(w, "home", nil)

}

func Room1(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	t, err := template.ParseFiles("nui/room.html", "nui/assets/css/theme2.css")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}

	t.ExecuteTemplate(w, "room1", nil)

}

func Room2(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	t, err := template.ParseFiles("nui/room2.html", "nui/assets/css/theme2.css")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}
	choixAvatarTeam1 := r.FormValue("position")
	nameTeam1 := r.FormValue("name_team1")

	fmt.Printf("choixAvatar: %v\n", choixAvatarTeam1)
	fmt.Printf("nameTeam1: %v\n", nameTeam1)

	team1 = structure.Team1{
		Name:   nameTeam1,
		Avatar: "./data/" + choixAvatarTeam1 + ".png",
		Score:  15,
	}

	jsonData, err := json.Marshal(team1)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile("./json/team1.json", jsonData, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "room2", nil)

}

func Manche(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	t, err := template.ParseFiles("nui/manche.html", "nui/assets/css/theme2.css")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}

	choixAvatarTeam2 := r.FormValue("position")

	nameTeam2 := r.FormValue("name_team2")

	fmt.Printf("choixAvatarTeam2: %v\n", choixAvatarTeam2)
	fmt.Printf("nameTeam1: %v\n", nameTeam2)

	team2 = structure.Team2{
		Name:   nameTeam2,
		Avatar: "./data/" + choixAvatarTeam2 + ".png",
		Jeton:  15,
	}

	jsonData, err := json.Marshal(team2)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile("./json/team2.json", jsonData, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.PickQuestion()

	t.ExecuteTemplate(w, "manche", nil)

}

func Game(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Printf("team1: %v\n", team1)
	fmt.Printf("team2: %v\n", team2)

	manche.Tour, _ = strconv.Atoi(r.FormValue("position"))

	t, err := template.ParseFiles("nui/game.html", "nui/assets/css/game.css")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}

	err = t.ExecuteTemplate(w, "game", map[string]any{
		"Team1":  team1,
		"Team2":  team2,
		"Manche": manche.Tour,
	})
	if err != nil {
		log.Fatal(err)
	}
}
