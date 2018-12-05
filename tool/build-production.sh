#!/usr/bin/env bash

set -ev

# Creating build folder
mkdir -p build/www

# Build analyze
cp -r analyze/tool ./build/analyze/

# Build roasterc
docker run --rm -v "$PWD/www":/usr/src/roaster -w /usr/src/roaster node /bin/bash -c "npm install && npm run build" && cp -r ./www/dist ./build/www/

# Build roasterd
docker run --rm -v "$PWD":/usr/src/roaster -w /usr/src/roaster/cmd/roasterd golang:1.11 go build -o /usr/src/roaster/build/roasterd

# Current build folder content
du -h -d1 build/*

# Packaging into roaster container
docker build -t roaster .
