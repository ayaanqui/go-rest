# 🖥 Go REST API

The repo uses the standard Go library [`net/http`](https://pkg.go.dev/net/http) to build a simple REST API and to learn how Go works.

## Instructions

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
