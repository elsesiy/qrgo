FROM golang:1.20 AS build-env

WORKDIR /qrgo
COPY main.go go.mod go.sum ./
COPY api ./api

RUN bash -c "go mod download &> /dev/null"

RUN CGO_ENABLED=0 GOOS=linux go build -o app -ldflags "-s -w" ./

FROM gcr.io/distroless/static:nonroot
COPY --from=build-env /qrgo/app .
CMD ["./app"]
