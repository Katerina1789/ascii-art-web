# Audit Guide: ASCII-Art-Web - Export File

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

### Test 2: File Export Functionality
**Check:** Can the file be exported successfully?

**How to Test:**
1. Navigate to `http://localhost:8080/`
2. Enter text: `Hello World`
3. Select banner: `standard`
4. Click "Generate ASCII Art"
5. Click export/download button
6. Verify file downloads

**Checklist:**
- [ ] Export button/link is visible
- [ ] File downloads successfully
- [ ] File has appropriate name
- [ ] Download completes without errors

---

### Test 3: Output Matching
**Check:** Does the exported file match the output?

**How to Test:**
```bash
# Generate ASCII art in browser
# Download the file
# Open the downloaded file

# Compare with browser output
cat ~/Downloads/ascii-art.txt
```

**Checklist:**
- [ ] Content matches exactly
- [ ] No extra characters added
- [ ] No missing characters
- [ ] Line breaks preserved
- [ ] Formatting intact

---

### Test 4: File Permissions
**Check:** Are the exported files read and write for the user?

**How to Test:**
```bash
# Check file permissions
ls -l ~/Downloads/ascii-art.txt

# Try to read the file
cat ~/Downloads/ascii-art.txt

# Try to edit the file
echo "# Test" >> ~/Downloads/ascii-art.txt
```

**Expected Permissions:**
```
-rw-r--r--  1 user user 1234 Jan 01 12:00 ascii-art.txt
```

**Checklist:**
- [ ] File is readable (r--)
- [ ] File is writable (rw-)
- [ ] User can open and edit file
- [ ] No permission errors

---

### Test 5: Content-Type Header

**How to Test:**
```bash
curl -v POST http://localhost:8080/export -d "ascii=Hello" --output /dev/null
```

**Expected Header:**
```
Content-Type: text/plain
```

**Checklist:**
- [ ] Content-Type header is present
- [ ] Media type is appropriate
- [ ] Charset optional (not required for audit)

---

### Test 6: Content-Length Header

**How to Test:**
```bash
curl -v POST http://localhost:8080/export -d "ascii=Hello" --output /dev/null
```

**Expected Header:**
```
Content-Length: 5
```

**Checklist:**
- [ ] Content-Length header is present
- [ ] Size matches actual file size
- [ ] Size is in bytes

---

### Test 7: Content-Disposition Header

**How to Test:**
```bash
curl -v POST http://localhost:8080/export -d "ascii=Hello" --output /dev/null
```

**Expected Header:**
```
Content-Disposition: attachment; filename="ascii-art.txt"
```

**Checklist:**
- [ ] Content-Disposition header is present
- [ ] Uses "attachment"
- [ ] Filename is specified
- [ ] Filename is descriptive

---

### Test 8: Download Button/Link
**Check:** Does the site have a clear button/link to download/export?

**How to Test:**
1. Navigate to `http://localhost:8080/`
2. Generate ASCII art
3. Look for export/download button

**Checklist:**
- [ ] Button/link is clearly visible
- [ ] Label is clear ("Download", "Export", "Save")
- [ ] Button is easy to find
- [ ] Button appears after generation
- [ ] Button styling is consistent

---

### Test 9: Project Standards
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

# Test stability
# - Export 20+ files
# - Try different text inputs
# - Test all banners
# - Check for memory leaks
```

---

## Basic Tests

### Test 10: Code Quality
**Check:** Does the code obey good practices?

**How to Test:**
```bash
# Review export handler code
cat internal/server/handlers.go
```

**Code Review Checklist:**
- [ ] Clear function names
- [ ] Proper error handling
- [ ] HTTP headers set correctly
- [ ] No code duplication
- [ ] Follows Go conventions
- [ ] Comments where needed

---

### Test 11: Instructions Clarity
**Check:** Are the instructions in the website clear?

**How to Test:**
1. Open `http://localhost:8080/`
2. Verify:
   - Clear labels for all fields
   - Export button is obvious
   - No confusion about how to download
   - Instructions (if any) are clear

