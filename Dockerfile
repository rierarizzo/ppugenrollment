# Base Go image
FROM golang:1.21.4-alpine3.18 as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o enrollmentApp .

RUN chmod +x /app/enrollmentApp

# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app
COPY --from=builder /app/enrollmentApp /app

CMD [ "/app/enrollmentApp" ]
