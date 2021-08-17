# Build Container
FROM golang:1.15.2-alpine AS builder
RUN apk update
RUN apk add git

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/lensent
COPY . .
RUN go build -o app main.go
RUN touch ./local.env
RUN echo "DB_HOST=${DB_HOST}" >> ./local.env
RUN echo "DB_USER=postgres" >> ./local.env
RUN echo "DB_PASSWORD=postgres" >> ./local.env
RUN echo "DB=lensent" >> ./local.env
RUN echo "DB_PORT=5432" >> ./local.env
COPY ./local.env ./app/.

# Runtime Container
FROM alpine
COPY --from=builder /go/src/lensent/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]