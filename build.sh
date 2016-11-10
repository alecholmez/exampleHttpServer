#!/bin/bash

echo "Building the http-server binary"
govendor sync

go install
echo "Done"
