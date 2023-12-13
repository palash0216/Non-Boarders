![Screenshot 2023-12-09 202610](https://github.com/palash0216/Non-Boarders/assets/75239216/db4fb529-721c-426a-80e1-0bc50727675b)

# Non-Borders 

This repository contains a RESTful API built with Go to facilitate CRUD operations for managing dayscholars in a college setting.

## Features

- **Create:** Add new dayscholars to the system.
- **Read:** Retrieve information about dayscholars.
- **Update:** Modify existing dayscholar details.
- **Delete:** Remove dayscholars from the system.

## Technologies Used

- **GIN Framework:** Fast and lightweight HTTP web framework for Go ([gin-gonic/gin](https://github.com/gin-gonic/gin))

- **DaemonCompile:** [DaemonCompile](https://github.com/daemoncompile)
You can use the go tool to install CompileDaemon:
```bash
go get github.com/githubnemo/CompileDaemon
```
To run Your Code use:
```bash
$ CompileDaemon -command="./MyProgram -my-options"
```
- **Gorm:** An ORM library for Golang, used here with PostgreSQL ([go-gorm/gorm](https://github.com/go-gorm/gorm))
To install Gorm in your system use:
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/"your DB"
```

- **Godotenv:** A Go port of the Ruby dotenv library, used for loading environment variables from a .env file ([joho/godotenv](https://github.com/joho/godotenv))
To install use the command:
```bash
go get github.com/joho/godotenv
```

- **TablePlus:** A modern, native app with an elegant UI for database management ([TablePlus](https://tableplus.com/))

- **Postman:** API development and testing made simple ([Postman](https://www.postman.com/))

## Getting Started: Prerequisites

Before running the API, make sure you have the following tools installed:

- Go (Latest Version)
- DaemonCompile (Latest Version)
- PostgreSQL
- Godotenv
- TablePlus (optional, for database visualization)
- Postman (optional, for testing and documentation)

## Setup

1. **Clone this repository:**

    ```bash
    git clone https://github.com/your-username/your-repo.git
    cd your-repo
    ```

2. **Install dependencies:**

    ```bash
    go mod download
    ```

3. **Create a `.env` file in the root directory and configure your environment variables:**

    ```env
    DB_USER=your_database_user
    DB_PASSWORD=your_database_password
    DB_NAME=your_database_name
    DB_URL="URL"
    PORT=PortNumber
    ```

4. **Run the API:**

    ```bash
    go run main.go
    ```

## API Endpoints

- **GET /Student:** Retrieve a list of all student.
- **GET /student/{id}:** Retrieve details of a specific student.
- **POST /student:** Add a new student to the system.
- **PUT /student/{id}:** Update information for a specific student.
- **DELETE /student/{id}:** Remove a student from the system.

## Usage

1. Use TablePlus or any preferred database management tool to view and manage the database.

2. Use Postman to test the API endpoints. Import the provided Postman collection for convenience.

3. Integrate this API into your frontend application to manage dayscholars seamlessly.


![Screenshot 2023-12-11 182259](https://github.com/palash0216/Non-Boarders/assets/75239216/27d67a38-4097-4f2f-b3ed-903badda8e03)

____________________________________________________________________________________________________________

## Sequence Diagram
(Create using ([Visula Paradigm](https://online.visual-paradigm.com/)))

![Screenshot 2023-12-13 110605](https://github.com/palash0216/Non-Boarders/assets/75239216/b6858309-f587-4e62-a04a-2b23e244031d)




