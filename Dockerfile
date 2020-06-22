FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build/cmd/app
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/cmd/app/main /app/
WORKDIR /app
EXPOSE 8080
CMD ["./main"]