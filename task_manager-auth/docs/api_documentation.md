# Task Manager API Documentation

Welcome to the Task Manager API! This API allows you to manage tasks with basic Create, Read, Update, and Delete (CRUD) operations.

## Getting Started

### Prerequisites
- Go installed (version 1.16 or higher recommended)
- Git (optional, for cloning the repository)

### Setup
1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd task_manager
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```
   The server will start on `localhost:8080` by default.

---

## API Endpoints

### 1. Get All Tasks
- **Endpoint:** `GET /tasks`
- **Description:** Returns a list of all tasks.
- **Response Example:**
  ```json
  [
    {"id": 1, "title": "Sample Task", "completed": false}
  ]
  ```

### 2. Get Task by ID
- **Endpoint:** `GET /tasks/{id}`
- **Description:** Returns a single task by its ID.
- **Response Example:**
  ```json
  {"id": 1, "title": "Sample Task", "completed": false}
  ```

### 3. Create a Task
- **Endpoint:** `POST /tasks`
- **Description:** Creates a new task.
- **Request Body Example:**
  ```json
  {"title": "New Task", "completed": false}
  ```
- **Response Example:**
  ```json
  {"id": 2, "title": "New Task", "completed": false}
  ```

### 4. Update a Task
- **Endpoint:** `PUT /tasks/{id}`
- **Description:** Updates an existing task by its ID.
- **Request Body Example:**
  ```json
  {"title": "Updated Task", "completed": true}
  ```
- **Response Example:**
  ```json
  {"id": 1, "title": "Updated Task", "completed": true}
  ```

### 5. Delete a Task
- **Endpoint:** `DELETE /tasks/{id}`
- **Description:** Deletes a task by its ID.
- **Response:**
  - Status 204 No Content

---

