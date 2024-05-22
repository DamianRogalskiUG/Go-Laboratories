package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type SharkAttack struct {
	Date     string `json:"date"`
	Year     string `json:"year"`
	Type     string `json:"type"`
	Country  string `json:"country"`
	Area     string `json:"area"`
	Location string `json:"location"`
}

var attacks []SharkAttack

func loadSampleData() {
	file, err := os.Open("global-shark-attack.json")
	if err != nil {
		fmt.Println("Error occured while opening the file:", err)
		return
	}
	defer file.Close()

	var data []SharkAttack
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Error occured while decoding the JSON:", err)
		return
	}

	for i := 0; i < 10; i++ {
		index := rand.Intn(len(data))
		attacks = append(attacks, data[index])
	}
}

func main() {
	loadSampleData()

	for _, attack := range attacks {
		fmt.Printf("%s: %s, %s, %s\n", attack.Date, attack.Type, attack.Country, attack.Area)
	}
}
