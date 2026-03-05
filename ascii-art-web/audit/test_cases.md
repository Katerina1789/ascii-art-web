# Test Cases for ASCII-Art-Web

## Functional Tests

### Test 1: Standard Banner - {123}
**Input:**
- Text: `{123}`
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

### Test 2: Standard Banner - <Hello> (World)!
**Input:**
- Text: `<Hello> (World)!`
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

---

### Test 3: Standard Banner - 123??
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

---

### Test 4: Shadow Banner - $% "=
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

---

### Test 5: Thinkertoy Banner - 123 T/fs#R
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

---

### Test 6: Newline Handling - Hello\nThere
**Input:**
- Text: `Hello\nThere` (or press Enter in textarea)
- Banner: `standard`

**Expected Output:**
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
```

---

## Error Handling Tests

### Test 7: 404 - Unknown Route
**Request:**
```bash
curl http://localhost:8080/nonexistent
```

**Expected Response:**
- Status Code: `404`
- Error Page with:
  - Title: "Error 404"
  - Message: "Page not found"
  - "Go Home" button

---

### Test 8: 400 - Invalid Banner
**Request:**
```bash
curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=invalid"
```

**Expected Response:**
- Status Code: `400`
- Error Page with:
  - Title: "Error 400"
  - Message: "Invalid banner name"
  - "Go Home" button

---

### Test 9: 400 - Wrong HTTP Method
**Request:**
```bash
curl -X GET http://localhost:8080/ascii-art
```

**Expected Response:**
- Status Code: `400`
- Error Page with:
  - Title: "Error 400"
  - Message: "Only POST method is allowed"
  - "Go Home" button

---

### Test 10: 404 - Template Not Found
**Scenario:** Delete templates/index.html temporarily

**Expected Response:**
- Status Code: `404`
- Error Page with:
  - Title: "Error 404"
  - Message: "Template not found"

---

## Static Files Tests

### Test 11: CSS Loading
**Request:**
```bash
curl http://localhost:8080/static/style.css
```

**Expected Response:**
- Status Code: `200`
- Content-Type: `text/css`
- CSS content returned

---

### Test 12: CSS in Browser
**How to Test:**
1. Open browser DevTools (F12)
2. Navigate to `http://localhost:8080/`
3. Check Network tab
4. Verify `/static/style.css` loads with status 200
5. Verify no 404 errors

---

## Banner Tests

### Test 13: All Three Banners Work
**Test each banner:**
- Standard: `Hello`
- Shadow: `Hello`
- Thinkertoy: `Hello`

**Expected:**
- All generate different ASCII art styles
- No errors
- Status: 200 OK

---

## Quick Test Commands

```bash
# Start server
go run cmd/main.go

# Test GET /
curl http://localhost:8080/

# Test POST with standard
curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=standard"

# Test 404
curl http://localhost:8080/invalid

# Test 400 - invalid banner
curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=wrong"

# Test 400 - wrong method
curl -X GET http://localhost:8080/ascii-art

# Test static files
curl http://localhost:8080/static/style.css
```

---

## Checklist for Auditors

- [ ] Project uses only standard Go packages
- [ ] HTML templates in `templates/` directory
- [ ] GET / returns main page
- [ ] POST /ascii-art processes form correctly
- [ ] All 3 banners work (standard, shadow, thinkertoy)
- [ ] 200 OK for successful requests
- [ ] 400 Bad Request for invalid input
- [ ] 404 Not Found for unknown routes
- [ ] 500 Internal Server Error handled
- [ ] Static files served correctly
- [ ] CSS loads without 404 errors
- [ ] Newlines handled correctly
- [ ] Test files exist in `testfiles/`
- [ ] Code follows good practices
- [ ] Server runs without crashes
