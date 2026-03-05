# ASCII-Art-Web

[![Go Version](https://img.shields.io/badge/Go-1.25.7-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](Dockerfile)
[![Zone01](https://img.shields.io/badge/Zone01-Athens-orange.svg)](https://zone01.gr/)

A Go web application that converts text into ASCII art using three banner styles. Built as part of the Zone01 Athens curriculum with advanced styling, Docker support, and comprehensive testing.

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Repository Structure](#repository-structure)
- [How to Run](#how-to-run)
- [Requirements](#requirements)
- [Algorithm](#algorithm)
- [Testing](#testing)
- [Docker Support](#docker-support)
- [Documentation](#documentation)
- [API Endpoints](#api-endpoints)
- [Error Handling](#error-handling)
- [Authors](#authors)
- [License](#license)

## Description

This web application provides a modern interface for generating ASCII art from text input. Users can select from three different banner styles (**standard**, **shadow**, **thinkertoy**) and see real-time results with advanced visual effects.

The application features a glassmorphism design with animated fireworks background, responsive layout, and comprehensive error handling. Built following Go best practices with full Docker containerization support.

## Features

### Core Functionality

- **Three Banner Fonts**: standard, shadow, thinkertoy
- **Web Interface**: Real-time ASCII art generation
- **HTTP Error Handling**: Proper 400, 404, 500 responses
- **Automatic Banner Management**: Downloads font files from zone01.gr
- **Input Validation**: ASCII 32-126 character filtering
- **Newline Support**: Handles `\n`, `\r\n`, `\\n` formats

### Advanced Features

- **Modern UI Design**: Glassmorphism with backdrop blur effects
- **Interactive Background**: Canvas-based fireworks animations
- **Responsive Layout**: Mobile, tablet, and desktop optimized
- **Docker Support**: Multi-stage build with Alpine Linux
- **Build Automation**: Makefile with test, coverage, and lint targets

### Professional Structure

- **Modular Architecture**: Separated ASCII logic and HTTP handlers
- **Comprehensive Testing**: 14 unit tests with edge case coverage
- **Clean Code**: Following Go conventions and best practices
- **Documentation**: Detailed architecture and flow diagrams
- **Audit Ready**: Complete test suites for evaluation

## Repository Structure

```
ascii-art-web/
├── cmd/                    # Application entry point
│   └── main.go
├── internal/utilities/     # Core application logic
│   ├── ascii/             # ASCII art generation
│   │   ├── banners/       # Font files (auto-downloaded)
│   │   ├── bannerHandler.go
│   │   ├── generator.go
│   │   └── helperFunctions.go
│   └── server/            # HTTP server components
│       ├── handlers.go
│       └── errors.go
├── templates/             # HTML templates
│   ├── index.html        # Main page with form
│   └── error.html        # Error page template
├── static/               # Frontend assets
│   ├── style.css         # Advanced CSS with animations
│   └── fireworkEffect.js # Canvas-based effects
├── testfiles/            # Comprehensive test suite
│   ├── ascii/           # ASCII logic tests
│   └── server/          # HTTP handler tests
├── audit/               # Audit test guides
│   ├── ascii-art-web.md # Core functionality tests
│   └── optionals/       # Optional feature tests
├── docs/                # Project documentation
├── scripts/             # Build and deployment scripts
│   ├── build-and-run.sh    # Docker build script
│   └── docker-version-updater.sh # Version management
├── Dockerfile           # Multi-stage Docker build
├── Makefile            # Build automation
├── go.mod              # Go module definition
└── README.md
```

## How to Run

### Quick Start

```bash
# Clone and navigate to project
cd ascii-art-web

# Run the application
go run ./cmd

# Open browser
open http://localhost:8080
```

### Using Make

```bash
# Build binary
make build

# Run application
make run

# Run with built binary
./bin/ascii-art-web
```

### Docker Deployment

```bash
# Build Docker image
docker build -t ascii-art-web .

# Run container
docker run -p 8080:8080 ascii-art-web

# Or use build script
./scripts/build-and-run.sh
```

## Requirements

- **Go**: Version 1.25.7 or higher
- **OS**: Linux, macOS, or Windows
- **Dependencies**: None (standard library only)
- **Docker**: Optional, for containerized deployment
- **Browser**: Modern browser with JavaScript support

### Input Format

- Text input via web form (textarea)
- Supports ASCII characters 32-126
- Newline handling: `\n`, `\r\n`, `\\n`
- Banner selection: standard, shadow, thinkertoy

### Output Format

ASCII art displayed in monospace font with proper alignment:

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
```

## Algorithm

The ASCII art generation follows a systematic process:

1. **Banner Loading**: Read 855-line font files (95 chars × 9 lines each)
2. **Input Filtering**: Keep ASCII 32-126, preserve newlines
3. **Line Splitting**: Handle different newline formats
4. **Character Mapping**: Convert each character to 8-row ASCII art
5. **Rendering**: Combine all rows into final output

### Character Mapping Formula

```go
position = (character - 32) × 9 + row + 1
```

**Example for 'A' (ASCII 65):**
- Index: (65 - 32) × 9 = 297
- Rows: asciiLines[298] to asciiLines[305]

### Banner File Structure

- **Total Lines**: 855 per banner file
- **Character Range**: ASCII 32-126 (95 printable characters)
- **Per Character**: 1 blank line + 8 art lines
- **Sources**: Downloaded from zone01.gr

## Testing

### Run All Tests

```bash
# Run tests with verbose output
make test

# Or use go directly
go test -v ./testfiles/...
```

### Test Coverage

```bash
# Generate coverage report
make coverage

# View coverage in browser
go tool cover -html=coverage.out
```

## Docker Support

### Multi-Stage Build

The Dockerfile uses a two-stage build process:

1. **Builder Stage**: golang:1.25.7 for compilation
2. **Runtime Stage**: alpine:3.19 for minimal deployment

### Image Features

- **Size**: ~34MB optimized image
- **Security**: Non-root user, minimal attack surface
- **Metadata**: Version, build date, maintainer labels
- **Certificates**: CA certificates for HTTPS support

### Build Scripts

- `scripts/build-and-run.sh`: Simple build and run script
- `scripts/docker-version-updater.sh`: Version management
- `Makefile`: Integrated Docker targets

## Documentation

- [Architecture](docs/architecture.md) - Detailed component breakdown
- [Project Flow](docs/project_flowchart.md) - Request flow diagrams
- [Fireworks Analysis](docs/stylisedBackrounds/fireworks-effect-technical-analysis.md) - Technical details
- [Audit Tests](audit/) - Comprehensive test guides

## API Endpoints

| Method | Endpoint     | Description                    |
|--------|--------------|--------------------------------|
| GET    | `/`          | Main page with form            |
| POST   | `/ascii-art` | Generate ASCII art             |
| GET    | `/static/*`  | Serve CSS and JavaScript files |

## Error Handling

- **400 Bad Request**: Invalid HTTP method or banner name
- **404 Not Found**: Unknown routes or missing templates
- **500 Internal Server Error**: ASCII generation failures

All errors display user-friendly pages with "Go Home" navigation.

## Authors

- **Marios Videnmager**  
- **Iosif Oikonomakis**  
- **Katerina Kasdanastasi**  

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
