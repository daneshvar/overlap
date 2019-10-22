package main

import (
	"daneshvar/overlap/api"
	"daneshvar/overlap/db"
	"daneshvar/overlap/log"
	"daneshvar/overlap/route"
	"os"
	"os/signal"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	bootstrap()

	route.Post("/", api.Add)
	route.Get("/", api.GetAll)
	route.Get("/unique", api.GetUnique)

	run(cfg.Addr)
	db.Close()
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	if fn := route.Route(string(ctx.URI().Path()), string(ctx.Method())); fn != nil {
		fn(ctx)
	} else {
		ctx.NotFound()
	}
}

func run(addr string) {
	http := fasthttp.Server{
		Handler:                       requestHandler,
		Name:                          "overlap",
		DisableHeaderNamesNormalizing: false,
		NoDefaultServerHeader:         true,
		TCPKeepalive:                  true,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	log.Info("HTTP: Running in ", addr)
	go func() {
		if err := http.ListenAndServe(addr); err != nil {
			if err.Error() == "HTTP: Server closed" {
				log.Info("Server closed")
			} else {
				log.Fatal(err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	<-quit
	log.Info("HTTP in Shutting down ...")
	<-time.After(5 * time.Second)
}
