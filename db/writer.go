package db

import (
	"bytes"
	"daneshvar/overlap/log"
	"daneshvar/overlap/rect"
	"io"
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
		now: time.Now().UnixNano(),
	}
}

func (w *Writer) Add(rc *rect.Rect) {
	w.lock.Lock()
	defer w.lock.Unlock()
	log.ErrorIf(write(&w.buf, rc, w.now))
}

func (w *Writer) Append() error {
	lock.Lock()
	defer lock.Unlock()

	if _, err := file.Seek(0, io.SeekEnd); err != nil {
		return err
	}

	_, err := file.Write(w.buf.Bytes())
	return err
}
