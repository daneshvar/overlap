package api

import (
	"daneshvar/overlap/db"
	"daneshvar/overlap/log"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

func GetAll(ctx *fasthttp.RequestCtx) {
	if rec, err := db.ReadAll(); err != nil {
		log.Error(err)
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
