# qrgo

[![Go Report Card](https://goreportcard.com/badge/github.com/elsesiy/qrgo)](https://goreportcard.com/report/github.com/elsesiy/qrgo)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/bf3c077335e046c5a7371c4900011618)](https://www.codacy.com/app/elsesiy/qrgo?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=elsesiy/qrgo&amp;utm_campaign=Badge_Grade)
[![Twitter](https://img.shields.io/badge/twitter-@elsesiy-blue.svg)](http://twitter.com/elsesiy)

:zap: Fast & simple service to generate QR codes from your CLI written in Go.

## Usage

The web app generates png output whereas UTF8 strings are used to display it for your CLI.

### Web

`curl -L qrgo.elsesiy.com/test`

or if you want to pipe the output of another command:

`echo "test" | curl -L qrgo.elsesiy.com/-`

### Local (via Docker)

Run `docker build -t elsesiy/qrgo . && docker run --rm -d -p 3000:3000 elsesiy/qrgo` and browse [here](http://localhost:3000).

### Local (from Source)

1. Install dependency

    `go mod download`

2. Build and run

    `go build && ./qrgo`

#### Links

- [go-qrcode](https://github.com/skip2/go-qrcode)
- [qrenco.de](https://github.com/chubin/qrenco.de)