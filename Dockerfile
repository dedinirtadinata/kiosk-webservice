# syntax=docker/dockerfile:1
FROM golang:1.19 AS build-stage

#set detination for copy
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY delivery delivery
COPY display display
COPY config.json ./

RUN CGO_ENABLED=1 GOOS=linux go build -o /docker-kiosk

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-kiosk /docker-kiosk

EXPOSE 8099
CMD ["/docker-kiosk"]
