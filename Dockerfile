FROM golang:1.16-alpine as builder

WORKDIR /app

COPY main.go ./ \
    go.mod ./

RUN go build -o todo

FROM alpine

COPY --from=builder /app /app

EXPOSE 8080/tcp

CMD ["/app/todo"]