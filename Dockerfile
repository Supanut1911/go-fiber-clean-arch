#first stage - builder
FROM golang:1.21-alpine3.18 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN  go build -o main main.go


#second stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 3000
CMD ["/app/main"]