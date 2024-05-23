# Tritura Software Developer - RESTful API Task Management System

## About

This RESTful API was built using [Go](https://go.dev/) and the [Gin Web Framework](https://github.com/gin-gonic/gin).

## Instructions to Install and Run the Application

1. Clone this repository.
2. run `$ go get -u`
3. run `$ go mod tidy`
4. Run application with `$ go run main.go`
5. Your application is now running locally on port 8080. You can test that your application is running by making a GET request to "http://localhost:8080/test" and you should get back a 200 HTTP response.

## Instructions to Run Tests

1. On the command line at the root of the project, execute: `$ go test`

## API Endpoints

### `POST /tasks`

Creates a new task given a title and description, and the new task will have a randomly generated ID and a status of "Pending".

#### Example Request Body

```
{
    "title": "Task Title",
    "description": "description"
}
```

#### Example Reponse

```
{
    "status": "success",
    "task": {
        "id": 5,
        "title": "Task Title",
        "description": "description",
        "status": "Pending"
    }
}
```

### `GET /tasks`

Returns a list of all tasks.

![image](https://github.com/SaturdayMornings/go-restful-api/assets/24395782/f8e6097a-e8b9-4b50-b50a-47b269296158)

#### Example Reponse

```
[
    {
        "id": 1,
        "title": "Task Title 1",
        "description": "description 1",
        "status": "Pending"
    },
    {
        "id": 2,
        "title": "Task Title 2",
        "description": "description 2",
        "status": "In Progress"
    },
    {
        "id": 3,
        "title": "Task Title 3",
        "description": "description 3",
        "status": "Completed"
    }
]
```

### `GET /tasks/:id`

Returns a tasks based on the provided ID.

#### Example Reponse

```
{
    "id": 2,
    "title": "Sample Title",
    "description": "Sample Description",
    "status": "Pending"
}
```

### `PUT /tasks/:id`

Replaces the task matching the provided ID and returns the updated task.

#### Example Request Body

```
{
    "title": "Updated Task Title",
    "description": "Updated description",
    "status": "Completed"
}
```

#### Example Reponse

```
{
    "status": "success",
    "task": {
        "id": 1,
        "title": "Task Title 1",
        "description": "description 1",
        "status": "Pending"
    }
}
```

### `DELETE /tasks/:id`

Deletes a task matching the provided ID, and will return a success status or error.

#### Example Reponse

```
{
    "status": "success"
}
```
