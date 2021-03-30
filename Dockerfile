FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 ARCH=amd64 GOOS=linux go build .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/meme-ddoser meme-ddoser

CMD /app/meme-ddoser