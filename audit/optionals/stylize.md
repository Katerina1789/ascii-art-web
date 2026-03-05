# Audit Guide: ASCII-Art-Web - Stylize

## Functional Tests

### Test 1: Color Visibility
**Check:** Do colors allow proper text visibility?

**How to Test:**
1. Navigate to `http://localhost:8080/`
2. Verify text is readable on white container
3. Check button text
4. Check ASCII output
5. Test error pages for readability

---

### Test 2: Design Consistency
**Check:** Does every page follow the same design?

**How to Test:**
1. Visit home page `/`
2. Generate ASCII art
3. Trigger error (visit `/invalid`)
4. Compare:
   - Color palette
   - Typography
   - Layout
   - Button styling

---

### Test 3: Responsive Design
**Check:** Does the site adapt to different screen sizes?

**How to Test:**
1. Open `http://localhost:8080/`
2. Resize browser window:
   - Desktop: 1920px wide
   - Tablet: 768px wide
   - Mobile: 375px wide
3. Verify:
   - Container adapts
   - Radio buttons stack on mobile
   - Canvas resizes
   - No broken layout

---

### Test 4: Interactive Design
**Check:** Does the site respond to user actions?

**How to Test:**
1. **Hover Effects:**
   - Hover over buttons
   - Hover over radio buttons

2. **Click Interactions:**
   - Click on page 
   - Click on the radio buttons
   - Type in textarea 
   - Submit form 

3. **Background Animation:**
   - Fireworks auto-launch
   - Rockets explode into particles
   - Grid boxes fill with colors

---

### Test 5: Project Standards
**Check:** Is the project complete and functional?

**How to Test:**
```bash
# Test compilation
go run ./cmd
# Expected: Server starts without errors

# Test stability
# - Generate ASCII 10+ times
# - Click fireworks 50+ times
# - Resize window multiple times
# - Trigger error pages
```

**Checklist:**
- [ ] Project compiles without errors
- [ ] All features implemented
- [ ] No crashes during use
- [ ] Proper error handling (400, 404, 500)
- [ ] Only standard Go packages used

---

## General Tests

### Test 6: Ease of Use
**Check:** Is the website easy to use?

**How to Test:**
1. Open `http://localhost:8080/` as first-time user
2. Verify:
   - Clear heading and labels
   - Obvious form fields
   - Single submit button
   - Result shows on same page
   - No instructions needed

---

### Test 7: Background Presence
**Check:** Does it have a background?

**How to Test:**
1. Open `http://localhost:8080/`
2. Verify animated fireworks background:
   - Black grid with gray lines
   - Rockets launching from bottom
   - Colorful particle explosions
   - Grid boxes filling with colors
   - Auto-launch every second

---

## Basic Tests

### Test 8: Performance
**Check:** Does the project run quickly and effectively?

**How to Test:**
```bash
# Server startup
time go run ./cmd
# Expected: < 2 seconds

# ASCII generation speed
time curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=standard"
# Expected: < 100ms response
```

---

### Test 9: Output Structure
**Check:** Is ASCII output well-structured and aligned?

**How to Test:**
1. Generate ASCII art: `Hello`
2. Verify:
   - Monospace font used
   - Characters align horizontally
   - 8 rows per character
   - No broken lines

3. Test edge cases:
   - Special characters: `{123}`
   - Newlines: `Hello\nWorld`

---

### Test 10: Code Quality
**Check:** Does code follow good practices?

**How to Test:**
```bash
# Check only standard library used
cat go.mod
# Expected: No external dependencies

# Check project structure
ls -R
# Expected: Organized folders (cmd/, internal/, static/, templates/)
```

**Code Review:**
- [ ] Clear package structure
- [ ] Separation of concerns (ASCII logic vs server logic)
- [ ] Error handling present
- [ ] Readable function names
- [ ] No code duplication

---

## Social

### Test 11: Learning Value
**Check:** Did you learn anything from this project?

**Discussion Points:**
- Canvas API and animations
- HTTP server in Go
- CSS responsive design
- Event handling
- Template rendering

---

### Test 12: Recommendation
**Check:** Would you recommend this as an example?

**Evaluation:**
- [ ] Exceeds basic requirements
- [ ] Well-structured code
- [ ] Polished user experience
- [ ] Good documentation
- [ ] Creative implementation

---

## Checklist for Auditors

### Functional
- [ ] Text is visible with good contrast
- [ ] Design is consistent across pages
- [ ] Layout is responsive (desktop/tablet/mobile)
- [ ] Interactive elements respond to user
- [ ] Project compiles and runs without errors

### General
- [ ] Website is easy to use
- [ ] Animated background is present

### Basic
- [ ] Performance is good (fast responses, smooth animations)
- [ ] ASCII output is properly aligned
- [ ] Code follows good practices
- [ ] Only standard Go packages used

### Social
- [ ] Project demonstrates learning
- [ ] Worthy of recommendation
