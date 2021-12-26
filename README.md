# 🖥 Go REST API

A simple REST API made using the standard Go library [`net/http`](https://pkg.go.dev/net/http) and PostgreSQL as the database. 

## Instructions

### Set up `db_config.json`
This project also uses a PostgreSQL database in order to run. To start, create a file called `db_config.json` in the project root and place the following in the file, replacing all content within the `[]` with the correct database values:

```json
{
    "db_name": "[database name]",
    "username": "[database username]",
    "password": "[database password]"
}
```

### Generate Binary

```
$ make build
```

or 

```
$ go build src/main.go
```

Creates an executable binary file called `main`. To run this file call `./main`, like so:

```
$ ./main
```

This should start the server on port `3001`.

### Run without Binary

Another way to run the server is by using the `make run` command.

```
$ make run
```

Running the command should also start the server on port `3001`. This command is equivalent to running `go run src/main.go`.
