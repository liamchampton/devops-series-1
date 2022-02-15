FROM golang:1.17.6-alpine3.15 as build

WORKDIR /home

COPY go.mod ./

COPY go.sum ./

RUN go mod download

RUN mkdir views

COPY *.go ./

COPY views ./views

RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o website

RUN chmod +x website

FROM alpine:3.15

COPY --from=build home/views /views
COPY --from=build home/website /

EXPOSE 3000

ENTRYPOINT [ "/website" ]