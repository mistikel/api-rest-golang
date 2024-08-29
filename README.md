# golang-rest-api

## Prerequisite tools

- Go version go1.21.0 
- Docker


## Command

- Install Dependency
  ```
  go mod download
  ```

- Test
  ```
  go test -race ./...
  ```

- Run with docker compose
  ```
  docker-compose up
  ```


- Run server without docker, make sure all env for DB is set, then exec mysql migration inside folder ./docs/sql

  ```
  export DB_USER= ""
  export DB_PASSWORD= ""
  export DB_HOST= ""
  export DB_PORT= ""
  export DB_NAME= ""

  go run main.go
  ```


- HTTP Call with empty Param
```
curl --location 'localhost:8080/record'

Response::
{
    "code": 0,
    "msg": "",
    "records": [
        {
            "id": 1,
            "name": "First",
            "totalMarks": 10,
            "createdAt": "2024-08-29T04:12:38Z"
        },
        {
            "id": 2,
            "name": "Second",
            "totalMarks": 22,
            "createdAt": "2024-08-27T08:18:27Z"
        },
        {
            "id": 3,
            "name": "Third",
            "totalMarks": 38,
            "createdAt": "2024-08-28T08:18:26Z"
        },
        {
            "id": 4,
            "name": "Fourth",
            "totalMarks": 60,
            "createdAt": "2024-08-26T08:18:26Z"
        }
    ]
}
```


- HTTP Call with minCount and maxCount Param
```
curl --location 'localhost:8080/record?minCount=20&maxCount=40'

Response:
{
    "code": 0,
    "msg": "",
    "records": [
        {
            "id": 2,
            "name": "Second",
            "totalMarks": 22,
            "createdAt": "2024-08-27T08:18:27Z"
        },
        {
            "id": 3,
            "name": "Third",
            "totalMarks": 38,
            "createdAt": "2024-08-28T08:18:26Z"
        }
    ]
}
```


- HTTP Call with startDate and endDate Param
```
curl --location 'localhost:8080/record?startDate=2024-08-28T04%3A12%3A38Z&endDate=2024-08-29T04%3A12%3A38Z'

Response:
{
    "code": 0,
    "msg": "",
    "records": [
        {
            "id": 3,
            "name": "Third",
            "totalMarks": 38,
            "createdAt": "2024-08-28T08:18:26Z"
        },
        {
            "id": 1,
            "name": "First",
            "totalMarks": 10,
            "createdAt": "2024-08-29T04:12:38Z"
        }
    ]
}
```

- HTTP Call with error state
```
curl --location 'localhost:8080/record?minCount=10'

Response:
{
    "code": 101,
    "msg": "code: 101, message: Failed to validate param, error: Key: 'RecordParam.MaxCount' Error:Field validation for 'MaxCount' failed on the 'required_with' tag"
}
```


