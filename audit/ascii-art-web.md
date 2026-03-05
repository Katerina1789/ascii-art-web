# Audit Guide: ASCII-Art-Web

## Functional Tests

### Test 1: Package Requirements
**Check:** Has the requirement for allowed packages been respected?

**How to Test:**
```bash
# Check go.mod for dependencies
cat go.mod
# Expected: Only standard Go packages (no external dependencies)
```

**Checklist:**
- [ ] Only standard library packages used
- [ ] No third-party dependencies in go.mod

---

### Test 2: HTML Files Presence
**Check:** Does the project contain HTML files?

**How to Test:**
```bash
# Check for HTML templates
ls templates/
# Expected: index.html, error.html (or similar)
```

**Checklist:**
- [ ] HTML files exist in `templates/` directory
- [ ] Main page template present
- [ ] Error page template present

---

### Test 3: Standard Banner - {123}
**Check:** Does it display the right result?

**Input:**
- Line 1: `{123}`
- Banner: `standard`

**Expected Output:**
```
   __                     __    
  / /  _   ____    _____  \ \   
 | |  / | |___ \  |___ /   | |  
/ /   | |   __) |   |_ \    \ \ 
\ \   | |  / __/   ___) |   / / 
 | |  |_| |_____| |____/   | |  
  \_\                     /_/   
                                
```

**How to Test:**
1. Navigate to `http://localhost:8080/`
2. Enter `{123}` in text field
3. Select `standard` banner
4. Click "Generate ASCII Art"
5. Verify output matches expected

---

### Test 4: Standard Banner - <Hello> (World)!
**Check:** Does it display the right result?

**Input:**
- Line 2: `<Hello> (World)!`
- Banner: `standard`

