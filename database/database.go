package database

import (
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func CreateDataBase() {
	var err error
	Db, err = sql.Open("sqlite3", "./usersDB.db")
	if err != nil {
		fmt.Println("Erreur ouverture de la base de donnée à la création de la table:")
		log.Fatal(err)
	}

	_, err = Db.Exec(`CREATE TABLE IF NOT EXISTS players
        (id INTEGER PRIMARY KEY AUTOINCREMENT,
		team1 TEXT DEFAULT '',
		team2 TEXT DEFAULT '',
        player1 TEXT DEFAULT '',
		player2 TEXT DEFAULT '',
		player2 TEXT DEFAULT '',
		player4 TEXT DEFAULT '',
		player5 TEXT DEFAULT '',
		player6 TEXT DEFAULT '',
		coinTeam1 INTEGER DEFAULT 15,
		coinTeam2 INTEGER DEFAULT 15,
        masterName TEXT DEFAULT '',
        nbManche INTEGER DEFAULT 0
        )`)
	if err != nil {
		log.Println("erreur creation de table users")
		log.Fatal(err)
	}

	_, err = Db.Exec(`CREATE TABLE IF NOT EXISTS card
	(id INTEGER PRIMARY KEY AUTOINCREMENT,
		questions TEXT DEFAULT '',
		responses TEXT DEFAULT ''
	) `)
	if err != nil {
		log.Fatal("erreur creation de table players")
		log.Fatal(err)
	}

}

func GeneratePlayers() {

	_, err := Db.Exec("INSERT INTO players (player1) VALUES ? ", GenerateRandomName())
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec("INSERT INTO players (player2) VALUES ? ", GenerateRandomName())
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec("INSERT INTO players (player3) VALUES ? ", GenerateRandomName())
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec("INSERT INTO players (player4) VALUES ? ", GenerateRandomName())
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec("INSERT INTO players (player5) VALUES ? ", GenerateRandomName())
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec("INSERT INTO players (player6) VALUES ? ", GenerateRandomName())
	if err != nil {
		log.Fatal(err)
	}

}

func GetPlayers() map[string]string {

	result := make(map[string]string, 6)
	rows, err := Db.Query("SELECT player1,player2,player3,player4,player5,player6 FROM players ")
	if err != nil {
		fmt.Println("Erreur function GetPlayers selecting element failed:")
		log.Fatal(err)
	}
	var players1, players2, players3, players4, players5, players6 string
	for rows.Next() {
		err = rows.Scan(&players1, &players2, &players3, &players4, &players5, &players6)
		if err != nil {
			log.Fatal(err)
		}
	}

	result["players1"] = players1
	result["players2"] = players2
	result["players3"] = players3
	result["players4"] = players4
	result["players5"] = players5
	result["players6"] = players6

	return result

}
