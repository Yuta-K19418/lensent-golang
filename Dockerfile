# Build Container
FROM golang:1.15.2-alpine AS builder
RUN apk update
RUN apk add git
# Set Environment Variable
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/lensent
COPY . .
RUN go build -o ./app/main main.go

# Runtime Container
FROM alpine
COPY --from=builder /go/src/lensent/app /app
# Set Environment Variable
EXPOSE 8080
ENTRYPOINT ["/app/main"]