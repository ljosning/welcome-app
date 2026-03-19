FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /app .

FROM alpine:3.21
COPY --from=builder /app /app
EXPOSE 3000
CMD ["/app"]
