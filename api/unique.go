package api

import (
	"daneshvar/overlap/db"
	"daneshvar/overlap/log"
	"daneshvar/overlap/rect"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/valyala/fasthttp"
)

func GetUnique(ctx *fasthttp.RequestCtx) {
	if uniques, err := db.ReadUnique(); err != nil {
		log.Error(err)
		ctx.WriteString(`{"status": "Error"}`)
	} else {
		var minReact rect.Rect
		var minTime int64 = math.MaxInt64
		found := false

		for rc, t := range uniques {
			if t > 0 && t < minTime {
				minTime = t
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
				Time:   time.Unix(0, minTime),
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
