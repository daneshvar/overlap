package db

import (
	"bytes"
	"daneshvar/overlap/rect"
	"encoding/binary"
	"io"
	"os"
	"time"
)

type Record struct {
	X      int16     `json:"x"`
	Y      int16     `json:"y"`
	Width  int16     `json:"width"`
	Height int16     `json:"height"`
	Time   time.Time `json:"time"`
}

const RecordSize = 16

func ReadAll() ([]Record, error) {
	lock.RLock()
	defer lock.RUnlock()

	var size int64

	if ofs, err := file.Seek(0, os.SEEK_END); err != nil {
		return nil, err
	} else {
		size = ofs
	}

	if _, err := file.Seek(0, os.SEEK_SET); err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	if _, err := file.Read(buf); err != nil && err != io.EOF {
		return nil, err
	}

	io := bytes.NewBuffer(buf)
	rec := make([]Record, 0)
	for i := size; i > 0; i -= RecordSize {
		var rc rect.Rect
		var t int64
		binary.Read(io, binary.BigEndian, &rc.X)      // 2
		binary.Read(io, binary.BigEndian, &rc.Y)      // 2
		binary.Read(io, binary.BigEndian, &rc.Width)  // 2
		binary.Read(io, binary.BigEndian, &rc.Height) // 2
		binary.Read(io, binary.BigEndian, &t)         // 8

		rec = append(rec, Record{
			X: rc.X, Y: rc.Y, Width: rc.Width, Height: rc.Height,
			Time: time.Unix(t, 0),
		})
	}

	return rec, nil
}

func ReadUnique() (map[rect.Rect]int64, error) {
	lock.RLock()
	defer lock.RUnlock()

	var size int64

	if ofs, err := file.Seek(0, os.SEEK_END); err != nil {
		return nil, err
	} else {
		size = ofs
	}

	if _, err := file.Seek(0, os.SEEK_SET); err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	if _, err := file.Read(buf); err != nil && err != io.EOF {
		return nil, err
	}

	io := bytes.NewBuffer(buf)
	rec := make(map[rect.Rect]int64, 0)
	for i := size; i > 0; i -= RecordSize {
		var rc rect.Rect
		var t int64
		binary.Read(io, binary.BigEndian, &rc.X)      // 2
		binary.Read(io, binary.BigEndian, &rc.Y)      // 2
		binary.Read(io, binary.BigEndian, &rc.Width)  // 2
		binary.Read(io, binary.BigEndian, &rc.Height) // 2
		binary.Read(io, binary.BigEndian, &t)         // 8

		if v, ok := rec[rc]; ok {
			if v > 0 {
				rec[rc] = 0
			}
		} else {
			rec[rc] = t
		}
	}

	return rec, nil
}
