package main

import (
	"daneshvar/overlap/db"
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Addr string
	DB   string
}

var cfg = config{
	Addr: ":80",
	DB:   "data.bin",
}

func bootstrap() {
	loadConfig("config.json")
}

func loadConfig(filename string) {
	if fileExists(filename) {
		log.Printf("Load Config: %s\n", filename)
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		dec := json.NewDecoder(f)
		err = dec.Decode(&cfg)
		_ = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Load Config: Default Config\n")
	}

	if err := db.Init(cfg.DB); err != nil {
		log.Fatalln(err)
	}
}
