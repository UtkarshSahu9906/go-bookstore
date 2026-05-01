# 📚 Go Bookstore API

A simple REST API built with **Go**, **Gorilla Mux**, and **GORM (MySQL)** that lets you manage a bookstore — create, read, update, and delete books.

---

## 🛠️ Tech Stack

| Tool | Purpose |
|---|---|
| [Go](https://go.dev/) | Programming language |
| [Gorilla Mux](https://github.com/gorilla/mux) | HTTP router |
| [GORM](https://gorm.io/) | ORM for database |
| [MySQL](https://www.mysql.com/) | Database |

---

## 📁 Project Structure

```
go-bookstore/
├── cmd/
│   └── main.go                  # Entry point — starts the server
├── pkg/
│   ├── config/
│   │   └── app.go               # Database connection
│   ├── controllers/
│   │   └── book-controller.go   # Request handlers
│   ├── models/
│   │   └── book.go              # Book model + database queries
│   ├── routes/
│   │   └── bookstore-routes.go  # API route definitions
│   └── utils/
│       └── utils.go             # Helper functions
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Prerequisites

Make sure you have these installed:

- [Go 1.22+](https://go.dev/dl/)
- [MySQL 8.0+](https://dev.mysql.com/downloads/installer/)

---

## 🚀 Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/your-username/go-bookstore.git
cd go-bookstore
```

### 2. Setup MySQL Database
Open MySQL and run:
```sql
CREATE DATABASE simplerest;
CREATE USER 'utkarsh'@'localhost' IDENTIFIED BY 'utkarsh';
GRANT ALL PRIVILEGES ON simplerest.* TO 'utkarsh'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```

### 3. Install dependencies
```bash
go mod tidy
```

### 4. Run the server
```bash
go run cmd/main.go
```

Server will start at **http://localhost:9010** 🎉

> If the terminal looks frozen/blank — that's normal! It means the server is running.

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/book/` | Get all books |
| `POST` | `/book/` | Add a new book |
| `GET` | `/book/{id}` | Get a book by ID |
| `PUT` | `/book/{id}` | Update a book |
| `DELETE` | `/book/{id}` | Delete a book |

---

## 🧪 Test the API

### ➕ Add a book
```bash
curl -X POST http://localhost:9010/book/ \
  -H "Content-Type: application/json" \
  -d '{"name":"Harry Potter","author":"Rowling","publication":"Bloomsbury"}'
```

### 📚 Get all books
```bash
curl http://localhost:9010/book/
```

### 🔍 Get a book by ID
```bash
curl http://localhost:9010/book/1
```

### ✏️ Update a book
```bash
curl -X PUT http://localhost:9010/book/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Harry Potter 2","author":"Rowling","publication":"Bloomsbury"}'
```

### 🗑️ Delete a book
```bash
curl -X DELETE http://localhost:9010/book/1
```

---

## 📦 Sample Response

```json
{
  "ID": 1,
  "CreatedAt": "2026-05-01T18:08:15+05:30",
  "UpdatedAt": "2026-05-01T18:08:15+05:30",
  "DeletedAt": null,
  "name": "Harry Potter",
  "author": "Rowling",
  "publication": "Bloomsbury"
}
```

---

## 🪟 Windows Users

Use `\"` instead of `'` in curl commands:
```cmd
curl -X POST http://localhost:9010/book/ -H "Content-Type: application/json" -d "{\"name\":\"Harry Potter\",\"author\":\"Rowling\",\"publication\":\"Bloomsbury\"}"
```

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).