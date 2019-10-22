package db

import (
	"bytes"
	"daneshvar/overlap/rect"
	"encoding/binary"
	"os"
	"sync"
	"time"
)

type Writer struct {
	now  int64
	lock sync.Mutex
	buf  bytes.Buffer
}

func NewWriter() *Writer {
	return &Writer{
		now: time.Now().Unix(),
	}
}

func (w *Writer) Add(rc rect.Rect) {
	w.lock.Lock()
	defer w.lock.Unlock()

	binary.Write(&w.buf, binary.BigEndian, rc.X)      // 2
	binary.Write(&w.buf, binary.BigEndian, rc.Y)      // 2
	binary.Write(&w.buf, binary.BigEndian, rc.Width)  // 2
	binary.Write(&w.buf, binary.BigEndian, rc.Height) // 2
	binary.Write(&w.buf, binary.BigEndian, w.now)     // 8
}

func (w *Writer) Save() error {
	lock.Lock()
	defer lock.Unlock()

	if _, err := file.Seek(0, os.SEEK_END); err != nil {
		return err
	}

	_, err := file.Write(w.buf.Bytes())
	return err
}
