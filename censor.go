package main

import (
	"bytes"
	"io"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func isText(MIMEtype string) bool {
	switch {
	case strings.HasPrefix(MIMEtype, "text/html"):
		return true
	case strings.HasPrefix(MIMEtype, "text/javascript"):
		return true
	case strings.HasPrefix(MIMEtype, "application/json"):
		return true
	case strings.HasPrefix(MIMEtype, "text/plain"):
		return true
	default:
		return false
	}
}

func editLinks(body []byte, originalUrl, originalCensor string) (string, error) {
	reader := bytes.NewReader(body)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	var errors = make([]error, 0)

	doc.Find("*[href],*[src],form[action]").Each(func(i int, s *goquery.Selection) {
		attr_name := "href"
		link, exists := s.Attr("href")
		if !exists {
			link, exists = s.Attr("src")
			if !exists {
				link, exists = s.Attr("action")
				if !exists {
					return
				} else {
					attr_name = "action"
				}
			} else {
				attr_name = "src"
			}
		}

		link, err = url.QueryUnescape(link)
		if err != nil {
			errors = append(errors, err)
			return
		}

		if strings.HasPrefix(link, "javascript:") {
			return
		}

		var rot13_buf strings.Builder
		c := strings.NewReader(link)
		r := rot13Reader{c}

		_, err = io.Copy(&rot13_buf, r)
		if err != nil {
			errors = append(errors, err)
			return
		}

		rot13_url, err := url.Parse(rot13_buf.String())
		if err != nil {
			errors = append(errors, err)
			return
		}

		originalUrlParsed, err := url.Parse(originalUrl)
		if err != nil {
			errors = append(errors, err)
			return
		}

		if rot13_url.Host == "" {
			rot13_url.Host = originalUrlParsed.Host
		}
		if rot13_url.Scheme == "" {
			rot13_url.Scheme = originalUrlParsed.Scheme
		}

		newLink := "//127.0.0.1:8080/proxy?u=" + url.QueryEscape(rot13_url.String()) + "&c=" + originalCensor

		s.SetAttr(attr_name, newLink)
	})

	if len(errors) != 0 {
		return "", errors[0]
	}

	new_body, err := doc.Html()
	if err != nil {
		return "", err
	}
	return new_body, nil
}

func censorBody(body string, censorRegex *regexp.Regexp) string {
	return censorRegex.ReplaceAllString(body, "censored")
}
