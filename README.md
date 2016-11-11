# Example HTTP Server

The point of this project is to demonstrate how simple it is to create a robust, multi-functioned API in Golang

1. [Prerequisites](#prerequisites)
2. [Install](#install)
3. [API Info](#api)

## Prerequisites

1. `Docker` (Recommended Docker For Mac)
2. `go` We recommend go 1.7

## Install

To get the server up and running run:
```bash
docker-compose up
```

### Alternatively Without Docker

To run the server without docker and seed data you will need to make the following change:

`config.toml`:
```toml
# mongo
[database]
host = "localhost"
port = 27017

# api documentation
[docs]
url = "../docs/index.html"
```
By doing this you will need to start a local instance of mongo as well:
```
mongod
```

Next is to edit the config file path in `main.go`

`main.go`:
```go
c := config.NewConfig("config.toml")
```
