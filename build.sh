#!/bin/bash

echo "Building the http-server docker images"
docker build -t alecholmez/http-server --force-rm .

docker pull mongo
echo "Done"