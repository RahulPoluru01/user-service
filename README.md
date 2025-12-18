## User with DOB and Calculated Age — Go Development Project->

This Go project helps users to store users data i.e NAME and DOB into Postgres DB and performs crud operations via API and addition to that 
it generates age dynamically based on DOB and while fetching user details.

---

## Tech Stack->

- [GoFiber](https://gofiber.io/) #Go web framework
- SQL (PostgreSQL) + [SQLC](https://sqlc.dev/) #for type safe code
- [Uber Zap](https://github.com/uber-go/zap) #for logging
- [go-playground/validator](https://github.com/go-playground/validator) #for input validation

---

## Project Structure

```text
user-service/
├── cmd/
│   └── server/
│       └── main.go            application entry point
│
├── config/                    configuration files
│
├── db/
│   ├── migrations/            database schema changes
│   │   └── 001_create_users.sql
│   │
│   ├── queries/               raw SQL used by SQLC
│   │   └── users.sql
│   │
│   └── sqlc/                  SQLC generated code
│       ├── db.go
│       ├── models.go
│       └── users.sql.go
│
├── internal/                  core application code
│   ├── handler/               HTTP request handlers
│   │   ├── health_handler.go
│   │   ├── user_handler.go
│   │   └── validator.go
│   │
│   ├── repository/            data access layer
│   │   ├── db.go
│   │   ├── user_repository.go
│   │   └── user_repository_impl.go
│   │
│   ├── service/               business logic
│   │   └── user_service.go
│   │
│   ├── routes/                route registration
│   │   └── routes.go
│   │
│   ├── middleware/            request lifecycle logic
│   │   ├── request_id.go
│   │   └── request_logger.go
│   │
│   └── logger/                logging setup
│       └── logger.go
│
├── go.mod
├── go.sum
├── sqlc.yaml
├── .gitignore
└── README.md
```
---

## Quick Start

Prerequisites
- go latest version
- postgresql 18+
- vscode go extension

## Installations
1.Clone the repository

```text
git clone https://github.com/RahulPoluru01/user-service.git

cd user-service
```

2.Database setup and schema migration
- Open PostgreSQL (pgAdmin or psql)
- Create database
  ```text
   CREATE DATABASE user_service;
  ```
- Connect to the database
 ```text
   \c user_service
```
- Run schema migration
```text
   \i db/migrations/001_create_users.sql
```
- Verify table creation
 ```
   \dt
```

3.Configure database connection

Update the PostgreSQL DSN in: 

cmd/server/main.go
```text
- postgres://postgres:password_here@localhost:5432/user_service?sslmode=disable
```
- Ensure the username, password, and database name match your local setup.

4.Run the application

From the project root:
```text
  go run cmd/server/main.go
```

The server starts on: 
http://localhost:3000


## Architecture Overview
### Frontend
- postman or browser

### Backend
- built with gofiber

- database users table:
```text

Field	|Type	 |Constraints
id	  |SERIAL|PRIMARY KEY
name	|TEXT	 |NOT NULL
dob	  |DATE	 |NOT NULL

```
---

## API usage:

### Create user:

POST /users
{
  "name": "Alice",
  "dob": "1990-05-10"
}


### Get user by ID:

GET /users/1


### List users:

GET /users


### Update user:

PUT /users/1
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}

### Delete user:

DELETE /users/1

---

## example screenshots:

<img width="720" height="641" alt="Screenshot (202)" src="https://github.com/user-attachments/assets/55b4686b-5348-49d9-8eff-1a9ac186024c" />




<img width="788" height="815" alt="Screenshot (203)" src="https://github.com/user-attachments/assets/f3521951-eed3-43ea-be85-4fc7595538b9" />




<img width="1030" height="375" alt="Screenshot (204)" src="https://github.com/user-attachments/assets/f1a48b5e-3b99-4f2d-ba5c-d6308bd5e0a9" />




<img width="792" height="646" alt="Screenshot (205)" src="https://github.com/user-attachments/assets/68d17c1d-41e3-4165-9fbf-61599502e0f9" />

---
