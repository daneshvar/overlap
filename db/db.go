package db

import (
	"daneshvar/overlap/log"
	"daneshvar/overlap/rect"
	"encoding/binary"
	"io"
	"os"
	"sync"
)

var file *os.File
var lock sync.Mutex

func Init(filename string) error {
	if f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644); err != nil {
		return err
	} else {
		file = f
	}

	log.Info("Open Database in: ", filename)
	return nil
}

func Close() {
	log.ErrorIf(file.Close())
}

func read(r io.Reader, rc *rect.Rect, time *int64) error {
	if err := binary.Read(r, binary.BigEndian, &rc.X); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &rc.Y); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &rc.Width); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &rc.Height); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, time); err != nil {
		return err
	}

	return nil
}

func write(w io.Writer, rc *rect.Rect, time int64) error {
	if err := binary.Write(w, binary.BigEndian, rc.X); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, rc.Y); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, rc.Width); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, rc.Height); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, time); err != nil {
		return err
	}

	return nil
}
