#!/bin/bash

# Docker Version Updater Script
# This script updates the version label in a Dockerfile and rebuilds the image

set -e  # Exit on error

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Default Dockerfile path
DOCKERFILE="${1:-Dockerfile}"

# Check if Dockerfile exists
if [ ! -f "$DOCKERFILE" ]; then
    echo -e "${RED}Error: Dockerfile not found at: $DOCKERFILE${NC}"
    echo "Usage: $0 [path/to/Dockerfile]"
    exit 1
fi

# Extract current version from Dockerfile
CURRENT_VERSION=$(grep -i "^LABEL version=" "$DOCKERFILE" | cut -d'"' -f2 | cut -d"'" -f2 | head -n1)

if [ -z "$CURRENT_VERSION" ]; then
    echo -e "${YELLOW}Warning: No version label found in Dockerfile${NC}"
    echo "Please enter the current version (e.g., 1.0.0):"
    read CURRENT_VERSION
    
    if [ -z "$CURRENT_VERSION" ]; then
        echo -e "${RED}Error: Version cannot be empty${NC}"
        exit 1
    fi
fi

echo -e "${GREEN}Current version: $CURRENT_VERSION${NC}"

# Parse version components
if [[ $CURRENT_VERSION =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
    MAJOR="${BASH_REMATCH[1]}"
    MINOR="${BASH_REMATCH[2]}"
    PATCH="${BASH_REMATCH[3]}"
else
    echo -e "${RED}Error: Invalid version format. Expected format: X.Y.Z (e.g., 1.0.0)${NC}"
    exit 1
fi

# Ask user for update type
echo ""
echo "What type of update do you want to perform?"
echo "1) Major (X.0.0) - Breaking changes"
echo "2) Minor (x.Y.0) - New features, backwards compatible"
echo "3) Patch (x.y.Z) - Bug fixes"
read -p "Enter choice (1-3): " UPDATE_TYPE

# Calculate new version
case $UPDATE_TYPE in
    1)
        NEW_VERSION="$((MAJOR + 1)).0.0"
        UPDATE_NAME="Major"
        ;;
    2)
        NEW_VERSION="$MAJOR.$((MINOR + 1)).0"
        UPDATE_NAME="Minor"
        ;;
    3)
        NEW_VERSION="$MAJOR.$MINOR.$((PATCH + 1))"
        UPDATE_NAME="Patch"
        ;;
    *)
        echo -e "${RED}Error: Invalid choice${NC}"
        exit 1
        ;;
esac

echo ""
echo -e "${YELLOW}$UPDATE_NAME update: $CURRENT_VERSION → $NEW_VERSION${NC}"
read -p "Proceed with this update? (y/n): " CONFIRM

if [[ ! $CONFIRM =~ ^[Yy]$ ]]; then
    echo "Update cancelled."
    exit 0
fi

# Update Dockerfile
echo ""
echo "Updating Dockerfile..."

# 1. Σβήνουμε το παλιό LABEL version (όπου κι αν είναι) για να μην έχουμε σκουπίδια
sed -i.bak '/^LABEL version=/d' "$DOCKERFILE"

# 2. Βάζουμε το νέο LABEL πάντα στο τέλος (για μέγιστη ταχύτητα cache)
# Ψάχνουμε αν υπάρχει CMD για να το βάλουμε ακριβώς από πάνω
if grep -q "^CMD" "$DOCKERFILE"; then
     # Βάλτο ΠΡΙΝ το CMD
     # (Η διπλή εντολή είναι για συμβατότητα Linux/Mac)
     sed -i '' "/^CMD/i \\
LABEL version=\"$NEW_VERSION\"
" "$DOCKERFILE" 2>/dev/null || sed -i "/^CMD/i LABEL version=\"$NEW_VERSION\"" "$DOCKERFILE"
else
     # Αν δεν βρει CMD, απλά πρόσθεσέ το στο τέλος του αρχείου
     echo "LABEL version=\"$NEW_VERSION\"" >> "$DOCKERFILE"
fi

echo -e "${GREEN}✓ Dockerfile updated (Label placed at the bottom)${NC}"

# Extract image name from Dockerfile or ask user
IMAGE_NAME=$(grep -i "^# Image:" "$DOCKERFILE" | cut -d: -f2 | xargs)

if [ -z "$IMAGE_NAME" ]; then
    read -p "Enter Docker image name (e.g., myapp): " IMAGE_NAME
    if [ -z "$IMAGE_NAME" ]; then
        IMAGE_NAME="myimage"
    fi
fi

# Build Docker image
echo ""
echo "Building Docker image..."
echo -e "${YELLOW}Image: $IMAGE_NAME:$NEW_VERSION${NC}"
echo ""

# Get current date for BUILD_DATE argument
BUILD_DATE=$(date +%Y-%m-%d)
echo "Build date: $BUILD_DATE"

# Ask for maintainer name
read -p "Enter maintainer name: " MAINTAINER
if [ -z "$MAINTAINER" ]; then
    MAINTAINER="Unknown"
fi
echo ""

if docker build --build-arg BUILD_DATE=$BUILD_DATE --build-arg MAINTAINER="$MAINTAINER" -t "$IMAGE_NAME:$NEW_VERSION" -t "$IMAGE_NAME:latest" -f "$DOCKERFILE" .; then
    echo ""
    echo -e "${GREEN}✅ Build successful!${NC}"
    echo ""
    echo "Tagged as:"
    echo "  - $IMAGE_NAME:$NEW_VERSION"
    echo "  - $IMAGE_NAME:latest"
    echo ""
    echo "Build date: $BUILD_DATE"
    echo "Maintainer: $MAINTAINER"
else
    echo ""
    echo -e "${RED}❌ Build failed.${NC}"
    echo "Restoring original Dockerfile..."
    mv "${DOCKERFILE}.bak" "$DOCKERFILE"
    exit 1
fi

# Clean up backup on success
rm -f "${DOCKERFILE}.bak"

echo ""
echo "Backup removed: ${DOCKERFILE}.bak"
echo -e "${GREEN}Done!${NC}"
