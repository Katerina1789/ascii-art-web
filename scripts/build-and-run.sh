#!/bin/bash

docker build -t ascii-art-web .
docker run -d -p 8080:8080 --name ascii ascii-art-web

# Clean unused Docker objects 
docker system prune -f