# Project Flow

## Server Startup

```
main()
  ├─ EnsureFontFiles()
  │   └─ Download missing banners from zone01.gr
  ├─ Register routes:
  │   ├─ GET  /          → HandleHome
  │   ├─ POST /ascii-art → HandleAsciiArt
  │   └─ GET  /static/*  → FileServer
  └─ Start server on :8080
```

## GET / (Home Page)

```
HandleHome(w, r)
  ├─ Validate path == "/"
  │   └─ NO  → Send404("Page not found")
  ├─ Parse templates/index.html
  │   └─ ERROR → Send404("Template not found")
  └─ Execute template with nil data
      └─ Return HTML to client
```

## POST /ascii-art (Generate ASCII)

```
HandleAsciiArt(w, r)
  ├─ Validate method == POST
  │   └─ NO → Send400("Only POST method allowed")
  ├─ Extract form values (text, banner)
  ├─ Validate banner ∈ {standard, shadow, thinkertoy}
  │   └─ NO → Send400("Invalid banner name")
  ├─ GenerateAsciiArt(text, banner)
  │   ├─ LoadAsciiFile(bannerPath)
  │   │   └─ Read 855 lines from banner file
  │   ├─ AsciiFilter(text)
  │   │   └─ Remove chars < 32 or > 126 (keep \n, \r)
  │   ├─ SplitInputLines(filteredText)
  │   │   └─ Split by \n (normalize \r\n, \\n)
  │   └─ RenderAscii(lines, asciiLines)
  │       └─ For each line:
  │           └─ PrintAsciiLine(line, asciiLines)
  │               └─ Generate 8 rows per character
  ├─ Check for errors
  │   └─ ERROR → Send500("Failed to generate ASCII art")
  ├─ Parse templates/index.html
  │   └─ ERROR → Send404("Template not found")
  └─ Execute template with result data
      └─ Return HTML with ASCII art to client
```

## Error Handling

```
Error Detected
  ├─ 400 Bad Request
  │   └─ Send400(w, message)
  │       ├─ Set status 400
  │       └─ Render error.html
  ├─ 404 Not Found
  │   └─ Send404(w, message)
  │       ├─ Set status 404
  │       └─ Render error.html
  └─ 500 Internal Server Error
      └─ Send500(w, message)
          ├─ Set status 500
          └─ Render error.html
```

## Banner File Management

```
EnsureFontFiles()
  └─ For each banner (standard, shadow, thinkertoy):
      ├─ FileExists(path)?
      │   └─ YES → Skip
      ├─ DownloadFile(url)
      │   ├─ HTTP GET to zone01.gr
      │   └─ Read response body
      └─ SaveFile(path, data)
          └─ Write to internal/utilities/ascii/banners/
```

## ASCII Generation Process

```
Input: "Hello\nWorld"
  ├─ AsciiFilter(text)
  │   └─ Keep ASCII 32-126 + \n, \r
  ├─ SplitInputLines(filtered)
  │   └─ Result: ["Hello", "World"]
  └─ RenderAscii(lines, asciiLines)
      └─ For each line:
          ├─ PrintAsciiLine("Hello", asciiLines)
          │   └─ 8 rows of ASCII art
          ├─ Add blank line separator
          └─ PrintAsciiLine("World", asciiLines)
              └─ 8 rows of ASCII art
```

## Character Mapping Formula

```
Character 'A' (ASCII 65):
  charIndex = 65 - 32 = 33
  start = 33 × 9 = 297
  
  For row 0-7:
    asciiLines[(297 + row) + 1]
  
  Example row 0: asciiLines[298]
  Example row 7: asciiLines[305]
```
