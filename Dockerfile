FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/app /app/app
COPY internal/pkg/stopwords.txt /app/internal/pkg/stopwords.txt

ARG DB_NAME
ARG URI
ARG GRPC_PORT

ENV DB_NAME=$DB_NAME
ENV URI=$URI
ENV GRPC_PORT=$GRPC_PORT

EXPOSE 50051

CMD ["/app/app"]
