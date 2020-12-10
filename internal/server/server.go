package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"rvdl/pkg/util"
)

var isTesting = util.EnvBool("RVDL_TESTING", false)

// TODO: There has to be a better way to do this
func redirect(res http.ResponseWriter, req *http.Request, url string, code int) {
	if isTesting {
		http.Redirect(res, req, "https://"+req.Host+"/"+url, code)
	} else {
		http.Redirect(res, req, url, code)
	}
}

// TODO: There has to be a better way to do this
func transform(res http.ResponseWriter, req *http.Request) {
	if isTesting {
		u, err := url.Parse(req.URL.String()[1:])
		if err != nil {
			fmt.Println(err)
			return
		}
		req.URL = u
	} else {
		req.URL.Scheme = "https"
		req.URL.Host = req.Host
	}
}

func ServeHTTP(res http.ResponseWriter, req *http.Request) {
	transform(res, req)

	log.Printf("%s %s", req.Method, req.URL)

	switch req.URL.Path {
	case "/":
		home(res, req)
	case "/favicon.ico":
		favicon(res, req)
	default:
		handleRvdl(res, req)
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Cache-Control", "public, max-age=86400")
	//if req.URL.Host != "redditvideodownload.com" {
	//	redirect(res, req, "https://redditvideodownload.com/", http.StatusFound)
	//	return
	//}
	res.Header().Set("Content-Type", "text/plain")
	http.ServeFile(res, req, "web/static/index.txt")
}

func favicon(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Cache-Control", "public, max-age=86400")
	res.WriteHeader(http.StatusNoContent)
}
