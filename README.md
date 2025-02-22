# 📚 Students API

A simple RESTful API to manage students using SQLite as the database. This project implements basic CRUD operations (Create, Read, Update, Delete) along with additional HTTP methods like `HEAD` to check if a student exists.

## ✨ Features  

- **Create**: Add a new student.  
- **Read**: Retrieve student details by ID or list all students.  
- **Update**: Modify an existing student's details.  
- **Delete**: Remove a student from the database.  
- **Existence Check**: Use the `HEAD` request to verify if a student exists without retrieving full data.  

---
## ⚙️ Prerequisites  

Before running the project, ensure you have:  

- [Go](https://golang.org/doc/install) (version 1.18+).  
- SQLite3 installed for database management.  
- [Postman](https://www.postman.com/) or any HTTP client for API testing.  

---
## 🚀 Installation  

### 1️⃣ Clone the Repository  

```sh```
git clone https://github.com/madihanazir/students-api.git
cd students-api
### 2️⃣ Install Dependencies

go mod tidy

### 3️⃣ SQLite Setup
SQLite is used as the database, and the required tables will be created automatically when the application runs.

If you want to inspect the database or perform manual queries, install SQLite by following the instructions here.

### ⚙️ Configuration
The project supports configuration via a .env file. Create one and add the following environment variables (optional):
Edit
HTTP_SERVER_ADDR=":8080"
STORAGE_PATH="storage/storage.db"

### ▶️ Running the Application
Start the API server with:
go run cmd/main.go
The API will be available at: http://localhost:8080.

### 🛠️ API Testing with Postman
✅ Install Postman
Download and install Postman.
✅ How to Use Postman for API Testing
Start the server:


go run cmd/main.go
Send API requests:

# Use Postman to make GET, POST, PUT, DELETE, and HEAD requests to test the API.
### 🗄️ Database Setup & Visualization
Option 1: Using TablePlus (Recommended)
Download TablePlus for your OS.
Open TablePlus and click "Create a new connection".
Select "SQLite", then choose storage/storage.db.
Click "Connect" to explore and manage the database visually.
### 📬 Contact
For any questions or suggestions, feel free to reach out:
📧 madihan541@gmail.com
