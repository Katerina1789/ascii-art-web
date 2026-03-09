# Project Flow

## Server Startup

```
main()
  ├─ EnsureFontFiles()
  │   ├─ Check banners/ directory
  │   ├─ Download missing standard.txt from zone01.gr
  │   ├─ Download missing shadow.txt from zone01.gr
  │   └─ Download missing thinkertoy.txt from zone01.gr
  ├─ Register HTTP routes:
  │   ├─ GET  /          → HandleHome
  │   ├─ POST /ascii-art → HandleAsciiArt
  │   ├─ POST /export    → ExportAsciiArt
  │   └─ GET  /static/*  → FileServer (CSS, JS)
  └─ Start HTTP server on :8080
      └─ Listen for incoming requests
```

## GET / (Home Page)

```
HandleHome(w, r)
  ├─ Validate request path == "/"
  │   └─ NO  → Send404("Page not found")
  ├─ Parse templates/index.html
  │   └─ ERROR → Send404("Template not found")
  └─ Execute template with nil data
      ├─ Render form with:
      │   ├─ Text input textarea
      │   ├─ Banner selection (standard/shadow/thinkertoy)
      │   └─ Submit button
      ├─ Load CSS and JavaScript assets
      ├─ Initialize fireworks canvas
      └─ Return HTML to client
```

## POST /ascii-art (Generate ASCII)

```
HandleAsciiArt(w, r)
  ├─ Validate HTTP method == POST
  │   └─ NO → Send400("Only POST method allowed")
  ├─ Extract form values:
  │   ├─ text := r.FormValue("text")
  │   └─ banner := r.FormValue("banner")
  ├─ Validate banner ∈ {standard, shadow, thinkertoy}
  │   └─ NO → Send400("Invalid banner name")
  ├─ GenerateAsciiArt(text, banner)
  │   ├─ LoadAsciiFile(bannerPath)
  │   │   ├─ Read 855 lines from banner file
  │   │   └─ Handle Windows/Unix line endings
  │   ├─ AsciiFilter(text)
  │   │   ├─ Keep ASCII chars 32-126
  │   │   ├─ Preserve newlines (\n, \r)
  │   │   └─ Remove non-printable characters
  │   ├─ SplitInputLines(filteredText)
  │   │   ├─ Normalize \r\n → \n
  │   │   ├─ Normalize \\n → \n
  │   │   └─ Split by \n into array
  │   └─ RenderAscii(lines, asciiLines)
  │       └─ For each line:
  │           ├─ PrintAsciiLine(line, asciiLines)
  │           │   └─ Generate 8 rows per character
  │           └─ Add blank line separator
  ├─ Check for generation errors
  │   └─ ERROR → Send500("Failed to generate ASCII art")
  ├─ Parse templates/index.html
  │   └─ ERROR → Send404("Template not found")
  └─ Execute template with result data:
      ├─ Text: original input
      ├─ Banner: selected banner name
      ├─ Result: generated ASCII art
      └─ Return HTML with ASCII art to client
```

## POST /export (File Export)

```
ExportAsciiArt(w, r)
  ├─ Validate HTTP method == POST
  │   └─ NO → Send400("Only POST method allowed")
  ├─ Extract form values:
  │   ├─ asciiResult := r.FormValue("ascii")
  │   └─ format := r.FormValue("format")
  ├─ Validate ASCII result not empty
  │   └─ NO → Send400("No ASCII art to export")
  ├─ Generate content based on format:
  │   ├─ format == "html":
  │   │   ├─ Wrap ASCII in HTML template
  │   │   ├─ contentType = "text/html; charset=utf-8"
  │   │   └─ filename = "ascii-art.html"
  │   ├─ format == "markdown":
  │   │   ├─ Wrap ASCII in markdown code block
  │   │   ├─ contentType = "text/markdown; charset=utf-8"
  │   │   └─ filename = "ascii-art.md"
  │   └─ default (txt):
  │       ├─ Use ASCII content as-is
  │       ├─ contentType = "text/plain; charset=utf-8"
  │       └─ filename = "ascii-art.txt"
  ├─ Set HTTP headers:
  │   ├─ Content-Type: format-specific MIME type
  │   ├─ Content-Length: file size in bytes
  │   └─ Content-Disposition: attachment with filename
  └─ Write file data to response
      └─ Browser initiates download
```

## GET /static/* (Static Assets)

```
FileServer Handler
  ├─ Strip "/static/" prefix from URL
  ├─ Serve files from static/ directory:
  │   ├─ style.css (7890 bytes)
  │   │   ├─ Glassmorphism styling
  │   │   ├─ Responsive design
  │   │   └─ Animation effects
  │   └─ fireworkEffect.js (11233 bytes)
  │       ├─ Canvas-based animations
  │       ├─ Particle physics
  │       └─ Interactive effects
  └─ Set appropriate Content-Type headers
```

## Error Handling Flow

```
Error Detected
  ├─ 400 Bad Request
  │   └─ Send400(w, message)
  │       ├─ Set HTTP status 400
  │       ├─ Parse templates/error.html
  │       └─ Render with ErrorData{400, message}
  ├─ 404 Not Found
  │   └─ Send404(w, message)
  │       ├─ Set HTTP status 404
  │       ├─ Parse templates/error.html
  │       └─ Render with ErrorData{404, message}
  └─ 500 Internal Server Error
      └─ Send500(w, message)
          ├─ Set HTTP status 500
          ├─ Parse templates/error.html
          └─ Render with ErrorData{500, message}

Error Page Features:
  ├─ Display status code and message
  ├─ "Go Home" button → /
  └─ Consistent styling with main page
```

