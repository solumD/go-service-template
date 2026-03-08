FROM golang:alpine AS builder

LABEL maintainer="so1umD"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/app . 

CMD ["./app"]