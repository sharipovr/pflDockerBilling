# 📦 pflDockerBilling (Billing Project)

This project implements a subscription and payment management system integrated with Stripe. This is a demo/portfolio project utilizing Stripe in test mode.

---

This project is designed to demonstrate my skills as a candidate for the position of **Senior Backend Engineer**.

---

## 📌 Project Goals

- **Demonstrate knowledge and practical experience** in developing and maintaining billing systems using modern technologies.
- **Create a system** as close as possible to real-world scenarios, including integration with third-party services (Stripe), cloud providers, and CI/CD systems.
- **Showcase my ability to develop**, test, document, and deploy applications in a production environment.

---

## 🚀 Quick Start (Running the project locally)

### 🛠 Requirements
- Docker and Docker Compose
- Go 1.22+
- Node.js 18+

### ▶️ Launching the project (Docker Compose)

Start with containers:
```bash
docker compose -f docker-compose.yml -f docker-compose.app.yml up -d --build
```

**Service Access:**

| Service   | URL                                            |
|-----------|------------------------------------------------|
| Frontend  | [http://localhost:5173](http://localhost:5173) |
| Backend   | [http://localhost:8080](http://localhost:8080) |
| Postgres  | postgres://localhost:5432                      |

### ⚙️ Running without Docker (for development)

Run Backend:

```bash
cd backend
go run cmd/main.go
```

Run Frontend:

```bash
cd frontend
npm install
npm run dev
```

---

## 📖 Deployment Instructions (CI/CD on Linode)

### Step 1: Server Setup (one-time setup)

- Install Docker and Docker Compose on your Linode server
- Clone the repository and create a `.env` file with your configuration settings

### Step 2: CI/CD Setup (GitHub Actions)

CI/CD is configured using GitHub Actions and Docker Hub.

On each push to the `main` branch:

- Docker images (`backend` and `frontend`) are built
- Images are uploaded to Docker Hub
- The server automatically fetches changes and restarts containers

> ⚠️ Ensure you've created secrets `DOCKERHUB_USERNAME`, `DOCKERHUB_TOKEN`, `SSH_PRIVATE_KEY`, and `SERVER_HOST` in GitHub Secrets!

---

## 🗂 Architectural Decisions

- **Backend:** Go, Gin, GORM, PostgreSQL
- **Frontend:** React, Vite, TypeScript, Axios, Tailwind CSS
- **Payments:** Stripe API (Subscriptions and Webhooks)
- **CI/CD:** Docker Compose, GitHub Actions, Docker Hub
- **Testing:** Go (sqlite3), React (Vitest, React Testing Library)

---

## 📚 API Documentation

API documentation is implemented using Swagger (OpenAPI).

- Access: `http://localhost:8080/swagger/index.html`

---

## 📋 License

MIT License
