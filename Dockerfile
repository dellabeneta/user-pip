FROM golang:alpine3.21 AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates git

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o pip .

FROM alpine:3.21

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/pip .

EXPOSE 8080

CMD ["./pip"]
