package router

import (
	"bytes"
	"github.com/microcosm-cc/bluemonday"
	"io"
	"net/http"
	"strings"
)

const (
	DecodeOne      = "&#34;"
	DecodeTwo      = "&#39;"
	DecodeThree    = "&amp; "
	DecodeFour     = "&amp;amp; "
	NewDecodeOne   = "\""
	NewDecodeTwo   = "'"
	NewDecodeThree = "& "
	NewDecodeFour  = "& "
)

func SanitizeBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := SanitizeBody(r)
		if err != nil {
			http.Error(w, "error sanitizing request body", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SanitizeBody(req *http.Request) error {
	p := bluemonday.StrictPolicy()
	bodyRead, err := io.ReadAll(p.SanitizeReader(req.Body))
	if err != nil {
		return err
	}

	bodyRead = htmlDecode(bodyRead)

	req.Body = io.NopCloser(bytes.NewReader(bodyRead))
	return nil
}

func htmlDecode(body []byte) []byte {
	body = []byte(strings.Replace(string(body), DecodeOne, NewDecodeOne, -1))
	body = []byte(strings.Replace(string(body), DecodeTwo, NewDecodeTwo, -1))
	body = []byte(strings.Replace(string(body), DecodeThree, NewDecodeThree, -1))
	body = []byte(strings.Replace(string(body), DecodeFour, NewDecodeFour, -1))
	return body
}