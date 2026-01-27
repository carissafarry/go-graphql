# Go GraphQL Server (gqlgen)

This repository contains a **GraphQL server implemented in Go** using **gqlgen**, following a **schema-first approach** and a **clean, modular project structure**.  
The implementation is based on and inspired by the tutorial:

> **Implementation GraphQL Server With Golang**  
> https://medium.com/@wahyubagus1910/implementation-graphql-server-with-golang-fb8f8303b4bc

---

## ✨ Features

- GraphQL server built with **Golang**
- Schema-first GraphQL using **gqlgen**
- Clear separation of concerns:
  - Resolver layer
  - Domain / DAO layer
- GraphQL Playground enabled
- Easy to extend and maintain

---

## 🗂 Project Structure
```bash
go-graphql/
├── cmd/app
│   ├── domain
│   │   └── dao             # data access objects
│   └── resolvers           # gql query/mutation resolvers
├── config                  # config related files
├── graph
│   ├── generated.go        # generated gql code
│   └── model               # GraphQL models
├── schema
│   └── schema.graphqls     # GraphQL schema definitions
├── tools
│   └── tool.go             # go:build tools
├── gqlgen.yml              # gqlgen config
├── server.go               # entrypoint
├── go.mod
└── go.sum
```

---

## 🧰 Prerequisites

- Go **1.18+** (recommended)
- Go modules enabled
- `gqlgen` installed as a tool

---

## ⚙️ Setup & Installation

### 1. Clone the repository
```bash
git clone https://github.com/carissafarry/go-graphql.git
cd go-graphql
```

### 2. Install tools
Add gqlgen and other tools to your module:
```bash
go install github.com/99designs/gqlgen@v0.17.86
```

### 3. Populate dependencies
This will download dependencies referenced both in your code and via tools.go.
```bash
go mod tidy
```

### 4. Generate GraphQL scaffolding
This reads your schema.graphqls and produces:
- GraphQL runtime code
- Models
- Resolver interfaces
```bash
gqlgen generate
```

---

### 🚀 Running the Server
The server runs at http://localhost:8080 by default.
Here you can write GraphQL queries and mutations interactively.
```bash
go run .
```
---

## 📄 Example Schema

The schema lives in:
```bash
schema/schema.graphqls
```
Example query:
```bash
query {
  users {
    id
    name
  }
}
```

---

## 📚 References
	•	gqlgen Documentation: https://gqlgen.com/
	•	Tutorial: https://medium.com/@wahyubagus1910/implementation-graphql-server-with-golang-fb8f8303b4bc 