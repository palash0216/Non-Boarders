
# Non-Borders API

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
  go get -u gorm.io/driver/sqlite
  ```

- **Godotenv:** A Go port of the Ruby dotenv library, used for loading environment variables from a .env file ([joho/godotenv](https://github.com/joho/godotenv))
  To install use the command:
  ```bash
  go get github.com/joho/godotenv
  ```

- **TablePlus:** A modern, native app with an elegant UI for database management ([TablePlus](https://tableplus.com/))

- **Postman:** API development and testing made simple ([Postman](https://www.postman.com/))

## Prerequisites

Before running the API, make sure you have the following tools installed:

- Go (at least version X.X.X)
- DaemonCompile (version X.X.X)
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
    ```

4. **Run the API:**

    ```bash
    go run main.go
    ```

## API Endpoints

- **GET /dayscholars:** Retrieve a list of all dayscholars.
- **GET /dayscholars/{id}:** Retrieve details of a specific dayscholar.
- **POST /dayscholars:** Add a new dayscholar to the system.
- **PUT /dayscholars/{id}:** Update information for a specific dayscholar.
- **DELETE /dayscholars/{id}:** Remove a dayscholar from the system.

## Usage

1. Use TablePlus or any preferred database management tool to view and manage the database.

2. Use Postman to test the API endpoints. Import the provided Postman collection for convenience.

3. Integrate this API into your frontend application to manage dayscholars seamlessly.

## Contributing

Feel free to contribute to this project by opening issues or submitting pull requests. Your feedback is highly appreciated.


