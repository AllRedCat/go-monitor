# API Endpoint Documentation

This document describes each endpoint provided by go-monitor, including method, path, request/response format, and functionality.

---

## REST Endpoints

### 1. GET `/info`
- **Description:** Returns system information (CPU, memory, etc.).
- **Method:** GET
- **Request Body:** None
- **Response:** JSON object with system info
- **Errors:**
  - 405 Method Not Allowed (if not GET)
  - 500 Internal Server Error (on failure)

---

### 2. GET `/docker`
- **Description:** Returns information about Docker containers.
- **Method:** GET
- **Request Body:** None
- **Response:** JSON array of containers
- **Errors:**
  - 405 Method Not Allowed (if not GET)
  - 204 No Content (if no containers)
  - 500 Internal Server Error (on failure)

---

### 3. POST `/exec`
- **Description:** Executes a script or command on the server.
- **Method:** POST
- **Request Body:** JSON object `{ "path": "<script_path>" }`
- **Response:** JSON object with execution result
- **Errors:**
  - 405 Method Not Allowed (if not POST)
  - 400 Bad Request (missing or invalid body)
  - 500 Internal Server Error (on failure)

---

## WebSocket Endpoints

### 4. `/ws/info`
- **Description:** Streams system metrics (CPU, memory, etc.) every 3 seconds.
- **Protocol:** WebSocket
- **Message Format:** JSON object with metrics
- **Errors:**
  - WebSocket upgrade failure
  - Internal errors logged, connection closes on error

---

### 5. `/ws/cpu`
- **Description:** Streams CPU usage data every 3 seconds.
- **Protocol:** WebSocket
- **Message Format:** JSON object with CPU usage
- **Errors:**
  - WebSocket upgrade failure
  - Internal errors logged, connection closes on error

---

## Notes
- All endpoints are defined in `internal/api/routes.go` and handled in `internal/api/handlers/`.
- WebSocket endpoints require a client that supports WebSocket protocol.
- For more details, see the handler source files.
