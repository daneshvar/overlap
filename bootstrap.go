package main

import (
	"daneshvar/overlap/db"
	"daneshvar/overlap/log"
	"encoding/json"
	"os"
)

type config struct {
	Addr string
	DB   string
}

var cfg = config{
	Addr: ":8080",
	DB:   "data.bin",
}

func bootstrap() {
	loadConfig("config.json")
}

func loadConfig(filename string) {
	if fileExists(filename) {
		log.Info("Load Config:", filename)
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
		log.Info("Load Config: Default Config")
	}

	if err := db.Init(cfg.DB); err != nil {
		log.Fatal(err)
	}
}
