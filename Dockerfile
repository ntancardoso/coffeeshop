FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . . 

# RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o coffeeshop_app .

FROM scratch

WORKDIR /app
COPY --from=builder /app/coffeeshop_app .
COPY --from=builder /app/config.json ./config.json
COPY --from=builder /app/templates ./templates

CMD ["./coffeeshop_app"]