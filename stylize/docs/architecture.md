# Architecture

## Repository Structure

```
ascii-art-web/
├── cmd/
│   └── main.go                          # Server entry point
├── internal/
│   └── utilities/
│       ├── ascii/
│       │   ├── banners/                 # Auto-downloaded font files
│       │   │   ├── standard.txt         # Standard ASCII font (855 lines)
│       │   │   ├── shadow.txt           # Shadow ASCII font (855 lines)
│       │   │   └── thinkertoy.txt       # Thinkertoy ASCII font (855 lines)
│       │   ├── bannerHandler.go         # Banner file management
│       │   ├── generator.go             # ASCII art generation orchestrator
│       │   └── helperFunctions.go       # Core ASCII processing functions
│       └── server/
│           ├── handlers.go              # HTTP request handlers
│           └── errors.go                # HTTP error responses
├── templates/
│   ├── index.html                       # Main page with form and results
│   └── error.html                       # Error page template
├── static/
│   └── style.css                        # Page styling
├── testfiles/
│   ├── ascii/
│   │   ├── bannerHandler_test.go        # Tests for banner file management
│   │   ├── generator_test.go            # Tests for ASCII generation
│   │   └── helperFunctions_test.go      # Tests for helper functions
│   └── server/
│       ├── handlers_test.go             # Tests for HTTP handlers
│       └── errors_test.go               # Tests for error handlers
├── audit/
│   └── test_cases.md                    # Manual test cases for auditing
├── docs/
│   ├── architecture.md                  # Architecture of the repository
│   └── project_flowchart.md             # Flow diagrams
├── go.mod                               # Go module definition
├── LICENSE                              # MIT License
└── README.md                            # Project documentation
```

---

## File Descriptions

### **cmd/main.go**
Application entry point that initializes and starts the HTTP server.

**Functions:**
- `main()` - Calls `EnsureFontFiles()` to download missing banners, registers three routes (GET `/`, POST `/ascii-art`, GET `/static/*`), and starts the server on port 8080

---

### **internal/utilities/ascii/bannerHandler.go**
Manages the lifecycle of banner font files including downloading, saving, and loading.

**Functions:**
- `EnsureFontFiles()` - Checks if all three banner files exist in `banners/` directory; downloads from zone01.gr if missing
- `FileExists(path)` - Returns true if file exists at given path
- `DownloadFile(url)` - Performs HTTP GET request and returns file content as byte array
- `SaveFile(path, data)` - Writes byte array to disk with 0644 permissions
- `LoadAsciiFile(path)` - Reads banner file and returns array of 855 lines (handles Windows/Unix line endings)

**Constants:**
- Banner URLs for standard, shadow, and thinkertoy fonts from zone01.gr
- Font path: `internal/utilities/ascii/banners`

---

### **internal/utilities/ascii/generator.go**
Orchestrates the ASCII art generation process by coordinating all helper functions.

**Functions:**
- `GenerateAsciiArt(text, banner)` - Main entry point that loads the specified banner file, filters the input text, splits it into lines, and renders the final ASCII art output

---

### **internal/utilities/ascii/helperFunctions.go**
Core ASCII processing functions that handle text manipulation and rendering.

**Functions:**
- `SplitInputLines(input)` - Normalizes line endings (converts `\r\n` to `\n`, `\\n` to `\n`) and splits input into array of lines
- `AsciiFilter(text)` - Filters out non-printable characters (keeps ASCII 32-126) while preserving newlines (`\n`, `\r`); returns filtered text and array of removed characters
- `PrintAsciiLine(line, asciiLines)` - Converts a single line of text into 8 rows of ASCII art by mapping each character to its corresponding banner lines using formula: `position = (char - 32) × 9 + row + 1`
- `RenderAscii(lines, asciiLines)` - Processes array of text lines, calls `PrintAsciiLine()` for each non-empty line, adds blank line separators between blocks, and returns complete ASCII art string

---

### **internal/utilities/server/handlers.go**
HTTP request handlers for the web application.

