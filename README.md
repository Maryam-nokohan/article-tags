# Article Tags Service

![Go Version](https://img.shields.io/badge/Go-1.24.5-blue)
![Docker](https://img.shields.io/badge/Docker-supported-blue)

A high-performance gRPC service written in Go for managing and processing article tags, backed by MongoDB.

## 🚀 Features

- **gRPC API:** Fast and efficient communication.
- **Worker Pool Pattern:** Efficient concurrency management for processing large articles without exhausting system resources.
- **MongoDB Integration:** Persistent storage for article metadata.
- **Dockerized:** Fully containerized setup with Docker Compose.
- **CI/CD:** Automated builds and Docker Hub pushes via GitHub Actions.

---

## ⚙️ Environment Variables

To run this project, create a `.env` file in the root directory. This file is used by Docker Compose to inject build arguments and runtime variables.

| Variable    | Description                     | Default                 |
| :---------- | :------------------------------ | :---------------------- |
| `DB_NAME`   | Name of the MongoDB database    | `article`               |
| `URI`       | MongoDB connection string       | `mongodb://mongo:27017` |
| `GRPC_PORT` | Port the gRPC server listens on | `50051`                 |

---

## 🏃 Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.24.5 (if running locally without Docker)

### Installation & Setup

1. **Clone the repository:**
```bash
   git clone [https://github.com/Maryam-nokohan/article-tags.git](https://github.com/Maryam-nokohan/article-tags.git)
   cd article-tags
```

2. **Configure Environment:**
Create a .env file:
```bash
DB_NAME=article
URI=mongodb://mongo:27017
GRPC_PORT=50051
```

3. **Run with Docker Compose:**
```bash
docker-compose up --build
```
4. **Run client test :**

```bash
go run cmd/client/main.go
```

## 🧪 CI/CD Pipeline

This project uses **GitHub Actions**. On every push to `main`:

1.  **Code Checkout**: The repository is cloned into the runner.
2.  **Docker Build**: The image is built using **GitHub Secrets** as build arguments (`--build-arg`) to ensure sensitive data is not hardcoded.
3.  **Docker Push**: The final image is tagged and pushed to **Docker Hub** (`mary1385/article-tags`).



### Required GitHub Secrets

To use the CI/CD pipeline, add these to your repository under `Settings > Secrets and variables > Actions`:

| Secret | Purpose |
| :--- | :--- |
| `DOCKERHUB_USERNAME` | Your Docker Hub ID |
| `DOCKERHUB_TOKEN` | Personal Access Token for Docker Hub |
| `DB_NAME` | Name of the MongoDB database |
| `URI` | MongoDB connection string |
| `GRPC_PORT` | The port for the gRPC server |

---

## 📂 Project Structure

```
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   │   ├── grpc
│   │   │   └── handler.go
│   │   └── mongo
│   │       └── mongoDB.go
│   ├── application
│   │   ├── article_service.go
│   │   └── tag_extractor_service.go
│   ├── configs
│   │   ├── config.go
│   │   └── load.go
│   ├── domain
│   │   ├── article.go
│   │   ├── repository.go
│   │   └── tag_extractor.go
│   ├── utils
│   │   ├── normalizeText.go
│   │   └── stopwords.txt
│   └── workerpool
│       └── pool.go
├── LICENSE
├── proto
│   ├── article_grpc.pb.go
│   ├── article.pb.go
│   └── article.proto
└── READMME.md

```
