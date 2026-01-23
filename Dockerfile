FROM golang:1.25-alpine AS build

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o task ./cmd

FROM alpine:3.23.2

RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /usr/src/app

COPY --from=build /usr/src/app/task .

EXPOSE 8080
CMD ["./task"]
