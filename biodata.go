package main

import (
	"encoding/json"
	"io/ioutil"
)

type biodata struct {
	Id      string
	Name    string
	Address string
	Reason  string
}

func getBiodatas() (biodatas []biodata) {
	fileBytes, err := ioutil.ReadFile("./data/biodatas.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &biodatas)

	if err != nil {
		panic(err)
	}
	return biodatas
}
