package reddit

import (
	"net/http"
	"net/url"
	"strings"
)

type V map[string]string

func (v V) Headers() http.Header {
	h := make(http.Header, len(v))
	for k, v := range v {
		h[k] = []string{v}
	}
	return h
}

func (v V) Encode() string {
	return v.encode(false)
}

func (v V) Query() string {
	return v.encode(true)
}

func (v V) EncodeReader() *strings.Reader {
	return strings.NewReader(v.Encode())
}

func (v V) encode(query bool) string {
	if v == nil {
		return ""
	}

	var buf strings.Builder

	if query {
		buf.WriteByte('?')
	}

	for k, v := range v {
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(v))
		buf.WriteByte('&')
	}

	return buf.String()
}
