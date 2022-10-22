## WIP : Leave Management System Backend

This project is a backend for a leave management system. It is built using [PocketBase](https://pocketbase.io/) as Framework

### Getting Started

Migrate the collections implemented for the System by running:

```bash
$ go run main.go migrate
```

In order to migrate DB for different environments, you can use the `--dir` flag:

```bash
$ go run main.go migrate --dir="./dev|test|prod_pb_data/"
```

Serving the API:

```bash
$ go run main.go serve
```

or

```bash
$ go run main.go serve --dir="./dev|test|prod_pb_data/"
```
