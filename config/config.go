package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadConfig(parseTo interface{}, names ...string) {
	var name string
	if len(names) == 0 {
		name = "config.json"
	} else {
		name = names[0]
	}

	d, err := ioutil.ReadFile(name)

	if err != nil {
		log.Panicln(err)
	}

	if err := json.Unmarshal(d, parseTo); err != nil {
		log.Panicln(err)
	}

	return
}
