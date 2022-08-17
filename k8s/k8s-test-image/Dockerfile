FROM golang:1.18.0-alpine3.15 as builder
WORKDIR /build
COPY main.go .

RUN go build -o main main.go
CMD ["/build/main"]

FROM scratch
WORKDIR /app
COPY --from=builder /build/main .


