# ASCII Art Web

Web application that converts text into ASCII art using three banner styles as part of the Zone01 Athens curriculum.

## Features

- Three banner fonts: **standard**, **shadow**, **thinkertoy**
- Web interface with real-time result display
- HTTP error handling (400, 404, 500)
- Automatic banner file management
- Full test coverage (hope the auditors agree - *fingers crossed*)

## Quick Start

**Requirements:** Go 1.20+

```bash
go run ./cmd
```

Open browser: **http://localhost:8080**

## Project Structure

```
cmd/                    # Server entry point
internal/utilities/
  ├── ascii/            # ASCII generation logic
  │   ├── banners/      # Font files (auto-downloaded)
  │   ├── bannerHandler.go
  │   ├── generator.go
  │   └── helperFunctions.go
  └── server/           # HTTP handlers
      ├── handlers.go
      └── errors.go
templates/              # HTML templates
static/                 # CSS file
testfiles/              # Unit tests
  ├── ascii/            # ASCII logic tests
  └── server/           # HTTP handler tests
audit/                  # Manual test cases
docs/                   # Architecture documentation
```

## API Endpoints

| Method | Endpoint     | Description            |
|--------|--------------|------------------------|
| GET    | `/`          | Main page              |
| POST   | `/ascii-art` | Generate ASCII art     |
| GET    | `/static/*`  | Serve static files     |

## Algorithm

1. **Load Banner** - Read 855-line banner file (95 chars × 9 lines)
2. **Filter Input** - Keep ASCII 32-126, preserve newlines
3. **Split Lines** - Handle `\n`, `\r\n`, `\\n`
4. **Render** - Convert each character to 8-row ASCII art
5. **Output** - Combine all rows into final result

**Character Mapping:** `position = (char - 32) × 9 + row + 1`

## Error Handling

- **400** - Invalid method or banner name
- **404** - Unknown route or missing template
- **500** - ASCII generation failure

## Testing

```bash
go test ./testfiles/...
```

**Coverage:** 13 tests across ASCII logic and HTTP handlers

## Documentation

- [Architecture](docs/architecture.md) - Detailed component breakdown
- [Flowchart](docs/project_flowchart.md) - Request flow diagrams
- [Audit Tests](audit/test_cases.md) - Manual test cases

## Authors

- Marios Videnmager
- Iosif Oikonomakis
- Katerina Kasdanastasi

## License

MIT License - See [LICENSE](LICENSE)
