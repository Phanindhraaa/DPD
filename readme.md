# 🐳 Multi-Service Dockerized Application with NGINX Reverse Proxy

This is a DevOps project demonstrating how to containerize and orchestrate a two-service application (Go + Flask) behind an NGINX reverse proxy using Docker Compose.

---

## 📦 Services Overview

| Service       | Description                         | Port   |
|---------------|-------------------------------------|--------|
| `service1`    | Go-based service with `/ping` & `/hello` endpoints | `8001` |
| `service2`    | Python Flask service with `/ping`, `/hello`, `/`  | `8002` |
| `nginx`       | Acts as a reverse proxy             | `8080` |

---

## 🧱 Tech Stack

- Docker
- Docker Compose
- NGINX
- Go (service1)
- Python Flask (service2)
- Shell scripting (for healthcheck)

---

## 🚀 Getting Started

### 1. Clone the Repo and run the docker-compose

```bash
git clone https://github.com/Phanindhraaa/DPD.git
cd DPD
docker compose up --build

## Visit in your browser or via curl:

http://localhost:8080/service1/ping

http://localhost:8080/service1/hello

http://localhost:8080/service2/ping

http://localhost:8080/service2/hello

http://localhost:8080/service2/

## Health Check 

./healthcheck.sh

OUTPUT:
🔍 Checking Service 1...
🔍 Checking Service 2...
✅ All services are healthy!

## Project Structure
DPD/
├── service_1/
│   ├── main.go
│   └── Dockerfile
├── service_2/
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
├── nginx/
│   ├── nginx.conf
│   └── Dockerfile
├── docker-compose.yml
├── healthcheck.sh
└── README.md




