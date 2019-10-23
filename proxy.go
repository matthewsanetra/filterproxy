package main

import (
	"bytes"
	"io"
	"log"
	"strings"

	"regexp"

	"github.com/valyala/fasthttp"
)

func handleProxyRequest(ctx *fasthttp.RequestCtx) {
	url_rot13 := ctx.QueryArgs().Peek("u")
	censored_expressions_rot13 := ctx.QueryArgs().Peek("c")

	var censored_buf strings.Builder
	c := bytes.NewReader(censored_expressions_rot13)
	r := rot13Reader{c}

	_, err := io.Copy(&censored_buf, r)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte("unable to write from rot13reader to string builder buffer for regex"))
		return
	}
	censored_expression := censored_buf.String()
	censor, err := regexp.Compile(censored_expression)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte("invalid censor regex"))
		return
	}

	var url_buf strings.Builder
	c = bytes.NewReader(url_rot13)
	r = rot13Reader{c}

	_, err = io.Copy(&url_buf, r)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte("unable to write from rot13reader to string builder buffer for url"))
		return
	}
	url := url_buf.String()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	fasthttp.Do(req, resp)

	ctx.SetContentType(string(resp.Header.ContentType()))
	ctx.SetStatusCode(resp.StatusCode())

	body, err := censorBody(resp.Body(), censor, string(url_rot13), string(censored_expressions_rot13))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte("unable to censor, error logged"))
		log.Println(err)
		return
	}

	ctx.SetBody([]byte(body))
}
