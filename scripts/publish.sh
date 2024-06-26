#!/bin/sh

set -e

# Get the last version in repository
CURRENT_VERSION=$(git describe --abbrev=0 --tags)

# Get the declared version in .version file
VERSION=$(cat .version)

# Check if versions are equal
if [ "$CURRENT_VERSION" = "$VERSION" ]; then
    echo "Versions are equal. Nothing to publish"
    exit 0
fi

# Tidy the repository
go mod tidy

# Create the new tag version
git tag "$VERSION"

# Push the new tag version
git push origin "$VERSION"

# Publish the new version
GOPROXY=proxy.golang.org go list -m github.com/Diegiwg/tt@"$VERSION"

# Test if the version was published
echo "Getting the version from https://pkg.go.dev/github.com/Diegiwg/tt@$VERSION..."
OK=$(curl -o /dev/null -X 'POST' -s -w "%{http_code}\n" https://pkg.go.dev/github.com/Diegiwg/tt@"$VERSION")

if [ "$OK" = "200" ]; then
    echo "Version $VERSION was published"
else
    echo "Version $VERSION was not published"
    exit 1
fi