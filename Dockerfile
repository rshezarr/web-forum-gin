FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o web_forum ./cmd/app/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app .
EXPOSE 9090
CMD ["/app/web_forum"]