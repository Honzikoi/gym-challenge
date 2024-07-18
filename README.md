# Challenge 2 Flutter & Go

## Work in Progress

This repository contains the backend code for the Challenge 2 project using Golang with Fiber for routing, GORM for ORM, and PostgreSQL as the database. The project is dockerized for easy setup and deployment.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- Docker installed on your machine
- Docker Compose installed on your machine

## How to Launch the Project

To launch the backend locally and check the Swagger APIs, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/your-repo-name.git
    cd your-repo-name
    ```

2. Create a `.env` file in the project root directory and add the following environment variables:
    ```plaintext
    DB_HOST=db
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    ```

3. Run the following command in the project directory with Docker to build and launch the backend:
    ```bash
    docker-compose up --build
    ```

4. To stop and remove the containers, run the following command:
    ```bash
    docker-compose down -v
    ```

## Project Structure


## Swagger 

```bash
swag init -g /routes.go -d ./cmd,./controllers,./models,./handlers --parseDependency --parseInternal 
```