FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o web_forum ./cmd/app/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/web_forum .
COPY --from=builder /app/database ./database
COPY --from=builder /app/configs ./configs
EXPOSE 9090
CMD ["/app/web_forum"]