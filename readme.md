# Student Management REST API (Go + SQLite)

A simple **RESTful API built with Go** that performs **CRUD operations for student records**.
This project demonstrates how to build a **production-style backend in Go** using the standard library, SQLite database, YAML configuration, and clean project structure.

The API allows you to **create, read, update, and delete student data** while following good backend practices like structured JSON responses and graceful server shutdown.

---

## 🚀 Features

* RESTful API built with **Go (Golang)**
* **CRUD operations** for student records
* **SQLite database integration**
* **YAML configuration file** for environment settings
* **Structured JSON responses**
* Clean **project architecture**
* **Graceful server shutdown**
* Uses Go’s **net/http standard library**

---

## 📁 Project Structure

```
student-api/
│
├── cmd/
│   └── student-api/
│       └── main.go          # Application entry point
│
├── config/
│   └── local.yaml           # Application configuration
│
├── internal/
│   ├── config/              # Config loader
│   ├── handlers/            # API request handlers
│   ├── storage/             # Database logic
│   │   └── sqlite/
│   └── types/               # Data models
│
├── pkg/
│   └── response/            # JSON response utility
│
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Configuration

The project uses a **YAML configuration file** to store application settings like the database path and server address.

Example `config/local.yaml`:

```yaml
env: "local"

storage_path: "./storage/storage.db"

http_server:
  address: "localhost:8080"
  timeout: 4s
  idle_timeout: 60s
```

---

## 🗄️ Database

The project uses **SQLite** as a lightweight file-based database.

Example student table structure:

```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT,
    age INTEGER
);
```

---

## 📡 API Endpoints

### Create Student

```
POST /students
```

Request Body

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 21
}
```

---

### Get Student By ID

```
GET /students/{id}
```

Response

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "age": 21
}
```

---

### Update Student

```
PUT /students/{id}
```

---

### Delete Student

```
DELETE /students/{id}
```

---

## 🧠 JSON Response Utility

The project includes a reusable **JSON response helper** that:

* Sets the correct `Content-Type: application/json`
* Encodes response data
* Handles encoding errors

Example:

```go
func JSON(w http.ResponseWriter, status int, data interface{}) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(data)
}
```

---

## 🛑 Graceful Server Shutdown

The server listens for system interrupts such as:

```
Ctrl + C
```

When triggered, the server:

* Stops accepting new requests
* Completes ongoing requests
* Closes the application safely

This ensures **no data loss or broken connections**.

---

## ▶️ Running the Project

### 1️⃣ Clone the repository

```
git clone https://github.com/your-username/student-api.git
cd student-api
```

### 2️⃣ Install dependencies

```
go mod tidy
```

### 3️⃣ Run the server

```
go run cmd/student-api/main.go -config config/local.yaml
```

Server will start at:

```
http://localhost:8080
```

---

## 🧪 Example Curl Request

Create a student:

```
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '{"name":"John","email":"john@example.com","age":22}'
```

---

## 📚 What I Learned

While building this project I learned:

* Building **REST APIs in Go**
* Using **net/http** for backend services
* Integrating **SQLite with Go**
* Writing **structured handlers**
* Managing **JSON serialization with struct tags**
* Creating reusable **utility packages**
* Implementing **graceful server shutdown**
* Organizing a **clean Go project architecture**

---

## 🛠 Tech Stack

* **Go (Golang)**
* **SQLite**
* **net/http**
* **YAML configuration**

---

## 👨‍💻 Author

Built as a **learning project to understand backend development in Go** and how to structure a real-world REST API.
