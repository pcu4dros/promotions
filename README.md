# Test project for MyTheresa

## Overview

This project is a **REST API** for handling products with or without discount,
built in **Golang** with **SQLite** as the database.

## Features

- Query products ("/products") with filters (category, price)
- JSON API responses
- Uses **SQLite** for lightweight data storage
- Implements **Dependency Injection** for flexibility and testability

## How to Run

### Prerequisites

#### 🛠️ How to Install Go

If you don't have Go installed, follow these steps:

Go to the [Go's official website](https://go.dev/doc/install) and follow the instructions

### 🌐 Run the Application

MacOs or Linux:

```bash
make run
```

Windows:


```bash
go run cmd/main.go
```

The API will start on [http://localhost:8080/products](http://localhost:8080/products)

## How to run the tests

MacOs or Linux:

```bash
make test
```

Windows:


```bash
go test -v ./...
```

## Design Decisions

### 🚀 Why SQLite?

- Lightweight & Embedded: No need for a separate database server, this is crucial
  to iterate fast.
- Perfect for small/mid-sized apps: Can handle a reasonable number
  of requests efficiently and a reasonable number of results (e.g for this case 200000)
- Easy to Set Up: No external dependencies—just a file-based DB or an in memory db
  pass in the string ":memory:" (this case)
- Relational: so there is more flexibility working with it

### 📦 Why a Single product Package?

- Keeps things simple and modular while still scalable.
  Right now, everything relates to products, so having only a product package
makes sense.
- Future features can be added as separate packages or as additional
  functions in the same package

### ⏳ Why Haven’t I Added More Packages Yet?

- The project is small at this stage -> Adding more packages too early could create
  unnecessary complexity.
- When new features or the scalability demand it, more packages could be added to
  isolate every domain/business logic.

### 🔧 Why Use Dependency Injection?

- Makes Testing Easy -> We can mock dependencies in unit tests.
  (e.g service tests)
- Flexible & Extensible -> Can swap implementations (e.g., replace SQLite with PostgreSQL).
- Decouples Components -> Each part of the app could depend on interfaces,
 not concrete implementations.

### ✨ Why Use Interfaces?

- Abstracts the Database -> The service layer doesn't depend on a specific database.
- Allows Mocking -> Makes unit testing possible without needing a real database.
- Future-Proofing -> Can easily replace critical components without affecting the
rest of the system.

### 🛠️ Why Start with Fewer External Dependencies?

- Reduces Complexity -> Fewer libraries = easier it is to understand the code.
- Faster Iteration -> Less time debugging third-party issues, more time focused
  on core functionality.
- Better Performance -> No unnecessary abstractions slowing things down.
- More Control -> We can decide when and how to introduce dependencies, rather than
  being forced into a framework.
- Easier Maintenance -> Fewer libraries mean fewer breaking changes and dependency
  upgrades in the future.
- Security Benefits -> Fewer dependencies mean fewer potential vulnerabilities
  in the project.

### Future Improvements

- Implement structured logging using libraries like Zap or Logrus for better
observability and performance.
- Add more interfaces if needed to decouple even more the code.
- Conduct benchmark testing to validate performance and ensure the system efficiently
handles 200000+ products.
- Implement pagination for product listings to improve response times and scalability.
- Evaluate the need for Docker containerization and implement it if required for
deployment consistency.
