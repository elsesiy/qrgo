FROM golang:1.13 AS build-env

WORKDIR /qrgo
COPY qrgo.go go.mod go.sum ./
COPY main ./main

RUN bash -c "go mod download &> /dev/null"

RUN CGO_ENABLED=0 GOOS=linux go build -o app -ldflags "-s -w" ./main

FROM gcr.io/distroless/static:latest
COPY --from=build-env /qrgo/app .
CMD ["./app"]