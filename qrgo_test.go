package qrgo

import (
	"net/http/httptest"
	"testing"
)

func TestGETQRCodePlain(t *testing.T) {
	testCases := map[string]struct {
		pathParam    string
		expectedBody string
	}{
		"valid path param": {"abc", `█████████████████████████████
█████████████████████████████
████ ▄▄▄▄▄ ██▀ ▄▄█ ▄▄▄▄▄ ████
████ █   █ █ ▄▄▀▀█ █   █ ████
████ █▄▄▄█ █ █ ▀▄█ █▄▄▄█ ████
████▄▄▄▄▄▄▄█ ▀▄█▄█▄▄▄▄▄▄▄████
████ ▀ ▄▄ ▄▀▀▄▄ ▀▀ ▄▄▄ █▀████
█████▄██▀█▄▄▀▄▀▀ ▀ ▄█▀▀ █████
█████▄██▄▄▄█ █▄█ ██ ██▀█▀████
████ ▄▄▄▄▄ █▀▀▄▄█▄█▀ ▄▀ ▀████
████ █   █ █  █ ▀ ▀█▄▀▄██████
████ █▄▄▄█ █▄██▀ ▀ ▄█▀███████
████▄▄▄▄▄▄▄█▄▄▄█▄██▄██▄▄█████
█████████████████████████████
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
`},
		"empty param": {"", ErrMissingParam.Error()},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/"+v.pathParam, nil)
			req.Header.Add("User-Agent", "curl")
			res := httptest.NewRecorder()

			QRServer(res, req)

			got := res.Body.String()
			want := v.expectedBody

			assertResult(t, got, want)
		})
	}
}

func TestGETQRCodeHTML(t *testing.T) {
	testCases := map[string]struct {
		pathParam    string
		expectedBody string
	}{
		"valid path param": {"abc", `<!DOCTYPE html>
<html>
<head>
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-104096180-2"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-104096180-2');
</script>
<title>qrgo</title>
</head>
<body style='background-color: black;color:white;'>
<a href="https://github.com/elsesiy/qrgo" class="github-corner" aria-label="View source on GitHub"><svg width="80" height="80" viewBox="0 0 250 250" style="fill:#fff; color:#151513; position: absolute; top: 0; border: 0; right: 0;" aria-hidden="true"><path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path><path d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2" fill="currentColor" style="transform-origin: 130px 106px;" class="octo-arm"></path><path d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z" fill="currentColor" class="octo-body"></path></svg></a><style>.github-corner:hover .octo-arm{animation:octocat-wave 560ms ease-in-out}@keyframes octocat-wave{0%,100%{transform:rotate(0)}20%,60%{transform:rotate(-25deg)}40%,80%{transform:rotate(10deg)}}@media (max-width:500px){.github-corner:hover .octo-arm{animation:none}.github-corner .octo-arm{animation:octocat-wave 560ms ease-in-out}}</style>


    <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN&#43;AAABD0lEQVR42uyYPY7DIBSEn7UFpY/AUTgaORpH4QiULqLM6vEnm8Rsuatlvoqgr3pixhAhhJC/zo5CyMvXVn8eFEahLH0Qq7sbzrtLCQaA&#43;OBQBQ8kChMhitETReFHYT&#43;&#43;KEyEmk0Xdzxn4V1daGXeJnnb9osLHYc0vyf8b6FET7/&#43;FvlEKYnCu6DbgoeDlrnWlYhN1&#43;hRUExeewSbJEfvIbKfR0mhkxtcy7xM0gEU3gVz5EDWMtfTN7QYhaHto0gVrtlcQ7g8aZ/bS1MYz9GjMN4na96AYD&#43;&#43;/ZcX&#43;ps3mTJJrXWhMBNqmUehMBFyiZXrAYVPQvtfTuRmkmsI/UkLtK&#43;/RaIwCoQQ8jt8BwAA///ImVmi9gMnrgAAAABJRU5ErkJggg==" alt="">


</body>
</html>`},
		"empty param": {"", `<!DOCTYPE html>
<html>
<head>
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-104096180-2"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-104096180-2');
</script>
<title>qrgo</title>
</head>
<body style='background-color: black;color:white;'>
<a href="https://github.com/elsesiy/qrgo" class="github-corner" aria-label="View source on GitHub"><svg width="80" height="80" viewBox="0 0 250 250" style="fill:#fff; color:#151513; position: absolute; top: 0; border: 0; right: 0;" aria-hidden="true"><path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path><path d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2" fill="currentColor" style="transform-origin: 130px 106px;" class="octo-arm"></path><path d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z" fill="currentColor" class="octo-body"></path></svg></a><style>.github-corner:hover .octo-arm{animation:octocat-wave 560ms ease-in-out}@keyframes octocat-wave{0%,100%{transform:rotate(0)}20%,60%{transform:rotate(-25deg)}40%,80%{transform:rotate(10deg)}}@media (max-width:500px){.github-corner:hover .octo-arm{animation:none}.github-corner .octo-arm{animation:octocat-wave 560ms ease-in-out}}</style>


    <pre>
Usage is simple, just append the content to be encoded to the Path, i.e.

$ curl -L qrgo.elsesiy.com/abc
    </pre>

    <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAAAz0lEQVR42uyX262CMQyDu0H239Ib9Ogk6Q0KvEJsg6q/0fdkNbcmSZIk/Yisu+BfWDc6AHmaQ1gxNsAN8nPdiAG3SsD80QKJ4T9z3mRWdSBL57TpVaUtDiwQnxpsaSDNgffVZs82sQAej6zpo6mAD4hq4b0kubN+0ADJnA+HELCIWQ4ZuFQQCuAwK/5XqjiwL2Jz6myEwBi0Il2sw+4LaXlgLGIjc+zSebmAKKRo1AC2Z0MJbFPG3SgOYCxiPWetx1WEBJAkSZK+WX8BAAD//yNx2Xt6EVY5AAAAAElFTkSuQmCC"
         alt="">


</body>
</html>`},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/"+v.pathParam, nil)
			res := httptest.NewRecorder()

			QRServer(res, req)

			got := res.Body.String()
			want := v.expectedBody

			assertResult(t, got, want)
		})
	}
}

func assertResult(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
