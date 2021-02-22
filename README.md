# feedback-api
REST API for Feedback App written in Go

# Entities
## Template
Template is used to store shared information when creating feedback for students. For example, if a class is taught and the teach has already written feedback for that class, a template can be created so that the same format can be used for different students. A template consists of the following properties:
```
{
  "class": "string", // Class ID
  "content": "string", // Template content
}
```

# API
## Template
```
  GET     /v1/template
  GET     /v1/template/{id}
  POST    /v1/template
  PUT     /v1/template
  DELETE  /v1/template/{id}
```

# Project structure
This project follows the structure described in Elton Minetto's great blog: [Clean Architecture, 2 years later ](https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/) 
