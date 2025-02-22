# Students API

A simple RESTful API to manage students using SQLite as the database. This project implements basic CRUD operations (Create, Read, Update, Delete) along with additional HTTP methods like `HEAD` to check if a student exists.

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
