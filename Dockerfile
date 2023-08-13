# syntax=docker/dockerfile:1
FROM golang:1.19

#set detination for copy
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-kiosk

EXPOSE 8099
CMD ["/docker-kiosk"]
