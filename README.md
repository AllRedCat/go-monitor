# go-monitor

## Overview

**go-monitor** is a system monitoring tool written in Go. It provides APIs and WebSocket endpoints to monitor system metrics, CPU usage, Docker containers, and execute commands remotely. The project is modular, with clear separation between API routing, handlers, models, services, and utilities.

## Features

- System metrics monitoring (CPU, memory, etc.)
- Real-time CPU usage via WebSocket
- Docker container management and monitoring
- Remote command execution
- RESTful API endpoints

## Folder Structure

```
go.mod, go.sum         # Go module files
README.md              # Project documentation
cmd/
	api/
		main.go            # Application entry point
internal/
	api/
		routes.go          # API route definitions
		handlers/          # HTTP/WebSocket handlers
			docker.go        # Docker-related handlers
			exec.go          # Command execution handlers
			system.go        # System metrics handlers
			ws_cpu.go        # WebSocket CPU usage handlers
			ws_metrics.go    # WebSocket metrics handlers
	models/              # Data models
		cpu_usage.go, cpu.go, exec.go, metrics.go
	services/            # Business logic/services
		cpu.go, docker.go, exec.go, metrics.go, system.go
	utils/
		websocket.go       # WebSocket utility functions
```

## Getting Started

### Prerequisites

- Go 1.25 or newer

### Installation

Clone the repository:

```bash
git clone https://github.com/AllRedCat/go-monitor.git
cd go-monitor
```

Install dependencies:

```bash
go mod tidy
```

### Running the Application

```bash
go run ./cmd/api/main.go
```

## API Endpoints

The API exposes endpoints for system metrics, Docker management, and command execution. WebSocket endpoints provide real-time updates for CPU and metrics.

Refer to the code in `internal/api/routes.go` and `internal/api/handlers/` for details.

You can see the endpoints documentation in [routes](/API_ENDPOINTS.md)

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements or new features.

## License

This project is licensed under the MIT License.
