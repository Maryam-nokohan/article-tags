# Article Tags Service

![Go Version](https://img.shields.io/badge/Go-1.24.5-blue)
![Docker](https://img.shields.io/badge/Docker-supported-blue)
![gRPC](https://img.shields.io/badge/gRPC-enabled-green)
![MongoDB](https://img.shields.io/badge/MongoDB-supported-brightgreen)
![NATS](https://img.shields.io/badge/NATS-event--driven-orange)
![S3](https://img.shields.io/badge/S3-storage-blue)

A high-performance event-driven article processing system written in Go. The service uses gRPC for article ingestion, NATS for asynchronous messaging, MongoDB for persistence, and S3-compatible storage for article content.

---

## 🚀 Features

* **gRPC API** for article publishing.
* **Event-Driven Architecture** using NATS.
* **MongoDB Integration** for article metadata persistence.
* **S3 Object Storage** for article content storage.
* **Asynchronous Processing Pipeline**.
* **Worker Pool Pattern** for efficient tag extraction.
* **Dockerized Deployment** with Docker Compose.
* **Clean Architecture** with Ports and Adapters pattern.

---

## 🏗️ Architecture

```text
                +----------------------+
                |  gRPC Publish Client |
                +----------+-----------+
                           |
                           v
                +----------------------+
                | PublishArticleService|
                +----------+-----------+
                           |
                           |
                           v
                       MongoDB

                           |
                           | Publish Event
                           v
                         NATS
                  article.created
                           |
                           v
                +----------------------+
                | Article Processing   |
                |      Service         |
                +----------+-----------+
                           |
                +----------+-----------+
                |                      |
                v                      v

         Extract Tags          Upload Content
                                   to S3

                +----------+-----------+
                           |
                           v
                       MongoDB
```

---

## ⚙️ Environment Variables

Create a `.env` file in the project root:

```env
# MongoDB
DB_NAME="article"
URI=mongodb://mongo:27017

# gRPC
GRPC_PORT=50051

# NATS
NATS_URL=nats://nats:4222

# S3
ENDPOINT=https://aws/articles-storage
REGION=uk
ACCESSKEY=YOUR_ACCESS_KEY
SECRETEKEY=YOUR_SECRET_KEY
BUCKET=articles-storage
```

### Variable Description

| Variable     | Description                           |
| ------------ | ------------------------------------- |
| `DB_NAME`    | MongoDB database name                 |
| `URI`        | MongoDB connection URI                |
| `GRPC_PORT`  | gRPC server port                      |
| `NATS_URL`   | NATS broker connection URL            |
| `ENDPOINT`   | S3-compatible object storage endpoint |
| `REGION`     | Storage region                        |
| `ACCESSKEY`  | Object storage access key             |
| `SECRETEKEY` | Object storage secret key             |
| `BUCKET`     | Bucket used to store article content  |

> Never commit real credentials to source control. Store them in a local `.env` file or CI/CD secrets.

---

## 🏃 Getting Started

### Prerequisites

* Docker
* Docker Compose
* Go 1.24.5+

---

### Clone Repository

```bash
git clone https://github.com/Maryam-nokohan/article-tags.git
cd article-tags
```

---

### Configure Environment

Create a `.env` file:

```env
DB_NAME=article
URI=mongodb://mongo:27017

GRPC_PORT=50051

BROKER_URL=nats://nats:4222
ARTICLE_CREATED_SUBJECT=article.created

S3_ENDPOINT=minio:9000
S3_ACCESS_KEY=minioadmin
S3_SECRET_KEY=minioadmin
S3_BUCKET=articles
S3_USE_SSL=false
```

---

### Run with Docker Compose

```bash
docker compose up --build
```

This starts:

* MongoDB
* NATS
* S3
* gRPC Server

---

## 🧪 Running Clients

### Publish an Article

```bash
go run cmd/clients/client_article_creation/main.go
```

This client sends an article to the gRPC server.

The server:

1. Stores article metadata.
2. Publishes an `article.created` event to NATS.

---

### Process Articles

```bash
go run cmd/clients/client_article_processing/main.go
```

The processing client subscribes to article creation events and performs:

* Tag extraction
* Content processing
* Upload to S3
* Metadata updates

---

## 📂 Project Structure

```text
├── cmd
│   ├── clients
│   │   ├── client_article_creation
│   │   │   └── main.go
│   │   └── client_article_processing
│   │       └── main.go
│   └── server
│       └── main.go
│
├── internal
│   ├── adapters
│   │   ├── grpc
│   │   ├── mongo
│   │   ├── nats
│   │   └── s3
│   │
│   ├── application
│   │   ├── publish_article_service.go
│   │   ├── article_processing_service.go
│   │   └── tag_extractor_service.go
│   │
│   ├── configs
│   │   ├── config.go
│   │   └── load.go
│   │
│   ├── domain
│   │   ├── article.go
│   │   └── event.go
│   │
│   ├── pkg
│   │   ├── normalizeText.go
│   │   ├── pool.go
│   │   └── stopwords.txt
│   │
│   └── ports
│       ├── repository.go
│       ├── message_broker.go
│       ├── object_storage.go
│       └── tag_extractor.go
│
├── proto
│   ├── article.proto
│   ├── article.pb.go
│   └── article_grpc.pb.go
│
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

## 🔄 Event Flow

### Article Creation

1. Client sends article via gRPC.
2. Server validates article.
3. Article metadata is stored in MongoDB.
4. `article.created` event is published to NATS.

### Article Processing

1. Processing service subscribes to `article.created`.
2. Event is consumed from NATS.
3. Article content is processed.
4. Tags are extracted using the worker pool.
5. Content is uploaded to S3.
6. Article metadata is updated in MongoDB.

---

## 🛠️ Technologies

* Go 1.24.5
* gRPC
* MongoDB
* NATS
* S3 / MinIO
* Docker
* Docker Compose

---

## 🧪 Testing

Run the publishing client:

```bash
go run cmd/clients/client_article_creation/main.go
```

Run the processing client:

```bash
go run cmd/clients/client_article_processing/main.go
```

---

## 📜 License

This project is licensed under the MIT License.
