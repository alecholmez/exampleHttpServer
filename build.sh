#!/bin/bash

echo "Building the http-server binary"
govendor sync

go install github.com/alecholmez/http-server
echo "Done"
