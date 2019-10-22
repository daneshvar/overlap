package main

import (
	"daneshvar/overlap/db"
	"daneshvar/overlap/route"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	bootstrap()

	route.Post("/", add)
	route.Get("/", get)
	route.Get("/unique", getUnique)

	run(cfg.Addr)
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	if fn := route.Route(string(ctx.URI().Path()), string(ctx.Method())); fn != nil {
		fn(ctx)
	} else {
		ctx.NotFound()
	}
}

func run(addr string) {
	defer db.Close()

	http := fasthttp.Server{
		Handler:                       requestHandler,
		Name:                          "overlap",
		DisableHeaderNamesNormalizing: false,
		NoDefaultServerHeader:         true,
		TCPKeepalive:                  true,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	log.Println("HTTP: Running in ", addr)
	go func() {
		if err := http.ListenAndServe(addr); err != nil {
			if err.Error() == "HTTP: Server closed" {
				log.Println("Server closed")
			} else {
				log.Fatalln(err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	<-quit
	log.Println("HTTP in Shutting down ...")
	<-time.After(2 * time.Second)
}
