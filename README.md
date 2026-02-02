# ASCII ART WEB

## Description

ASCII Art Web is a web application that generates ASCII art from text input using different banner styles. The application provides a clean web interface where users can enter text, select from three different banner styles (standard, shadow, thinkertoy), and generate ASCII art output.

## Authors

- Your Team Name

## Usage: How to Run

1. Make sure you have Go installed on your system
2. Clone or download the project
3. Navigate to the project directory
4. Run the server:
   ```bash
   go run main.go
   ```
5. Open your web browser and go to `http://localhost:8080`
6. Enter your text, select a banner style, and click "Generate ASCII Art"

## Implementation Details: Algorithm

### Server Architecture
- **HTTP Server**: Built using Go's standard `net/http` package
- **Template Engine**: Uses Go's `html/template` for rendering HTML pages
- **Static File Serving**: Serves CSS and other static assets from `/static/` directory

### HTTP Endpoints
- `GET /`: Serves the main homepage with the input form
- `POST /ascii-art`: Processes form data and generates ASCII art
- `GET /static/*`: Serves static files (CSS, images, etc.)

### Error Handling
The application implements proper HTTP status codes:
- **200 OK**: Successful requests
- **400 Bad Request**: Invalid input (empty text, invalid banner)
- **404 Not Found**: Page not found
- **405 Method Not Allowed**: Wrong HTTP method
- **500 Internal Server Error**: Server-side errors

### File Structure
```
ascii-art-web/
├── main.go              # Main server file
├── templates/           # HTML templates
│   ├── index.html       # Homepage
│   ├── result.html      # Result page
│   └── error.html       # Error page
├── static/              # Static assets
│   └── css/
│       └── style.css    # Stylesheet
└── README.md           # This file
```

### ASCII Art Generation
The current implementation includes a placeholder ASCII art generator that demonstrates the concept. For production use, this should be replaced with the actual ASCII art generation logic from your previous ascii-art project.

### Features
- Responsive web design
- Clean, modern UI
- Form validation
- Error handling with custom error pages
- Support for three banner types: standard, shadow, thinkertoy
- Mobile-friendly interface