**Expected Output:**
```
   __  _    _          _   _          __            __ __          __                 _       _  __    _
  / / | |  | |        | | | |         \ \          / / \ \        / /                | |     | | \ \  | |
 / /  | |__| |   ___  | | | |   ___    \ \        | |   \ \  /\  / /    ___    _ __  | |   __| |  | | | |
< <   |  __  |  / _ \ | | | |  / _ \    > >       | |    \ \/  \/ /    / _ \  | '__| | |  / _` |  | | | |
 \ \  | |  | | |  __/ | | | | | (_) |  / /        | |     \  /\  /    | (_) | | |    | | | (_| |  | | |_|
  \_\ |_|  |_|  \___| |_| |_|  \___/  /_/         | |      \/  \/      \___/  |_|    |_|  \__,_|  | | (_)
                                                   \_\                                           /_/       
                                                                                                           
```

**How to Test:**
1. Navigate to `http://localhost:8080/`
2. Enter `<Hello> (World)!` in text field
3. Select `standard` banner
4. Click "Generate ASCII Art"
5. Verify output matches expected

---

### Test 5: Standard Banner - 123??
**Check:** Does it display the right result?

**Input:**
- Text: `123??`
- Banner: `standard`

**Expected Output:**
```
                     ___    ___  
 _   ____    _____  |__ \  |__ \ 
/ | |___ \  |___ /     ) |    ) |
| |   __) |   |_ \    / /    / / 
| |  / __/   ___) |  |_|    |_|  
|_| |_____| |____/   (_)    (_)  
                                 
                                 
```

**How to Test:**
1. Navigate to `http://localhost:8080/`
2. Enter `123??` in text field
3. Select `standard` banner
4. Click "Generate ASCII Art"
5. Verify output matches expected

---

### Test 6: Shadow Banner - $% "=
**Check:** Does it display the right result?

**Input:**
- Text: `$% "=`
- Banner: `shadow`

**Expected Output:**
```
                        _|  _|  
  _|   _|_|    _|       _|  _|  
_|_|_| _|_|  _|                _|_|_|_|_|
_|_|       _|                            
  _|_|   _|  _|_|              _|_|_|_|_|
_|_|_| _|    _|_|                        
  _|                                     
```

**How to Test:**
1. Navigate to `http://localhost:8080/`
2. Enter `$% "=` in text field
3. Select `shadow` banner
4. Click "Generate ASCII Art"
5. Verify output matches expected

---

### Test 7: Thinkertoy Banner - 123 T/fs#R
**Check:** Does it display the right result?

**Input:**
- Text: `123 T/fs#R`
- Banner: `thinkertoy`

**Expected Output:**
```
  0    --  o-o        o-O-o     o  o-o      | |  o--o
 /|   o  o    |         |      /   |       -O-O- |   |
o |     /   oo          |     o   -O-  o-o  | |  O-Oo 
  |    /      |         |    /     |    \  -O-O- |  \ 
o-o-o o--o o-o          o   o      o   o-o  | |  o   o
```

**How to Test:**
1. Navigate to `http://localhost:8080/`
2. Enter `123 T/fs#R` in text field
3. Select `thinkertoy` banner
4. Click "Generate ASCII Art"
5. Verify output matches expected

---

### Test 8: Graphical Representation
**Check:** Does it display an understandable graphical representation?

**How to Test:**
1. Generate ASCII art with any input
2. Verify:
   - Output is displayed clearly
   - Monospace font is used
   - Characters are aligned properly
   - Result is readable
   - No broken formatting

**Checklist:**
- [ ] ASCII art is clearly visible
- [ ] Proper font rendering
- [ ] Correct alignment
- [ ] No visual glitches

---

### Test 9: Page Navigation
**Check:** Can you navigate between all available pages?

**How to Test:**
```bash
# Test main page
curl http://localhost:8080/

# Test ASCII generation
curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=standard"

# Test static files
curl http://localhost:8080/static/style.css

# Test unknown route
curl http://localhost:8080/nonexistent
```

**Checklist:**
- [ ] Main page loads (GET /)
- [ ] ASCII generation works (POST /ascii-art)
- [ ] Static files accessible (GET /static/*)
- [ ] Error pages display correctly

---

### Test 10: 404 Status Implementation
**Check:** Does the project implement 404 status?

**How to Test:**
```bash
# Test unknown route
curl -I http://localhost:8080/nonexistent

# Expected: HTTP/1.1 404 Not Found
```

**Checklist:**
- [ ] Returns 404 status code
- [ ] Displays error page
- [ ] Error message is clear
- [ ] "Go Home" link/button present

---

### Test 11: 400 Status - Bad Request
**Check:** Does the project handle HTTP status 400?

**How to Test:**
```bash
# Test invalid banner
curl -I -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=invalid"

# Test wrong HTTP method
curl -I -X GET http://localhost:8080/ascii-art

# Expected: HTTP/1.1 400 Bad Request
```

**Checklist:**
- [ ] Returns 400 for invalid banner
- [ ] Returns 400 for wrong HTTP method
- [ ] Displays error page
- [ ] Error message explains the issue

---

### Test 12: 500 Status - Internal Server Error
**Check:** Does the project handle HTTP status 500?

**How to Test:**
```bash
# Simulate server error (e.g., missing banner file)
# Or test with malformed data
```

**Checklist:**
- [ ] Returns 500 for server errors
- [ ] Displays error page
- [ ] Doesn't expose internal details
- [ ] Graceful error handling

---

### Test 13: Server-Client Communication
**Check:** Is communication between server and client well established?

**How to Test:**
1. Open browser DevTools (F12)
2. Navigate to `http://localhost:8080/`
3. Generate ASCII art
4. Check Network tab:
   - Request method
   - Response status
   - Response time
   - Headers

**Checklist:**
- [ ] Form submission works
- [ ] Response is received
- [ ] Data is displayed correctly
- [ ] No network errors

---

### Test 14: HTTP Method Validation
**Check:** Does the server use the right HTTP method?

**How to Test:**
```bash
# GET for main page
curl -X GET http://localhost:8080/
# Expected: 200 OK

# POST for ASCII generation
curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=standard"
# Expected: 200 OK

# GET for static files
curl -X GET http://localhost:8080/static/style.css
# Expected: 200 OK
```

**Checklist:**
- [ ] GET / returns main page
- [ ] POST /ascii-art generates ASCII
- [ ] GET /static/* serves files
- [ ] Wrong methods return 400

---

### Test 15: Stability
**Check:** Did the site work without crashing at any time?

**How to Test:**
```bash
# Run multiple requests
for i in {1..50}; do
  curl -X POST http://localhost:8080/ascii-art -d "text=Test$i&banner=standard"
done

# Check server logs for errors
```

**Checklist:**
- [ ] No crashes during testing
- [ ] Handles multiple requests
- [ ] No memory leaks
- [ ] Consistent performance

---

### Test 16: Server Language
**Check:** Is the server written in Go?

**How to Test:**
```bash
# Check for Go files
ls cmd/*.go internal/**/*.go

# Check go.mod
cat go.mod

# Run with Go
go run ./cmd
```

**Checklist:**
- [ ] Server is written in Go
- [ ] go.mod file present
- [ ] Compiles with `go run`
- [ ] Uses Go standard library

---

### Test 17: Project Standards
**Check:** Is this project up to every standard?

**Evaluation Criteria:**
- [ ] **Empty Work:** Project is complete and functional
- [ ] **Incomplete Work:** All features implemented
- [ ] **Invalid Compilation:** Builds successfully
- [ ] **Cheating:** Original work
- [ ] **Crashing:** Runs without crashes
- [ ] **Leaks:** No resource leaks

**How to Test:**
```bash
# Test compilation
go run ./cmd

# Test all features
# - Generate ASCII with all banners
# - Test error handling
# - Test static files
# - Check for crashes
```

---

## General Tests

### Test 18: HTTP Handlers and Patterns
**Check:** Does the server present all needed handlers and patterns?

**How to Test:**
```bash
# Review server code
cat internal/server/handlers.go
```

**Expected Handlers:**
- [ ] GET / - Main page handler
- [ ] POST /ascii-art - ASCII generation handler
- [ ] GET /static/* - Static file handler
- [ ] 404 handler for unknown routes
- [ ] Error handlers (400, 404, 500)

---

## Basic Tests

### Test 19: Performance
**Check:** Does the server run quickly and effectively?

**How to Test:**
```bash
# Test response time
time curl http://localhost:8080/

# Test ASCII generation speed
time curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=standard"

# Expected: < 100ms response time
```

**Checklist:**
- [ ] Fast startup time
- [ ] Quick response times
- [ ] No unnecessary data requests
- [ ] Efficient resource usage
- [ ] No recursive issues

---

### Test 20: Code Quality
**Check:** Does the code obey good practices?

**How to Test:**
```bash
# Review code structure
ls -R

# Check for organization
cat internal/utilities/ascii/generator.go
cat internal/server/handlers.go
```

**Code Review Checklist:**
- [ ] Clear package structure
- [ ] Separation of concerns
- [ ] Proper error handling
- [ ] Readable function names
- [ ] No code duplication
- [ ] Comments where needed
- [ ] Follows Go conventions

---

### Test 21: Test Files
**Check:** Is there a test file for this code?

**How to Test:**
```bash
# Check for test files
ls testfiles/

# Run tests
go test ./testfiles/...
```

**Checklist:**
- [ ] Test files exist in `testfiles/`
- [ ] Tests for ASCII logic
- [ ] Tests for HTTP handlers
- [ ] All tests pass

---

### Test 22: Test Coverage
**Check:** Are the tests checking each possible case?

**How to Test:**
```bash
# Run tests with coverage
go test -cover ./testfiles/...
```

**Test Cases to Check:**
- [ ] All three banners tested
- [ ] Special characters tested
- [ ] Newline handling tested
- [ ] Error cases tested (400, 404, 500)
- [ ] HTTP method validation tested
- [ ] Edge cases covered

---

### Test 23: Instructions Clarity
**Check:** Are the instructions in the website clear?

**How to Test:**
1. Open `http://localhost:8080/`
2. Verify:
   - Clear heading
   - Labeled text input
   - Clear banner selection
   - Obvious submit button
   - Result display is clear

**Checklist:**
- [ ] User knows what to do
- [ ] Labels are descriptive
- [ ] No confusion about inputs
- [ ] Instructions (if any) are clear

---

### Test 24: API Usage
**Check:** Does the project run using an API?

**How to Test:**
```bash
# Test API endpoints
curl http://localhost:8080/
curl -X POST http://localhost:8080/ascii-art -d "text=Test&banner=standard"
curl http://localhost:8080/static/style.css
```

**Checklist:**
- [ ] RESTful endpoints
- [ ] Proper HTTP methods (GET, POST)
- [ ] Correct status codes
- [ ] Form data handling
- [ ] Response format is consistent

---

## Social

### Test 25: Learning Value
**Check:** Did you learn anything from this project?

**Discussion Points:**
- HTTP server implementation in Go
- Template rendering
- Error handling
- ASCII art generation algorithms
- Web form processing

---

### Test 26: Open Source Potential
**Check:** Can it be open-sourced / used for other sources?

**Evaluation:**
- [ ] Well-documented
- [ ] Clear code structure
- [ ] Reusable components
- [ ] Good examples for learning
- [ ] License included

---

### Test 27: Recommendation
**Check:** Would you recommend this as an example?

**Evaluation:**
- [ ] Meets all requirements
- [ ] Clean implementation
- [ ] Good documentation
- [ ] Follows best practices
- [ ] Production-ready

---

## Quick Test Commands

```bash
# Start server
go run ./cmd

# Test GET /
curl http://localhost:8080/

# Test POST with standard
curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=standard"

# Test all banners
curl -X POST http://localhost:8080/ascii-art -d "text=Test&banner=standard"
curl -X POST http://localhost:8080/ascii-art -d "text=Test&banner=shadow"
curl -X POST http://localhost:8080/ascii-art -d "text=Test&banner=thinkertoy"

# Test 404
curl -I http://localhost:8080/invalid

# Test 400 - invalid banner
curl -I -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=wrong"

# Test 400 - wrong method
curl -I -X GET http://localhost:8080/ascii-art

# Test static files
curl http://localhost:8080/static/style.css

# Run tests
go test ./testfiles/...
```

---

## Checklist for Auditors

### Functional
- [ ] Only standard Go packages used
- [ ] HTML files present in templates/
- [ ] Standard banner: {123} displays correctly
- [ ] Standard banner: <Hello> (World)! displays correctly
- [ ] Standard banner: 123?? displays correctly
- [ ] Shadow banner: $% "= displays correctly
- [ ] Thinkertoy banner: 123 T/fs#R displays correctly
- [ ] Graphical representation is understandable
- [ ] All pages are accessible
- [ ] 404 status implemented
- [ ] 400 status handled
- [ ] 500 status handled
- [ ] Server-client communication works
- [ ] Correct HTTP methods used
- [ ] Site runs without crashing
- [ ] Server is written in Go
- [ ] Project meets all standards

### General
- [ ] All HTTP handlers and patterns present

### Basic
- [ ] Server runs quickly and effectively
- [ ] Code follows good practices
- [ ] Test files exist
- [ ] Tests check all cases
- [ ] Website instructions are clear
- [ ] Project uses API endpoints

### Social
- [ ] Project demonstrates learning
- [ ] Can be open-sourced
- [ ] Worthy of recommendation
