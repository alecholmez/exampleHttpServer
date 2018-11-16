#!/bin/bash

echo "Building the http-server docker container"
docker build -t alecholmez/http-server:latest .
echo "Done"
