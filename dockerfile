# Base Go image
FROM golang:1.21.5-alpine3.19 as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o enrollmentApp ./cmd/api

RUN chmod +x /app/enrollmentApp

# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app
COPY --from=builder /app/enrollmentApp /app

CMD [ "/app/enrollmentApp" ]
