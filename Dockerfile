# syntax=docker/dockerfile:1

FROM golang:1.22.3-alpine3.20
LABEL version="1.0"
LABEL license="MIT"
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o /docker-ascii
EXPOSE 8080
CMD ["/docker-ascii"]
