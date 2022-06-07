FROM golang:1.16-alpine as builder

WORKDIR /app

COPY main.go ./

COPY go.mod ./

RUN go build -o todo

FROM alpine

WORKDIR /app

COPY --from=builder /app /app

EXPOSE 8080 8080

ENTRYPOINT ["/app/todo"]
