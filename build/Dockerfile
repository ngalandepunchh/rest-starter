FROM golang:1.12 AS builder

WORKDIR /app
ADD . /app

RUN go build -v -o app

FROM gcr.io/distroless/base
COPY --from=builder /app/app /
CMD ["/app"]