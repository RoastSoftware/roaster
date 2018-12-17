#!/usr/bin/env bash

set -ev

# Creating build folder
mkdir -p build/{www,doc/restapi}

# Copying dependency lists
cp requirements.freeze.txt build/

# Build roasterc
docker run --rm -v "$PWD/www":/usr/src/roaster -w /usr/src/roaster node /bin/bash -c "npm install && npm run build" && cp -r ./www/dist ./build/www/

# Build roasterd
docker run --rm -v "$PWD":/usr/src/roaster -w /usr/src/roaster/cmd/roasterd golang:1.11 go build -o /usr/src/roaster/build/roasterd

# Build API docs
docker run --rm -v $PWD/doc/restapi:/doc quay.io/bukalapak/snowboard html \
	-o index.html Roaster-REST-API.md && \
	cp ./doc/restapi/index.html ./build/doc/ && \
	rm -f ./doc/restapi/index.html

# Current build folder content
du -h -d1 build/*

# Packaging into roaster container
docker build -t roaster .
