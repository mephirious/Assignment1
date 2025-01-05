# Advanced Programming 1

Create a concurrent web server using Go's net/http package. The server should
handle multiple incoming requests, process data concurrently using goroutines, coordinate
using channels, use select-case for multiplexing and ensure thread safety with mutex for
shared resources.

## File Structure
```plaintext
./Assignment1/
├── cmd
├── go.mod
├── .http
├── readme.md
├── struct.go
└── server.go
```
   

## How to run:
1. **Library Management System:**
   ```bash
   go run ./cmd/Library/main.go
   ```
2. **Shapes and Geometry:**
   ```bash
   go run ./cmd/Shapes/main.go
   ```
3. **Employee Management System:**
   ```bash
   go run ./cmd/Employee/main.go
   ```
4. **Bank Account System:**
   ```bash
   go run ./cmd/Bank/main.go
   ```
