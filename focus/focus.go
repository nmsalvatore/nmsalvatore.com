package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

type Focus struct {
	options Options
}

func (f Focus) Work() string {
	return RandomOption(f.options.Type.Work)
}

func (f Focus) Movement() string {
	return RandomOption(f.options.Type.Movement)
}

type Options struct {
	Type struct {
		Work     []string `json:"work"`
		Movement []string `json:"movement"`
	} `json:"options"`
}

type Result struct {
	Work     string `json:"work"`
	Movement string `json:"movement"`
}

func main() {
	options := ReadOptionsAsJSON()
	focus := Focus{options}
	result := Result{
		focus.Work(),
		focus.Movement(),
	}
	WriteResultsAsJSON(result)
}

func ReadOptionsAsJSON() Options {
	contents, err := os.ReadFile("./options.json")
	if err != nil {
		log.Fatal(err)
	}

	var options Options
	err = json.Unmarshal(contents, &options)
	if err != nil {
		log.Fatal(err)
	}
	return options
}

func RandomOption(options []string) string {
	return options[rand.Intn(len(options))]
}

func WriteResultsAsJSON(result Result) {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("../public/results.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
