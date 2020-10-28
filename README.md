# Project-Demo BE

This API is an example.

## Installation

Install the dependencies that are on the go.mod file.

## Usage

At first you need to create an .env file where the API connection info will be located.

DB_USER: A user with permissions in your database.

DB_HOST: Where your database engine runs, like localhost.

DB_PROT: Your db engine exposed port.



```.env
DB_NAME = <db_name>
DB_USER = <db_user>
DB_DRIVER = postgres
DB_HOST = <your_host>
DB_PORT = <your_host_port>
```

Then, run the following:

```go
go run main.go
```