## Banner File Management

```
EnsureFontFiles()
  └─ For each banner (standard, shadow, thinkertoy):
      ├─ bannerPath := "internal/utilities/ascii/banners/" + name + ".txt"
      ├─ FileExists(bannerPath)?
      │   └─ YES → Skip to next banner
      ├─ DownloadFile(zone01URL)
      │   ├─ HTTP GET to zone01.gr
      │   ├─ Read response body
      │   └─ Return []byte content
      └─ SaveFile(bannerPath, data)
          ├─ Create directories if needed
          └─ Write file with 0644 permissions

Banner URLs:
  ├─ standard: https://zone01.gr/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt
  ├─ shadow: https://zone01.gr/git/root/public/raw/branch/master/subjects/ascii-art/shadow.txt
  └─ thinkertoy: https://zone01.gr/git/root/public/raw/branch/master/subjects/ascii-art/thinkertoy.txt
```

## ASCII Generation Process

```
Input: "Hello\nWorld"
  ├─ AsciiFilter("Hello\nWorld")
  │   ├─ Check each character:
  │   │   ├─ 'H' (72) ∈ [32-126] ✓
  │   │   ├─ 'e' (101) ∈ [32-126] ✓
  │   │   ├─ '\n' (10) → preserve ✓
  │   │   └─ Continue for all chars...
  │   └─ Result: "Hello\nWorld" (no changes)
  ├─ SplitInputLines("Hello\nWorld")
  │   ├─ Normalize line endings
  │   └─ Result: ["Hello", "World"]
  └─ RenderAscii(["Hello", "World"], asciiLines)
      ├─ Process "Hello":
      │   ├─ PrintAsciiLine("Hello", asciiLines)
      │   │   ├─ For each char in "Hello":
      │   │   │   ├─ 'H' → index (72-32)*9 = 360
      │   │   │   ├─ 'e' → index (101-32)*9 = 621
      │   │   │   └─ Continue for 'l', 'l', 'o'...
      │   │   └─ Generate 8 rows of ASCII art
      │   └─ Add to result
      ├─ Add blank line separator
      └─ Process "World" (same process)
```

## Character Mapping Formula

```
Character Mapping:
  Input: character (e.g., 'A' = ASCII 65)
  ├─ charIndex = character - 32
  │   └─ 'A': 65 - 32 = 33
  ├─ startLine = charIndex * 9
  │   └─ 'A': 33 * 9 = 297
  └─ For rows 0-7:
      ├─ Row 0: asciiLines[297 + 0 + 1] = asciiLines[298]
      ├─ Row 1: asciiLines[297 + 1 + 1] = asciiLines[299]
      └─ Row 7: asciiLines[297 + 7 + 1] = asciiLines[305]

Banner File Structure (855 lines total):
  ├─ Character 32 (space): lines 1-9
  ├─ Character 33 (!): lines 10-18
  ├─ Character 34 ("): lines 19-27
  └─ Character 126 (~): lines 847-855

Each character block:
  ├─ Line 1: blank line
  └─ Lines 2-9: 8 rows of ASCII art
```

## Frontend Interactive Effects

```
Fireworks Animation System:
  ├─ Canvas Initialization
  │   ├─ gridCanvas (background grid)
  │   ├─ warpCanvas (warp effects)
  │   └─ fireworksCanvas (particle effects)
  ├─ Auto-Launch Timer
  │   └─ Every 1000ms: launch random firework
  ├─ User Interaction
  │   └─ Click anywhere: launch firework at cursor
  └─ Animation Loop
      ├─ Update particle physics
      ├─ Apply gravity and friction
      ├─ Remove expired particles
      └─ Render all effects

Particle Physics:
  ├─ Firework Launch
  │   ├─ Initial velocity: upward trajectory
  │   └─ Gravity: -0.1 acceleration
  └─ Explosion
      ├─ Create 50-100 particles
      ├─ Random velocities in all directions
      ├─ Color variations
      └─ Fade out over time
```

## Docker Build Process

```
Docker Build (Multi-stage):
  ├─ Stage 1: Builder (golang:1.25.7)
  │   ├─ WORKDIR /app
  │   ├─ COPY go.mod ./
  │   ├─ RUN go mod download
  │   ├─ COPY . .
  │   └─ RUN go build -o ascii-art ./cmd/main.go
  └─ Stage 2: Runtime (alpine:3.19)
      ├─ RUN apk add --no-cache ca-certificates
      ├─ WORKDIR /app
      ├─ COPY --from=builder /app/ascii-art .
      ├─ COPY --from=builder /app/templates ./templates
      ├─ COPY --from=builder /app/static ./static
      ├─ EXPOSE 8080
      └─ CMD ["./ascii-art"]

Result: ~34MB optimized image
```

## Testing Flow

```
Test Execution (make test):
  ├─ ASCII Logic Tests
  │   ├─ bannerHandler_test.go (2 tests)
  │   ├─ generator_test.go (2 tests)
  │   └─ helperFunctions_test.go (4 tests)
  └─ Server Tests
      ├─ handlers_test.go (3 tests)
      └─ errors_test.go (3 tests)

Total: 14 tests, all passing
Coverage: Available via 'make coverage'
```
