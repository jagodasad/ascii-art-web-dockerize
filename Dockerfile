
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go build -o main server.go

FROM alpine3.13
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8070
CMD [ "/app/main" ]