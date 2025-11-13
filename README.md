# Eteleverse API

Go web API application with health check and hello world endpoints.

## Features

- Health check endpoint at `/health`
- Hello world endpoint at `/hello`
- Built with go-chi router
- Dockerized application
- Jenkins CI/CD pipeline

## Running with Docker Compose

```bash
docker-compose up -d
```

## API Endpoints

### Health Check
```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "message": "Service is healthy",
  "status": "ok"
}
```

### Hello World
```bash
curl http://localhost:8080/hello
```

Response:
```json
{
  "message": "Hello, World!",
  "status": "success"
}
```

## Running Tests

```bash
go test -v ./test/...
```

## Building Locally

```bash
go build -o main .
./main
```

## Jenkins Pipeline

The Jenkinsfile includes the following stages:
- Checkout
- Build
- Test
- Docker Build
- Docker Deploy
- Health Check
