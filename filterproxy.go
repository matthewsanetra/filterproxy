package main

import (
	"bytes"
	"log"

	"github.com/valyala/fasthttp"
)

const addr = "0.0.0.0:8080"

func main() {
	if err := fasthttp.ListenAndServe(addr, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	if bytes.HasPrefix(ctx.Path(), []byte("/proxy")) {
		handleProxyRequest(ctx)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Println("non /proxy request passed through, please fix in reverse proxy")
	}
}
