FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/server .

COPY internal/utils/stopwords.txt ./internal/utils/stopwords.txt

ENV PORT=50051
ENV DB_NAME=articles
ENV URI=mongodb://mongo:27017

EXPOSE 50051

CMD ["./server"]