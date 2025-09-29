FROM golang:1.25 AS build-env

WORKDIR /qrgo
COPY main.go go.mod go.sum ./
COPY api ./api

RUN bash -c "go mod download &> /dev/null" && \
  CGO_ENABLED=0 GOOS=linux go build -o app -ldflags "-s -w" ./

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=build-env /qrgo/app .
CMD ["./app"]
