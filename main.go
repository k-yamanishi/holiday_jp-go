package main

import (
	"io/ioutil"
	"os"
	"time"

	log "github.com/inconshreveable/log15"
	yaml "gopkg.in/yaml.v2"
)

//This go:generate command is not a comment. It is executed at build time.
//go:generate go-assets-builder -s="/data" -o holiday.go data

var (
	// DateFormat is "yyyy-MM-dd"
	DateFormat = "2006-01-02"

	// JST is Japan's time zone
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func main() {
	os.Exit(run())
}

// 0:not working day
// 1:normal day
// 2:error
func run() int {
	f, errYml := Assets.Open("/holidays.yml")
	if errYml != nil {
		log.Error("Open", "Error", errYml)
		return 2
	}
	defer f.Close()

	b, errRead := ioutil.ReadAll(f)
	if errRead != nil {
		log.Error("Read", "Error", errRead)
		return 2
	}

	holiday := make(map[interface{}]interface{})
	errYamlUnmarshal := yaml.Unmarshal(b, &holiday)
	if errYamlUnmarshal != nil {
		log.Error("YAML Unmarshal", "Error", errYamlUnmarshal)
		return 2
	}
	//test code
	//now := time.Date(2023, 5, 3, 9, 0, 0, 0, time.Local)
	now := time.Now().In(JST)
	d := now.Format(DateFormat)

	if _, ok := holiday[d]; ok {
		return 0
	}

	return 1
}
