package api

import (
	"daneshvar/overlap/db"
	"daneshvar/overlap/log"
	"daneshvar/overlap/rect"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/valyala/fasthttp"
)

func Add(ctx *fasthttp.RequestCtx) {
	var data struct {
		Main  rect.Rect
		Input []rect.Rect
	}

	if err := json.Unmarshal(ctx.PostBody(), &data); err == nil {
		if err := adds(&data.Main, data.Input); err != nil {
			ctx.WriteString(`{"status": "Error"}`)
		} else {
			ctx.WriteString(`{"status": "OK"}`)
		}
	} else {
		log.Error(err)
		ctx.SetStatusCode(http.StatusBadRequest)
	}
}

func adds(main *rect.Rect, inputs []rect.Rect) error {
	wg := sync.WaitGroup{}
	w := db.NewWriter()

	for i := range inputs {
		wg.Add(1)
		go func(rc rect.Rect) {
			defer wg.Done()
			if main.IsOverlap(rc) {
				w.Add(&rc)
			}
		}(inputs[i])
	}
	wg.Wait()

	return w.Append()
}
