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
		ctx.SetBody([]byte("error writing from rot13reader to string builder buffer for regex, error logged"))
		log.Println(err)
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
		ctx.SetBody([]byte("error writing from rot13reader to string builder buffer for url, error logged"))
		log.Println(err)
		return
	}
	url := url_buf.String()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	fasthttp.Do(req, resp)

	contentType := string(resp.Header.ContentType())

	ctx.SetContentType(contentType)
	ctx.SetStatusCode(resp.StatusCode())

	if !isText(contentType) {
		ctx.SetBody(resp.Body())
		return
	}

	body, err := editLinks(resp.Body(), string(url_rot13), string(censored_expressions_rot13))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte("error editing links, error logged"))
		log.Println(err)
		return
	}

	body = censorBody(body, censor)
	ctx.SetBody([]byte(body))
}
