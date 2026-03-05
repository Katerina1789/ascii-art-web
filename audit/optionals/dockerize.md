# Audit Guide: ASCII-Art-Web - Dockerize

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

### Test 2: Dockerfile Presence
**Check:** Does the project have a Dockerfile?

**How to Test:**
```bash
# Check for Dockerfile
ls -la | grep Dockerfile
# Expected: Dockerfile exists in project root
```

---

### Test 3: Build Docker Image
**Check:** Can the Docker image be built successfully?

**How to Test:**
```bash
# Build the Docker image
docker image build -f Dockerfile -t ascii-art-web-docker .

# Verify image was created
docker images
```

**Expected Output:**
```
REPOSITORY               TAG        IMAGE ID       CREATED         SIZE
ascii-art-web-docker     latest     85a65d66ca39   7 seconds ago   795MB
```

**Checklist:**
- [ ] Build completes without errors
- [ ] Image appears in `docker images` list
- [ ] Image has reasonable size

---

### Test 4: Run Docker Container
**Check:** Can the container be started using the built image?

**How to Test:**
```bash
# Run the container
docker container run -p 8080:8080 --detach --name dockerize ascii-art-web-docker

# Verify container is running
docker ps -a
```

**Expected Output:**
```
CONTAINER ID   IMAGE                  COMMAND      CREATED        STATUS        PORTS                    NAMES
51c2efe2d366   ascii-art-web-docker   "./server"   6 seconds ago  Up 6 seconds  0.0.0.0:8080->8080/tcp   dockerize
```

**Checklist:**
- [ ] Container starts without errors
- [ ] Container status shows "Up"
- [ ] Port mapping is correct (8080:8080)

---

### Test 5: Container File System
**Check:** Does the container's file system contain expected files?

**How to Test:**
```bash
# Access container shell
docker exec -it dockerize /bin/bash

# List files
ls -l
```

**Expected Output:**
```
-rwxr-xr-x   1 root root 10833387 Sep  8 10:31 server
drwxr-xr-x   2 root root     4096 Sep  8 10:51 static
drwxr-xr-x   2 root root     4096 Sep  8 10:51 templates
```

**Checklist:**
- [ ] Server binary exists and is executable
- [ ] `static/` directory present
- [ ] `templates/` directory present
- [ ] Banner files accessible (if needed)

**Exit container:**
```bash
exit
```

---

### Test 6: Application Functionality
**Check:** Does the application work inside the container?

**How to Test:**
```bash
# Test the application
curl http://localhost:8080/

# Test ASCII generation
curl -X POST http://localhost:8080/ascii-art -d "text=Hello&banner=standard"

# Test in browser
# Open: http://localhost:8080/
```

**Checklist:**
- [ ] Home page loads successfully
- [ ] ASCII art generation works
- [ ] Static files (CSS) load correctly
- [ ] All three banners work

---

### Test 7: Docker Metadata
**Check:** Does the Dockerfile contain metadata?

**How to Test:**
```bash
# Inspect the Dockerfile
cat Dockerfile

# Check image metadata
docker inspect ascii-art-web-docker
```

**Expected Metadata:**
- [ ] LABEL with maintainer/author
- [ ] LABEL with description
- [ ] EXPOSE directive for port
- [ ] WORKDIR set appropriately
- [ ] CMD or ENTRYPOINT defined

---

### Test 8: Project Standards
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
# Test stability
# Generate ASCII 20+ times
# Check container logs
docker logs dockerize

# Check for errors
# Expected: No error messages
```

---

## General Tests

### Test 9: Build Script
**Check:** Does the project present a script to build images and containers?

**How to Test:**
```bash
# Look for build scripts
ls -la | grep -E "build|docker|script"
# Examples: build.sh, docker-build.sh, Makefile
```

**Checklist:**
- [ ] Script exists to simplify build process
- [ ] Script builds image
- [ ] Script runs container
- [ ] Script includes cleanup commands

**Example script commands:**
```bash
# Build
./build.sh

# Or if using Makefile
make build
make run
```

---

## Basic Tests

### Test 10: Performance
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

---

### Test 11: Code Quality
**Check:** Does the code obey good practices?

**How to Test:**
```bash
# Review Dockerfile
cat Dockerfile
```

**Dockerfile Best Practices:**
- [ ] Multi-stage build (if applicable)
- [ ] Minimal base image (alpine, scratch)
- [ ] Proper layer caching
- [ ] No unnecessary files in image
- [ ] Security considerations (non-root user)
- [ ] .dockerignore file present

---

### Test 12: Test Files
**Check:** Is there a test file for this code?

**How to Test:**
```bash
# Check for tests
ls testfiles/

# Run tests
go test ./testfiles/...
```

**Checklist:**
- [ ] Test files exist
- [ ] Tests cover ASCII logic
- [ ] Tests cover HTTP handlers
- [ ] All tests pass

---

### Test 13: Test Coverage
**Check:** Are the tests checking each possible case?

**How to Test:**
```bash
# Run tests with coverage
go test -cover ./testfiles/...
```

**Checklist:**
- [ ] Tests for all banners
- [ ] Tests for error cases
- [ ] Tests for edge cases
- [ ] Good coverage percentage

---

### Test 14: Instructions Clarity
**Check:** Are the instructions in the website clear?

**How to Test:**
1. Open `http://localhost:8080/`
2. Verify:
   - Clear labels
   - Obvious form fields
   - Banner selection is clear
   - Submit button is visible
   - Results display clearly

---

### Test 15: API Usage
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
- [ ] Correct status codes (200, 400, 404, 500)
- [ ] JSON or form data handling

---

## Social

### Test 16: Learning Value
**Check:** Did you learn anything from this project?

**Discussion Points:**
- Docker containerization
- Multi-stage builds
- Container orchestration
- Image optimization
- Deployment strategies

---

### Test 17: Open Source Potential
**Check:** Can it be open-sourced / used for other sources?

**Evaluation:**
- [ ] Well-documented
- [ ] Clear README
- [ ] Reusable Dockerfile
- [ ] Good project structure
- [ ] License included

---

### Test 18: Recommendation
**Check:** Would you recommend this as an example?

**Evaluation:**
- [ ] Follows Docker best practices
- [ ] Clean implementation
- [ ] Good documentation
- [ ] Easy to deploy
- [ ] Production-ready

---

## Cleanup Commands

```bash
# Stop container
docker stop dockerize

# Remove container
docker rm dockerize

# Remove image
docker rmi ascii-art-web-docker

# Clean up all
docker stop dockerize && docker rm dockerize && docker rmi ascii-art-web-docker
```

---

## Checklist for Auditors

### Functional
- [ ] Only standard Go packages used
- [ ] Dockerfile exists
- [ ] Docker image builds successfully
- [ ] Container runs successfully
- [ ] File system contains expected files
- [ ] Dockerfile contains metadata
- [ ] Project meets all standards

### General
- [ ] Build script provided for simplification

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
