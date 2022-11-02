package jMU

import (
	"encoding/json"
	"log"
	"os"
)

type Details struct {
	Name   string
	Mobile float64
	Email  string
}

//doubt on marshalling multiple values declared on same data type

func Marshalling() {
	_, err := os.Create("scrap.txt")
	if err != nil {
		log.Println(err)
	}
	Person1 := Details{
		Name:   "teja",
		Mobile: 9863733828,
		Email:  "Teja@kit.com",
	}

	Jm, err := json.Marshal(Person1)
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile("scrap.txt", Jm, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}
