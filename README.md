# gorunner

`gorunner` is a CLI tool written in Go that monitors changes in Go files and automatically restarts the server whenever files in the specified directory are modified. It works similarly to `nodemon`, but is tailored specifically for Go projects.

## Features

- Monitors all `.go` files in the directory of the specified main file.
- Automatically restarts the Go server on file changes.
- Lightweight and built using only Go's standard library.

## Installation

1. Clone the repository or download the `gorunner` code.
2. Build the binary:

   ```bash
   go build -o gorunner
