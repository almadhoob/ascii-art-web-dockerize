# syntax=docker/dockerfile:1

FROM golang:1.22.3-alpine3.20
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-ascii
EXPOSE 8080
CMD ["/docker-ascii"]

LABEL version="1.0"
LABEL license="MIT"