**Checklist:**
- [ ] User knows how to export
- [ ] Button label is clear
- [ ] No ambiguity in UI
- [ ] Tooltips/help text (if needed)

---

### Test 12: API Usage
**Check:** Does the project run using an API?

**How to Test:**
```bash
# Test API endpoints
curl http://localhost:8080/
curl -X POST http://localhost:8080/ascii-art -d "text=Test&banner=standard"
curl http://localhost:8080/export
```

**Checklist:**
- [ ] RESTful endpoints
- [ ] Proper HTTP methods
- [ ] Correct status codes
- [ ] API for export functionality

---

### Test 13: Multiple Export Formats
**Check:** Can you export in multiple formats?

**How to Test:**
1. Generate ASCII art
2. Check available export options:
   - `.txt` (plain text)
   - `.html` (HTML format)
   - `.md` (Markdown format)
   - Other formats

**Checklist:**
- [ ] Multiple format options available
- [ ] Each format works correctly
- [ ] Format selection is clear
- [ ] Content-Type matches format

---

## Advanced Tests

### Test 14: File Naming
**Check:** Are exported files named appropriately?

**How to Test:**
1. Export multiple files
2. Check filenames

**Good Practices:**
- [ ] Descriptive names
- [ ] Includes timestamp (optional)
- [ ] Includes banner name (optional)
- [ ] Valid filename characters
- [ ] Proper file extension

**Examples:**
- `ascii-art.txt`
- `ascii-art-standard-2024-01-15.txt`
- `hello-world-shadow.txt`

---

### Test 15: Special Characters in Export
**Check:** Does export handle special characters correctly?

**How to Test:**
1. Generate ASCII with special chars: `{123} <Hello>`
2. Export file
3. Open and verify content

**Checklist:**
- [ ] Special characters preserved
- [ ] No encoding issues
- [ ] UTF-8 encoding (if applicable)
- [ ] Line breaks correct

---

### Test 16: Large Output Export
**Check:** Can large ASCII outputs be exported?

**How to Test:**
1. Generate large ASCII art (multiple lines, long text)
2. Export file
3. Verify complete content

**Checklist:**
- [ ] No truncation
- [ ] Complete content exported
- [ ] File size is correct
- [ ] No performance issues

---

### Test 17: Error Handling
**Check:** Does export handle errors gracefully?

**How to Test:**
```bash
# Try to export without generating ASCII
curl http://localhost:8080/export

# Try invalid export requests
curl -X POST http://localhost:8080/export
```

**Checklist:**
- [ ] Appropriate error messages
- [ ] Correct status codes
- [ ] No crashes
- [ ] User-friendly error pages

---

## Social

### Test 18: Learning Value
**Check:** Did you learn anything from this project?

**Discussion Points:**
- HTTP headers for file downloads
- Content-Disposition usage
- File handling in Go
- Browser download mechanisms

---

### Test 19: Open Source Potential
**Check:** Can it be open-sourced / used for other sources?

**Evaluation:**
- [ ] Well-documented
- [ ] Reusable export logic
- [ ] Clear code structure
- [ ] Good examples for learning

---

### Test 20: Recommendation
**Check:** Would you recommend this as an example?

**Evaluation:**
- [ ] Proper HTTP headers implementation
- [ ] Clean export functionality
- [ ] Good user experience
- [ ] Follows web standards
- [ ] Production-ready

---

## Checklist for Auditors

### Functional
- [ ] Only standard Go packages used
- [ ] File export works successfully
- [ ] Exported file matches output
- [ ] File permissions are correct (read/write)
- [ ] Content-Type header is present
- [ ] Content-Length header is present
- [ ] Content-Disposition header is present
- [ ] Clear download button/link exists
- [ ] Project meets all standards

### Basic
- [ ] Code follows good practices
- [ ] Website instructions are clear
- [ ] Project uses API endpoints
- [ ] Multiple export formats available (bonus)

### Social
- [ ] Project demonstrates learning
- [ ] Can be open-sourced
- [ ] Worthy of recommendation
