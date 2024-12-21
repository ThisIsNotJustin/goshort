# GoShort!

### URL Shortener Project in Go with Redis, Go Fiber, and Docker

This project is a URL shortener built using Go, Go Fiber, Redis, and Docker for containerization. The application allows you to shorten URLs, manage them, and store URL mappings in a Redis database.

## Features
- URL shortening and redirection
- Redis caching for fast lookups
- RESTful API to handle URL creation, retrieval, edit, and deletion
- Built with Go and Go Fiber for high performance
- Dockerized for easy deployment and development

## Requirements
- [Docker](https://www.docker.com/get-started)
- [Go](https://go.dev/dl/)
- [Redis](https://redis.io/docs/latest/)
- [Postman](https://www.postman.com/) for API testing

## Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/ThisIsNotJustin/goshort.git
```

### 2. Create a .env file in the root of project
```bash
DB_ADDR="db:6379"          # Address for Redis container
DB_PASS=""                 # Password for Redis
APP_PORT=:8080             # Port to run the Go application
DOMAIN="localhost:8000"    # Domain for shortening URLs
API_QUOTA=10               # Max number of URLs
```

### 3. Dockerize the application
```bash
docker-compose up --build
```
This will:
- Build the Go app
- Start Redis in a container
- Run the Go app inside a container

### 4. Test the API with Postman
Test the API endpoints, I recommend using Postman. 
- POST to create a shortened URL from a long URL.

```bash
{
    "url": "https://www.youtube.com/watch?v=xm3YgoEiEDc",
    "short": "",
    "expiry": 24
}
```

- This will receive a response body consisting of
```bash
{
    "url": "https://www.youtube.com/watch?v=xm3YgoEiEDc",
    "short": "localhost:8000/exampleUrl", # random short generated when not provided
    "expiry": 24, # expires in 24 hours
    "rate_limit": 9, # 9 api uses remaining
    "rate_limit_reset": 29 # api uses resets in 29 minutes
}
```