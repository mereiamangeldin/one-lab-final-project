FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

#RUN go build -o myapp ./cmd

EXPOSE 8080

CMD ["go", "run", "./cmd/main.go"]