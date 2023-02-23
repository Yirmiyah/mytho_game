package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mytho_game/structure"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var quizz structure.Quizz

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateRandomName() string {

	lengthString := 6
	rand.Seed(time.Now().UnixNano())
	randomString := RandStringBytes(lengthString)

	return randomString
}

func CreateJsonQuizz() {

	f, err := excelize.OpenFile("database/Questions_Reponses_SEO.xlsx")
	if err != nil {
		fmt.Println("Erreur ouverture fichier Excel")
		log.Fatal(err)
	}

	for i := 2; i < 94; i++ {
		quizz.Question = append(quizz.Question, f.GetCellValue("Sheet1", "A"+strconv.Itoa(i)))
		quizz.Response = append(quizz.Response, f.GetCellValue("Sheet1", "B"+strconv.Itoa(i)))
		quizz.Level = append(quizz.Level, f.GetCellValue("Sheet1", "C"+strconv.Itoa(i)))
	}

	jsonData, err := json.Marshal(quizz)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile("./json/quizz.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func PickQuestion() {

	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(93)
	fmt.Printf("number: %v\n", number)
	fmt.Printf("quizz.Question: %v\n", quizz.Question[number])
	fmt.Printf("quizz.Response: %v\n", quizz.Response[number])
	fmt.Printf("quizz.Level: %v\n", quizz.Level[number])
}
