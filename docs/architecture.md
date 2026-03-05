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
│   ├── style.css                        # Advanced CSS with glassmorphism
│   └── fireworkEffect.js                # Canvas-based fireworks animation
├── testfiles/
│   ├── ascii/
│   │   ├── bannerHandler_test.go        # Tests for banner file management
│   │   ├── generator_test.go            # Tests for ASCII generation
│   │   └── helperFunctions_test.go      # Tests for helper functions
│   └── server/
│       ├── handlers_test.go             # Tests for HTTP handlers
│       └── errors_test.go               # Tests for error handlers
├── audit/
│   ├── ascii-art-web.md                 # Core functionality audit guide
│   └── optionals/
│       ├── dockerize.md             # Docker implementation audit
│       ├── export-file.md           # File export feature audit
│       └── stylize.md               # Advanced styling audit
├── docs/
│   ├── architecture.md                  # This file - system architecture
│   ├── project_flowchart.md             # Request flow diagrams
│   └── stylisedBackrounds/
│       └── fireworks-effect-technical-analysis.md # Fireworks implementation
├── scripts/
│   ├── build-and-run.sh                    # Docker build script
│   └── docker-version-updater.sh           # Version management script
├── .gitignore                          # Git version control exclusion rules
├── Dockerfile                           # Multi-stage Docker build
├── go.mod                              # Go module definition (v1.25.7)
├── LICENSE                             # MIT License
├── Makefile                            # Build automation and testing
└── README.md                           # Project documentation
```

---

## File Descriptions

### **cmd/main.go**
Application entry point that initializes and starts the HTTP server.

**Functions:**
- `main()` - Calls `EnsureFontFiles()` to download missing banners, registers routes (GET `/`, POST `/ascii-art`, GET `/static/*`), and starts the server on port 8080

**Routes Registered:**
- `http.HandleFunc("/", server.HandleHome)` - Main page
- `http.HandleFunc("/ascii-art", server.HandleAsciiArt)` - ASCII generation
- `http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))` - Static assets

---

### **internal/utilities/ascii/banners/standard.txt**
Standard ASCII font file containing character definitions.

**Structure:**
- 855 lines total (95 characters × 9 lines each)
- ASCII range: 32-126 (printable characters)
- Each character: 1 blank line + 8 art lines
- Classic block-style ASCII art font
- Auto-downloaded from zone01.gr platform

**Usage:**
- Default banner selection
- Clean, readable ASCII output
- Compatible with all printable characters

---

### **internal/utilities/ascii/banners/shadow.txt**
Shadow ASCII font file with 3D shadow effect styling.

**Structure:**
- 855 lines total (95 characters × 9 lines each)
- Shadow/3D effect styling
- Darker, more dramatic appearance
- Auto-downloaded from zone01.gr platform

**Features:**
- 3D shadow effects
- Bold character representation
- Enhanced visual impact
- Professional presentation style

---

### **internal/utilities/ascii/banners/thinkertoy.txt**
Thinkertoy ASCII font file with unique artistic styling.

**Structure:**
- 855 lines total (95 characters × 9 lines each)
- Artistic, creative character designs
- Unique visual style
- Auto-downloaded from zone01.gr platform

**Features:**
- Creative character interpretations
- Artistic flair and personality
- Distinctive visual appearance
- Alternative aesthetic option

---

### **internal/utilities/ascii/bannerHandler.go**
Manages the lifecycle of banner font files including downloading, saving, and loading.

**Functions:**
- `EnsureFontFiles()` - Checks if all three banner files exist in `banners/` directory; downloads from zone01.gr if missing
- `FileExists(path)` - Returns true if file exists at given path
- `DownloadFile(url)` - Performs HTTP GET request and returns file content as byte array
- `SaveFile(path, data)` - Writes byte array to disk with 0644 permissions, creates directories if needed
- `LoadAsciiFile(path)` - Reads banner file and returns array of 855 lines (handles Windows/Unix line endings)

**Constants:**
- Banner URLs for standard, shadow, and thinkertoy fonts from zone01.gr
- Font path: `internal/utilities/ascii/banners/`

---

### **internal/utilities/ascii/generator.go**
Orchestrates the ASCII art generation process by coordinating all helper functions.

**Functions:**
- `GenerateAsciiArt(text, banner)` - Main entry point that loads the specified banner file, filters the input text, splits it into lines, and renders the final ASCII art output

**Process Flow:**
1. Load banner file using `LoadAsciiFile()`
2. Filter input text with `AsciiFilter()`
3. Split into lines with `SplitInputLines()`
4. Render ASCII art with `RenderAscii()`

---

### **internal/utilities/ascii/helperFunctions.go**
Core ASCII processing functions that handle text manipulation and rendering.

**Functions:**
- `SplitInputLines(input)` - Normalizes line endings (converts `\r\n` to `\n`, `\\n` to `\n`) and splits input into array of lines
- `AsciiFilter(text)` - Filters out non-printable characters (keeps ASCII 32-126) while preserving newlines (`\n`, `\r`); returns filtered text and array of removed characters
- `PrintAsciiLine(line, asciiLines)` - Converts a single line of text into 8 rows of ASCII art by mapping each character to its corresponding banner lines using formula: `position = (char - 32) × 9 + row + 1`
- `RenderAscii(lines, asciiLines)` - Processes array of text lines, calls `PrintAsciiLine()` for each non-empty line, adds blank line separators between blocks, and returns complete ASCII art string

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

### **internal/utilities/server/handlers.go**
HTTP request handlers for the web application.

**Functions:**
- `HandleHome(w, r)` - Handles GET requests to `/`; validates path is exactly `/`, parses `templates/index.html`, and executes template with nil data (empty form)
- `HandleAsciiArt(w, r)` - Handles POST requests to `/ascii-art`; validates HTTP method is POST, extracts form values (text, banner), validates banner name is one of {standard, shadow, thinkertoy}, calls `GenerateAsciiArt()`, and re-renders `index.html` with result data

**Validation Logic:**
- Method validation: Only POST allowed for `/ascii-art`
- Banner validation: Must be "standard", "shadow", or "thinkertoy"
- Error handling: Returns appropriate HTTP status codes

---

### **templates/index.html**
Main page template containing the input form, result display area, and interactive effects.

**Features:**
- Text input field (textarea) for user input with placeholder text
- Radio buttons for banner selection (standard, shadow, thinkertoy) with visual previews
- Submit button that sends POST request to `/ascii-art`
- Conditional result display using `{{if .Result}}<pre>{{.Result}}</pre>{{end}}`
- Results displayed on same page after form submission
- Links to CSS stylesheet and JavaScript effects
- Canvas elements for fireworks background animation
- Glassmorphism design with backdrop blur effects

---

### **templates/error.html**
Error page template that displays HTTP error information with consistent styling.

**Features:**
- Displays `{{.StatusCode}}` (400, 404, or 500)
- Displays `{{.Message}}` with error description
- "Go Home" button linking back to `/`
- Consistent styling with main page
- Responsive design

---

### **static/style.css**
Advanced stylesheet implementing modern design patterns and responsive layout.

**Features:**
- **Glassmorphism Design**: Semi-transparent containers with backdrop blur
- **Responsive Layout**: Mobile-first design with media queries
- **Animation Effects**: Smooth transitions and hover effects
- **Typography**: Modern font stack with proper hierarchy
- **Form Styling**: Custom radio buttons and interactive elements
- **Result Display**: Monospace font with proper overflow handling
- **Color Scheme**: Dark theme with light accents
- **Canvas Integration**: Proper layering for background effects

**Key Classes:**
- `.container` - Main glassmorphism container
- `.effect-card` - Interactive banner selection cards
- `.result-container` - ASCII art display area
- `.btn` - Gradient buttons with hover effects

---

### **static/fireworkEffect.js**
Canvas-based JavaScript animation system for interactive background effects.

**Features:**
- **Multi-Canvas System**: Separate canvases for grid, warp, and fireworks
- **Particle Physics**: Realistic firework explosions with gravity
- **Interactive Elements**: Click-to-launch fireworks
- **Auto-Launch**: Automatic fireworks every second
- **Grid Animation**: Animated background grid with color fills
- **Performance Optimized**: Efficient rendering and cleanup

**Classes:**
- `Firework` - Individual firework rockets
- `Particle` - Explosion particles with physics
- `GridEffect` - Background grid animation

---

### **testfiles/ascii/bannerHandler_test.go**
Unit tests for banner file management functions.

**Tests:**
- `TestFileExists_SaveLoad` - Tests `FileExists()`, `SaveFile()`, and `LoadAsciiFile()` using temporary directory
- `TestDownloadFile_HTTPServer` - Tests `DownloadFile()` using mock HTTP server with test data

---

### **testfiles/ascii/generator_test.go**
Unit tests for ASCII art generation orchestration.

**Tests:**
- `TestGenerateAsciiArt_InvalidBanner` - Verifies error handling for invalid banner names
- `TestGenerateAsciiArt_EmptyText` - Tests generation with empty input string

---

### **testfiles/ascii/helperFunctions_test.go**
Unit tests for core ASCII processing functions.

**Tests:**
- `TestSplitInputLines` - Tests newline normalization and splitting with various formats
- `TestAsciiFilter` - Tests character filtering (removes non-printable chars, preserves newlines)
- `TestPrintAsciiLine` - Tests single line ASCII art rendering with 855-line mock data
- `TestRenderAscii` - Tests multi-line rendering with blank line handling

---

### **testfiles/server/handlers_test.go**
Unit tests for HTTP request handlers.

**Tests:**
- `TestHandleHome` - Tests GET `/` returns status 200 with proper template
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

### **audit/ascii-art-web.md**
Core functionality audit guide and test cases.

**Contents:**
- 27 comprehensive test cases
- Functional, General, Basic, and Social test categories
- Step-by-step testing instructions
- Expected outputs and validation criteria
- HTTP status code verification
- Banner output validation

**Test Categories:**
- Package requirements validation
- HTML template presence
- ASCII generation accuracy
- Error handling (400, 404, 500)
- Server stability and performance

---

### **audit/optionals/dockerize.md**
Docker implementation audit guide.

**Contents:**
- 18 Docker-specific test cases
- Dockerfile validation
- Image build verification
- Container execution testing
- Multi-stage build validation
- Security and optimization checks

**Test Areas:**
- Docker image creation
- Container functionality
- Build script automation
- Image size optimization
- Metadata and labeling

---

### **audit/optionals/export-file.md**
File export functionality audit guide.

**Contents:**
- 20 export-related test cases
- HTTP header validation (Content-Type, Content-Length, Content-Disposition)
- File permission testing
- Multiple format support
- Download functionality verification

**Features Tested:**
- File export mechanisms
- Browser download integration
- File format handling
- User experience validation

---

### **audit/optionals/stylize.md**
Advanced styling and UI audit guide.

**Contents:**
- 12 styling-focused test cases
- CSS animation validation
- Responsive design testing
- Interactive element verification
- Visual design consistency
- Performance optimization

**Test Areas:**
- Glassmorphism effects
- Canvas animations
- Mobile responsiveness
- User interaction feedback
- Visual hierarchy and accessibility

---

### **docs/architecture.md**
This file - comprehensive system architecture documentation.

**Contents:**
- Complete repository structure overview
- Detailed file descriptions for all 27+ files
- Technical specifications and implementation details
- Integration points and dependencies
- Performance considerations and error handling

**Sections:**
- Repository structure visualization
- File-by-file technical documentation
- Data flow diagrams
- Banner file structure specifications

---

### **docs/project_flowchart.md**
Detailed request flow diagrams and system workflows.

**Contents:**
- Server startup sequence
- HTTP request handling flows
- ASCII generation process
- Error handling workflows
- Banner file management
- Docker build process

**Diagrams:**
- Visual flowcharts for each major process
- Step-by-step breakdowns
- Decision trees and error paths
- Integration points and data flow

---

### **docs/stylisedBackrounds/fireworks-effect-technical-analysis.md**
Technical documentation for the fireworks animation system.

**Contents:**
- Canvas API implementation details
- Particle physics algorithms
- Performance optimization techniques
- Animation loop architecture
- Event handling mechanisms

**Technical Details:**
- Multi-canvas rendering strategy
- Particle lifecycle management
- Physics simulation (gravity, friction)
- Memory management and cleanup
- Browser compatibility considerations

---

### **scripts/build-and-run.sh**
Simple Docker build and run automation script.

**Features:**
- Builds Docker image with tag `ascii-art-web`
- Runs container with port mapping 8080:8080
- Provides quick deployment workflow
- Handles container cleanup and restart

**Usage:**
```bash
./scripts/build-and-run.sh
```

---

### **scripts/docker-version-updater.sh**
Version management and Docker build automation script.

**Features:**
- Updates version labels in Dockerfile
- Manages build metadata (build date, maintainer)
- Automates Docker image tagging
- Provides versioned releases
- Integrates with CI/CD pipelines

**Functions:**
- Version bumping (major, minor, patch)
- Automated build date injection
- Docker image metadata management

---

### **.gitignore**
Git version control exclusion rules.

**Excluded Items:**
- Binary files (`bin/`, `*.exe`)
- Build artifacts (`coverage.out`)
- IDE files (`.vscode/`, `.idea/`)
- OS files (`.DS_Store`, `Thumbs.db`)
- Temporary files (`*.tmp`, `*.log`)
- Docker build cache

**Benefits:**
- Clean repository
- Reduced repository size
- No sensitive data leaks
- Cross-platform compatibility

---

### **Dockerfile**
Multi-stage Docker build configuration for optimized deployment.

**Stages:**
1. **Builder Stage** (`golang:1.25.7`): Compiles Go application
2. **Runtime Stage** (`alpine:3.19`): Minimal deployment image

**Features:**
- **Optimized Size**: ~34MB final image
- **Security**: Non-root execution, minimal attack surface
- **Metadata**: Build date, version, maintainer labels
- **Dependencies**: CA certificates for HTTPS support
- **Port Exposure**: Port 8080 for web server

---

### **go.mod**
Go module definition and dependency management.

**Contents:**
- Module name: `ascii-art-web`
- Go version requirement: `1.25.7`
- No external dependencies (standard library only)
- Clean dependency tree for security and performance

**Benefits:**
- Reproducible builds
- Version-locked Go runtime
- Zero external attack surface
- Fast compilation and deployment

---

### **LICENSE**
MIT License file defining project usage rights and restrictions.

**Features:**
- MIT License (permissive open source)
- Allows commercial and private use
- Requires attribution
- No warranty or liability
- Compatible with Zone01 requirements

**Rights Granted:**
- Use, copy, modify, merge, publish, distribute
- Sublicense and sell copies
- Full commercial usage rights

---

### **Makefile**
Build automation and development workflow management.

**Targets:**
- `build` - Compile binary to `bin/ascii-art-web`
- `run` - Start development server
- `test` - Run all unit tests with verbose output
- `coverage` - Generate and display test coverage report
- `check` - Format code and run linting
- `clean` - Remove generated files
- `all` - Run complete pre-PR check suite

---

### **README.md**
Comprehensive project documentation and overview.
**Contents:**
- Project description and features
- Installation and usage instructions
- Requirements and dependencies
- API endpoints and error handling
- Author information and license details
