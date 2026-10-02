# Notes API

A fast, robust, and clean **Notes management API** built with **Go (Golang)**. This project leverages **SQLC** for type-safe database queries and uses **PostgreSQL** as its primary data store.

## 🚀 Features
* **User Authentication**: Secure registration and login endpoints.
* **Task Management**: Full CRUD operations for user tasks (sub-routed under `/tasks`).
* **Profile Management**: View and update the authenticated user's profile.
* **Middleware Protection**: Route grouping with strict authentication checks (`h.AuthMiddleware`).
* **Type-Safe DB Queries**: SQLC integration to compile raw SQL into boilerplate-free Go code.
* **Database Migrations**: Automatic database schema management with PostgreSQL.
* **Dockerized Environment**: Quick setup with Docker Compose.


## 📁 Project Structure
* `cmd/app/` — Main entry point of the Go application.
* `config/` — Application configuration logic and environment setups.
* `internal/` — Business logic, custom API handlers, and SQLC-generated database packages.
* `migrations/` — Database schema files (SQL migrations).
* `pkg/` — Reusable helper packages and utilities.
* `sqlc.yaml` — SQLC generator configuration file.

## 🛠️ Prerequisites
Before running the project, make sure you have the following installed:
* [Go](https://go.dev) (1.22+ recommended)
* [Docker](https://docker.com) & [Docker Compose](https://docker.com)
* [SQLC](https://sqlc.dev) (Optional, only needed if you modify SQL queries)

## ⚡ Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com
cd notes-api
```

### 2. Run with Docker Compose (Recommended)
You can launch the API and the PostgreSQL database with a single command:
```bash
docker-compose up --build
```

### 3. Running Locally
If you want to run the Go application outside of Docker while connecting to a local database:
```bash
# Install dependencies
go mod download

# Run the app
go run cmd/app/main.go
```

## 🛠️ Database Development (SQLC)
If you update any SQL queries or schemas, regenerate the Go code using:
```bash
sqlc generate
```

## 🔌 API Endpoints

Routes within the protected group require a valid authentication token via `h.AuthMiddleware`.

### 🔑 Public Routes (Authentication)

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/auth/register` | Register a new user account |
| `POST` | `/auth/login` | Log in and receive an auth token |

### 🔒 Protected Routes (Auth Required)

#### User Profile

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/users/me` | Get the current authenticated user's profile |
| `PUT` | `/users/me` | Update the current user's profile |

#### Tasks (Notes)

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/tasks` | Create a new task |
| `GET` | `/tasks` | Get all tasks belonging to the user |
| `GET` | `/tasks/{id}` | Get a specific task by ID |
| `PUT` | `/tasks/{id}` | Update an existing task by ID |
| `DELETE` | `/tasks/{id}` | Delete a task by ID |

## 📝 License
This project is licensed under the [MIT License](LICENSE).
