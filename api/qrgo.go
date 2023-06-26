package qrgo

import (
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/skip2/go-qrcode"
)

// ErrMissingParam is returned if path param evaluation results empty
var ErrMissingParam = errors.New("no content provided for encoding")

// ErrEncodingFailed is returned if the query couldn't be successfully encoded
var ErrEncodingFailed = errors.New("couldn't encode provided content")

var plainTextUserAgents = []string{"curl", "wget", "fetch", "httpie"}

var tmpl = `<!DOCTYPE html>
<html>
<head>
<title>qrgo</title>
<script defer src="/_vercel/insights/script.js"></script>
</head>
<body style='background-color: black;color:white;'>
<a href="https://github.com/elsesiy/qrgo" class="github-corner" aria-label="View source on GitHub"><svg width="80" height="80" viewBox="0 0 250 250" style="fill:#fff; color:#151513; position: absolute; top: 0; border: 0; right: 0;" aria-hidden="true"><path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path><path d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2" fill="currentColor" style="transform-origin: 130px 106px;" class="octo-arm"></path><path d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z" fill="currentColor" class="octo-body"></path></svg></a><style>.github-corner:hover .octo-arm{animation:octocat-wave 560ms ease-in-out}@keyframes octocat-wave{0%,100%{transform:rotate(0)}20%,60%{transform:rotate(-25deg)}40%,80%{transform:rotate(10deg)}}@media (max-width:500px){.github-corner:hover .octo-arm{animation:none}.github-corner .octo-arm{animation:octocat-wave 560ms ease-in-out}}</style>

{{ if .QRCode }}
    <img src="data:image/png;base64,{{ .QRCode }}" alt="">
{{ else }}
    <pre>
Usage is simple, just append the content to be encoded to the Path, i.e.

$ curl -L qrgo.elsesiy.com/abc
    </pre>

    <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAAAz0lEQVR42uyX262CMQyDu0H239Ib9Ogk6Q0KvEJsg6q/0fdkNbcmSZIk/Yisu+BfWDc6AHmaQ1gxNsAN8nPdiAG3SsD80QKJ4T9z3mRWdSBL57TpVaUtDiwQnxpsaSDNgffVZs82sQAej6zpo6mAD4hq4b0kubN+0ADJnA+HELCIWQ4ZuFQQCuAwK/5XqjiwL2Jz6myEwBi0Il2sw+4LaXlgLGIjc+zSebmAKKRo1AC2Z0MJbFPG3SgOYCxiPWetx1WEBJAkSZK+WX8BAAD//yNx2Xt6EVY5AAAAAElFTkSuQmCC"
         alt="">
{{ end }}

</body>
</html>`

// Result represents the object evaluated in the html template
type Result struct {
	QRCode string
}

// QRServer is the main http handler
func QRServer(w http.ResponseWriter, r *http.Request) {
	// Determine User-Agent
	agent := r.Header.Get("User-Agent")
	plainTextResponse := isPlainTextResponse(agent)

	if plainTextResponse {
		plainTextHandler(w, r)
	} else {
		htmlHandler(w, r)
	}
}

func plainTextHandler(w http.ResponseWriter, r *http.Request) {
	pathParam := r.URL.Path[1:]

	if pathParam == "" {
		_, _ = fmt.Fprint(w, ErrMissingParam)
		return
	}

	pathParam = FixRequestPath(pathParam)
	res, err := getQRCodeFromCommand(pathParam, false)
	if err != nil {
		_, _ = fmt.Fprint(w, err)
		return
	}

	_, _ = fmt.Fprint(w, res)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	var result Result

	tpl, err := template.New("result").Parse(tmpl)
	if err != nil {
		_, _ = fmt.Fprint(w, err)
		return
	}

	pathParam := r.URL.Path[1:]

	if pathParam == "" || pathParam == "favicon.ico" {
		_ = tpl.Execute(w, result)
		return
	}

	pathParam = FixRequestPath(pathParam)
	res, err := getQRCodeFromCommand(pathParam, true)
	if err != nil {
		_, _ = fmt.Fprint(w, err)
		return
	}

	result.QRCode = res
	_ = tpl.Execute(w, result)
}

func getQRCodeFromCommand(s string, png bool) (string, error) {
	var res string
	if !png {
		code, err := qrcode.New(s, qrcode.Medium)
		if err != nil {
			return "", ErrEncodingFailed
		}
		res = code.ToSmallString(false)
	} else {
		png, err := qrcode.Encode(s, qrcode.Medium, 256)
		if err != nil {
			return "", ErrEncodingFailed
		}
		res = base64.StdEncoding.EncodeToString(png)
	}

	return res, nil
}

func isPlainTextResponse(ua string) bool {
	for _, el := range plainTextUserAgents {
		if strings.Contains(strings.ToLower(ua), el) {
			return true
		}
	}

	return false
}

// FixRequestPath rectifies issues caused by https://github.com/zeit/now/issues/3086
func FixRequestPath(path string) string {
	u, _ := url.ParseRequestURI(path)
	if u != nil && (u.Scheme == "" || u.Host == "") {
		components := strings.SplitN(path, "/", 2)
		res := components[0] + "//" + components[1]
		return res
	}

	return path
}
