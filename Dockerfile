FROM golang:1.22.0

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN echo 'DSN=' > .env

RUN go build -o main

CMD ["./main"]

EXPOSE 9090
