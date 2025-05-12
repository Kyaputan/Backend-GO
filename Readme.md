# 🚀 Go Backend Showcase: Fiber | Gin

> A lightning-fast journey through two of the hottest Go web frameworks.

---

## 📘 Introduction

Welcome to the **Go Backend Showcase**, where we explore **Fiber** and **Gin**, two of the most popular and performant web frameworks in the Go ecosystem:

- ⚡ **Fiber** – Inspired by Express.js, built on top of Fasthttp for extreme speed.
- 🍃 **Gin** – A minimalist, high-performance HTTP web framework with a martini-like API.

Whether you’re building microservices, REST APIs, or full-featured web apps, this repo will guide you through core concepts and sample implementations for each framework.

---

## 🏗️ Project Structure

```bash
go-backend-showcase/
│
├── fiber_app/          # Fiber-based backend project
│   └── ...             # Handlers, routes, middleware, etc.
│
├── gin_app/            # Gin-based backend project
│   └── ...             # Controllers, routes, middleware, etc.
│
├── go.mod              # Module definition
├── go.sum              # Dependency lock file
└── README.md           # You’re reading it now!
```

---

## 🧰 Tech Stack

| Feature                | Fiber                                        | Gin                                              |
|------------------------|----------------------------------------------|--------------------------------------------------|
| Routing                | ✅ Express-like, fast                        | ✅ Martini-like, concise                          |
| Middleware             | ✅ Built-in, chaining                        | ✅ Built-in, easy to plug-in                      |
| HTTP Engine            | ⚡ Fasthttp (blazing fast)                   | 🚀 net/http (battle-tested)                      |
| Performance            | 🔥 Very high (benchmarks favor Fiber)        | ⚡ Very high (lightweight)                       |
| Templating             | 🔄 Using any Go template engine              | 🔄 Using any Go template engine                  |
| ORM Integration        | 🔄 GORM, SQLX, ent-go, etc.                  | 🔄 GORM, SQLX, ent-go, etc.                      |
| Context & Params       | ✅ Simple context, path params               | ✅ Detailed context, path params                  |
| Swagger / OpenAPI      | 🔄 fiber-swagger plugin                      | 🔄 swaggo/gin-swagger                            |
| Hot Reloading          | 🔄 air, fresh, or other third-party tools    | 🔄 air, fresh, or other third-party tools        |
| Learning Curve         | 🎈 Gentle for Express.js users               | 🎈 Gentle for Go developers                       |
| Community & Ecosystem  | 📈 Growing fast                              | 📊 Mature, large user base                       |

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-backend-showcase.git
cd go-backend-showcase
```

### 2. Install Go and Set Up Environment

- Install Go (v1.20+) from [golang.org](https://golang.org/dl/)
- Ensure `$GOPATH/bin` is in your `PATH`

### 3. Download Dependencies

```bash
go mod download
```

### 4. Run Each Project

#### ▶️ Fiber

```bash
cd fiber_app
go run main.go
# Visit http://localhost:3000
```

#### ▶️ Gin

```bash
cd gin_app
go run main.go
# Visit http://localhost:8080
```

---

## 🧪 Features Demonstrated

- 🚀 CRUD APIs with JSON responses  
- 🔄 Middleware for logging, CORS, and error handling  
- 🗄️ Database integration using GORM (SQLite/MySQL)  
- 📜 Request validation and binding  
- 📦 Modular router grouping  
- ⚙️ Swagger/OpenAPI auto-generated docs  
- 🔧 Hot-reload development workflow  

---

## 📚 Recommended Learning Resources

- [Fiber Official Docs](https://docs.gofiber.io/)  
- [Gin Official Docs](https://gin-gonic.com/docs/)  
- [Go by Example](https://gobyexample.com/)  
- [Building RESTful APIs with Go](https://www.youtube.com/watch?v=SonwZ6MF5BE)  

---

## ❤️ Contributing

Contributions are welcome! Feel free to:
- 🐛 Report bugs  
- ✨ Suggest new features  
- 📚 Improve documentation  
- 🛠️ Submit pull requests  

Please fork the repo, create a branch, and open a PR—let’s build together!

---

## 📄 License

MIT License © 2025 Rachata Singkhet (Caption)

---

> “Go beyond frameworks—master the language.” – Caption
