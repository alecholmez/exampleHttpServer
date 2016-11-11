#!/bin/bash

echo "Building documentation page"
aglio --theme-variables streak --theme-template triple -i docs/main.md -o docs/index.html
echo "Done"
