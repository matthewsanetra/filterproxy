package main

import (
	"log"
	"strings"

	"github.com/valyala/fasthttp"
)

func main() {
	if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())

	// TODO: SERVE INDEX STATIC FILE WITH NGINX, ONLY PASSTHROUGH /proxy

	if strings.HasPrefix(path, "/proxy") {
		handleProxyRequest(ctx)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Println("non /proxy request passed through, please fix in reverse proxy")
	}
}
