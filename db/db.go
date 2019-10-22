package db

import (
	"log"
	"os"
	"sync"
)

var file *os.File
var lock sync.RWMutex

func Init(filename string) error {
	if f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644); err != nil {
		return err
	} else {
		file = f
	}

	log.Println("Open Database in: ", filename)
	return nil
}

func Close() {
	file.Close()
}
