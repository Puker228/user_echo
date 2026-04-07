FROM golang:1.25.5-alpine AS builder

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o app ./cmd/app

FROM alpine:3.22
WORKDIR /usr/src/app

ENV GIN_MODE=release

COPY --from=builder /usr/src/app/app .

CMD ["./app"]
