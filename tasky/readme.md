## Tasky
Tasky is a Task Management Web App.
It involves common features like authentication and integrates Go,React and PostgreSql.

## Project Idea For Tasky
A task management app where users can:
1. signup and login
2. create,edit and delete tasks
3. mark tasks as complete or incomplete
4. organize task into categories(eg. work,personal)
5. view completed tasks in a summary.

## Tech Stack
Backend (Go) - build a REST API using the Fiber framework (fast and lightweight)
Database - Use PostgreSQL
Frontend (React)
Authentication: JWT for token-based sessions

## Features and API endpoints
### Authentication
1. POST /signup: register a user
2. POST /login: authenticate a user and return a JWT.

### Tasks
1. POST /task - create a new task
2. GET /task - retrieve all taks for the logged-in user
3. PUT /tasks/:id - update a task
4. DELETE /taks/:id - delete a task

### Summary
1. GET /tasks/summary: Retrieve completed task statistics.



## STEP 1: Database Scheme Design
Use postgresql to store users and tasks.

Tables
1. Users Table
Responsible for storing user credentials and profile details.
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY, --unique id for each user
    email VARCHAR(255) UNIQUE NOT NULL, -- user's email (unique)
    password_hash TEXT NOT NULL, -- hashed password
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- time for registration
);
```

2. Tasks Table
```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY, -- unique id for each task
    user_id INT NOT NULL, -- reference to the user who owns the task
    title VARCHAR(256) NOT NULL, -- task title
    description TEXT, -- optional task description
    is_completed BOOLEAN DEFAULT FALSE, --status of the task
    category VARCHAR(50), -- task category (eg, work,personal)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE; --links to users table.
); -- remember semi-colon
```

## Create the database
```sql
CREATE DATABASE tasky; --remember the semi-colon

-- switch to the tasky database
\c tasky

-- after creating want to make sure the table was created.
\dt

-- list all databases
\l
```

The `.env` file contains database credentials

## Changing password of postgres
To change password of postgres user.
```sql
ALTER USER postgres WITH PASSWORD 'your_password';
```

Things to notice:
- the password should be in 'single quotes
- remember the semi colon;

## Writing database interaction functions
1. create a models package inside your project, this holds all the database interaction code.
2. create `user.go` for user functions add it to `models/user.go`.

