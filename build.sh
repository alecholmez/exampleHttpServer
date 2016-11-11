#!/bin/bash

echo "Building the http-server binary"
govendor sync

go install github.com/alecholmez/http-server/h2tp
echo "Done"
