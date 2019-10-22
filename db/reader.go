package db

import (
	"bytes"
	"daneshvar/overlap/rect"
	"io"
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

// Walk of all records
func Walk(fn func(*rect.Rect, int64)) error {
	lock.Lock()
	defer lock.Unlock()

	var size int64

	if ofs, err := file.Seek(0, io.SeekEnd); err != nil {
		return err
	} else {
		size = ofs
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return err
	}

	buf := make([]byte, size)
	if _, err := file.Read(buf); err != nil && err != io.EOF {
		return err
	}

	r := bytes.NewBuffer(buf)
	var rc rect.Rect
	var t int64

	for i := size; i > 0; i -= RecordSize {
		if err := read(r, &rc, &t); err != nil {
			return err
		}

		fn(&rc, t)
	}

	return nil
}

// ReadAll returns All records
func ReadAll() ([]Record, error) {
	rec := make([]Record, 0)
	err := Walk(func(rc *rect.Rect, t int64) {
		rec = append(rec, Record{
			X: rc.X, Y: rc.Y, Width: rc.Width, Height: rc.Height,
			Time: time.Unix(0, t),
		})
	})

	if err != nil {
		return nil, err
	}

	return rec, nil
}

// ReadUnique returns map[Rect]Time of all Rect but time value for duplicates Rect is zero
func ReadUnique() (map[rect.Rect]int64, error) {
	rec := make(map[rect.Rect]int64, 0)
	err := Walk(func(rc *rect.Rect, t int64) {
		if v, ok := rec[*rc]; ok {
			if v > 0 {
				rec[*rc] = 0
			}
		} else {
			rec[*rc] = t
		}
	})

	if err != nil {
		return nil, err
	}

	return rec, nil
}
