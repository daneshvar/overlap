package main

import (
	"daneshvar/overlap/db"
	"daneshvar/overlap/rect"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

func add(ctx *fasthttp.RequestCtx) {
	var data struct {
		Main  rect.Rect
		Input []rect.Rect
	}

	log.Println(string(ctx.PostBody()))

	if err := json.Unmarshal(ctx.PostBody(), &data); err == nil {
		log.Println(data)
		if err := adds(&data.Main, data.Input); err != nil {
			log.Println(err)
			ctx.WriteString(`{"status": "Error"}`)
		} else {
			ctx.WriteString(`{"status": "OK"}`)
		}
	} else {
		log.Println(err)
		ctx.SetStatusCode(http.StatusBadRequest)
	}
}

func get(ctx *fasthttp.RequestCtx) {
	if rec, err := db.ReadAll(); err != nil {
		log.Println(err)
		ctx.WriteString(`{"status": "Error"}`)
	} else {
		b, err := json.Marshal(rec)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.Write(b)
	}
}

func getUnique(ctx *fasthttp.RequestCtx) {
	if uniques, err := db.ReadUnique(); err != nil {
		log.Println(err)
		ctx.WriteString(`{"status": "Error"}`)
	} else {
		var minReact rect.Rect
		var minTime int64 = math.MaxInt64
		found := false

		for rc, time := range uniques {
			if time > 0 && time < minTime {
				minTime = time
				minReact = rc
				found = true
			}
		}

		if found {
			rec := db.Record{
				X:      minReact.X,
				Y:      minReact.Y,
				Width:  minReact.Width,
				Height: minReact.Height,
				Time:   time.Unix(minTime, 0),
			}
			b, err := json.Marshal(rec)
			if err != nil {
				fmt.Println(err)
				return
			}
			ctx.Write(b)
		} else {
			ctx.WriteString(`{"status": "Error", "Message": "Not Found Unique"}`)
		}
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
				log.Println(rc)
				w.Add(rc)
			}
		}(inputs[i])
	}
	wg.Wait()

	return w.Save()
}
