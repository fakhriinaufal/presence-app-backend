# Stage 1
FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o main


# Stage 2
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8000
CMD [ "./main" ]