#!/bin/bash

echo "Building the http-server binary"
dep ensure -v

go install github.com/alecholmez/http-server/h2tp
echo "Done"
