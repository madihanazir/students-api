# Students API

A simple RESTful API to manage students using SQLite as the database. This project implements basic CRUD operations (Create, Read, Update, Delete) along with additional HTTP methods like `HEAD` to check if a student exists.
# 📂 Project Structure
pgsql
Copy
Edit
students-api/
│── cmd/
│   └── main.go       # Main entry point
│── storage/
│   ├── storage.go    # SQLite database connection
│── handlers/
│   ├── student.go    # Student API handlers
│── models/
│   ├── student.go    # Student struct
│── storage/
│   ├── storage.db    # SQLite database (auto-created)
│── README.md         # Documentation
│── go.mod            # Go module file

## Features

- **Create**: Add a new student.
- **Read**: Get a student's details by their ID, or list all students.
- **Update**: Update an existing student's details.
- **Delete**: Remove a student.
- **Check if a student exists**: Use the `HEAD` request to verify if a student is available without retrieving the entire data.

## Prerequisites

- [Go](https://golang.org/doc/install) version 1.18+ must be installed on your machine.
- SQLite3 for database management.
- [Postman](https://www.postman.com/) or any HTTP client for testing the API.

## Installation

### Clone the Repository
To get started, clone the repository to your local machine:
```bash
git clone https://github.com/madihanazir/students-api.git
cd students-api
```
## Install Dependencies
The project depends on some external libraries which can be installed using go mod:

```bash
Copy
Edit
go mod tidy
```
## SQLite Setup
SQLite is used as the database, and the necessary tables will be created automatically when you run the application.

If you want to inspect the database or perform manual queries, you can install SQLite by following the instructions here.

## Configuration
The project comes with a configuration file under internal/config. To use the default configuration, simply create a .env file and add the following variables (optional):

```bash
Copy
Edit
```
HTTP_SERVER_ADDR=":8080"
STORAGE_PATH="storage/storage.db"
## Running the Application
To start the API server, run:

```bash
Copy
Edit
```
go run main.go
The API will be available at http://localhost:8080.
## API Testing with Postman
Postman allows you to test your API endpoints easily.

## Installation:
-Download and install Postman.
-How to Use Postman for API Testing:
-Start the Server
-Make sure your server is running:

sh:
go run cmd/main.go
The server should be listening at your localhost.

## 🗄️ Database Setup & Visualization
-Option 1: Using TablePlus (Recommended)
-Download TablePlus for your OS.
-Open TablePlus and click "Create a new connection".
-Select "SQLite", then choose storage/storage.db.
-Click "Connect" to explore and manage the database visually.

Feel free to reach me at madihan541@gmail.com

