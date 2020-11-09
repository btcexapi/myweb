FROM golang:1.13.8 AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
ENV GOPROXY https://goproxy.io
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build  -o bin/myweb main.go

FROM alpine:3.7
COPY --from=builder /app/bin/myweb /app
WORKDIR /app
CMD "./myweb"
EXPOSE 8888