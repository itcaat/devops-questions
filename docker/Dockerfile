FROM golang:latest

WORKDIR /app

COPY . .

RUN go get ./...
RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]
