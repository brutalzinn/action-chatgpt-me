FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY . .

RUN apk update && apk add --no-cache git
RUN go build -o /bin/app main.go

ENTRYPOINT ["app"]