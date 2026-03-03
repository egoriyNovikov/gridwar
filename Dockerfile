FROM golang:1.26-alpine

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "-mod=mod", "app/cmd/server/main.go"]