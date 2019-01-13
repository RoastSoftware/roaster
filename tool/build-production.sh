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
git clone https://github.com/danielgtaylor/aglio.git /tmp/roaster-build-aglio
pushd /tmp/roaster-build-aglio
docker build -t aglio .
popd
docker run --rm -v $PWD/doc/restapi:/tmp -t aglio --theme-template triple \
	-i /tmp/Roaster-REST-API.md -o /tmp/index.html \
	&& mv -f ./doc/restapi/index.html ./build/doc/restapi/
rm -rf /tmp/roaster-build-aglio

# Current build folder content
du -h -d1 build/*

# Packaging into roaster container
docker build -t roaster .
