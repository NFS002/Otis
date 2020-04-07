#!/bin/bash
# This script recursively searches all files and subdirectories for all *.mod files,
# and removes any lines that end with '//_LOCAL'
# This script is intended to be run as part of a preprocessing stage to prepare the project for deployment \
# Note that this script is only tested on MAC OS X, but should work on Linux also.

find "$(cd ..; pwd)" -name "go.mod" -exec sed -i '' '/^.*\/\/_LOCAL.*$/d' {} + ;