**Functions:**
- `HandleHome(w, r)` - Handles GET requests to `/`; validates path is exactly `/`, parses `templates/index.html`, and executes template with nil data (empty form)
- `HandleAsciiArt(w, r)` - Handles POST requests to `/ascii-art`; validates HTTP method is POST, extracts form values (text, banner), validates banner name is one of {standard, shadow, thinkertoy}, calls `GenerateAsciiArt()`, and re-renders `index.html` with result data

---

### **internal/utilities/server/errors.go**
HTTP error response handlers that render error pages with appropriate status codes.

**Type:**
- `ErrorData` - Struct containing `StatusCode` (int) and `Message` (string) passed to error template

**Functions:**
- `Send400(w, message)` - Sets HTTP status 400 Bad Request, renders `error.html` with message (used for invalid method or banner name)
- `Send404(w, message)` - Sets HTTP status 404 Not Found, renders `error.html` with message (used for unknown routes or missing templates)
- `Send500(w, message)` - Sets HTTP status 500 Internal Server Error, renders `error.html` with message (used for ASCII generation failures)

---

### **templates/index.html**
Main page template containing the input form and result display area.

**Features:**
- Text input field (textarea) for user input
- Radio buttons for banner selection (standard, shadow, thinkertoy)
- Submit button that sends POST request to `/ascii-art`
- Conditional result display using `{{if .Result}}<pre>{{.Result}}</pre>{{end}}`
- Results displayed on same page after form submission
- Links to CSS stylesheet

---

### **templates/error.html**
Error page template that displays HTTP error information.

**Features:**
- Displays `{{.StatusCode}}` (400, 404, or 500)
- Displays `{{.Message}}` with error description
- "Go Home" button linking back to `/`

---

### **static/style.css**
Stylesheet for all pages.

**Features:**
- Page layout and form styling
- `.result` class with `overflow-x: auto` and `max-width: 100%` to handle long ASCII art lines
- Error page styling
- Responsive design elements

---

### **testfiles/ascii/bannerHandler_test.go**
Unit tests for banner file management functions.

**Tests:**
- `TestFileExists_SaveLoad` - Tests `FileExists()`, `SaveFile()`, and `LoadAsciiFile()` using temporary directory
- `TestDownloadFile_HTTPServer` - Tests `DownloadFile()` using mock HTTP server

---

### **testfiles/ascii/generator_test.go**
Unit tests for ASCII art generation.

**Tests:**
- `TestGenerateAsciiArt_InvalidBanner` - Verifies error handling for invalid banner names
- `TestGenerateAsciiArt_EmptyText` - Tests generation with empty input string

---

### **testfiles/ascii/helperFunctions_test.go**
Unit tests for core ASCII processing functions.

**Tests:**
- `TestSplitInputLines` - Tests newline normalization and splitting
- `TestAsciiFilter` - Tests character filtering (removes non-printable chars)
- `TestPrintAsciiLine` - Tests single line ASCII art rendering with 855-line mock data
- `TestRenderAscii` - Tests multi-line rendering with blank line handling

---

### **testfiles/server/handlers_test.go**
Unit tests for HTTP request handlers.

**Tests:**
- `TestHandleHome` - Tests GET `/` returns status 200
- `TestHandleAsciiArt_InvalidMethod` - Tests GET request to `/ascii-art` returns 400
- `TestHandleAsciiArt_InvalidBanner` - Tests POST with invalid banner name returns 400

---

### **testfiles/server/errors_test.go**
Unit tests for error handlers.

**Tests:**
- `TestSend400` - Verifies 400 status code and error template rendering
- `TestSend404` - Verifies 404 status code and error template rendering
- `TestSend500` - Verifies 500 status code and error template rendering

---

## Data Flow

```
User Input → HandleAsciiArt → GenerateAsciiArt → LoadAsciiFile
                                              ↓
                                         AsciiFilter
                                              ↓
                                      SplitInputLines
                                              ↓
                                         RenderAscii
                                              ↓
                                       PrintAsciiLine (per line)
                                              ↓
                                      Return ASCII Art
```

## Banner File Structure

- **Total lines:** 855 (95 characters × 9 lines)
- **ASCII range:** 32-126 (printable characters)
- **Per character:** 1 blank line + 8 art lines
- **Character mapping:** `position = (char - 32) × 9 + row + 1`
