# feedback-api
REST API for Feedback App written in Go. To learn more, read below

- [Local Development](#local-development)
- [Docker](#docker)
- [Entities](#entities)
- [API](#api)
- [Database](#database)
- [Project Structure](#project-structure)


# Local Development
## Build
```
make
```
## Run
```
make run-api
```

## Run Unit Tests
```
make test
```

# Docker
The docker environment spins up a postgres database as well as the API to listen on port 8080
## Build
```
docker-compose build
```

## Run
```
docker-compose up
```

# Entities
## Template
Template is used to store shared information when creating feedback for students. For example, if a class is taught and the teacher has already written feedback for that class, a template can be created so that the same format can be used for different students. A template consists of the following properties:
```
{
  "class": "string", // Class ID
  "content": "string", // Template content
}
```

# API
## Template
```
  GET     /v1/templates
  GET     /v1/templates/{id}
  POST    /v1/templates
  PUT     /v1/templates
  DELETE  /v1/templates/{id}
```

# Database
## Migrations
Migrations for updating the postgres db is handled using Robert's [pgmgr](https://github.com/rnubel/pgmgr) utility. The migration scripts can be found under `ops/db/migrate`

# Project structure
This project follows the structure described in Elton Minetto's great blog: [Clean Architecture, 2 years later ](https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/)

- `/entity` holds the models and their core business logic
- `/usecase` holds the overall usecases for the app and structures them as services
- `/api` contains the api-specific logic to run the app (handlers, middleware, router, etc...)
- `/ops` handles any platform operations
