package main

import (
	"bytes"
	"log"

	"github.com/valyala/fasthttp"
)

const addr = "127.0.0.1:8080"

func main() {
	if err := fasthttp.ListenAndServe(addr, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	// TODO: SERVE INDEX STATIC FILE WITH NGINX, ONLY PASSTHROUGH /proxy

	if bytes.HasPrefix(ctx.Path(), []byte("/proxy")) {
		handleProxyRequest(ctx)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Println("non /proxy request passed through, please fix in reverse proxy")
	}
}
