FROM golang:alpine as BUILDER

WORKDIR /go/src/todo/
COPY . .

RUN go build -o todo .

FROM alpine

WORKDIR /todo

COPY --from=BUILDER /go/src/todo/ /todo/

CMD ["./cmd/todo"]

