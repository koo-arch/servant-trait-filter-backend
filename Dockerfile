FROM golang:1.24.3-bookworm

WORKDIR /backend

RUN go install github.com/air-verse/air@v1.61.7

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["air"]