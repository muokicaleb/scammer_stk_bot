FROM golang:latest AS builder

WORKDIR /

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app .


FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ=Africa/Nairobi

WORKDIR /

COPY --from=builder /app /.env ./

EXPOSE 8080

RUN chmod +x ./app

ENTRYPOINT ["./app"]
