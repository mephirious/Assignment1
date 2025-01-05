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
```bash
   go run main.go
```

## Features:
1. **POST /data:**
Accepts data in JSON format (e.g., {"key": "value"}) and stores it in
an in-memory database (a map). 
2. **GET /data:**
Returns the entire in-memory database as JSON
3. **GET /stats:**
Returns the number of requests handled so far. 
4. **DELETE /data/{key}**
Deletes a specific key from the in-memory database. 