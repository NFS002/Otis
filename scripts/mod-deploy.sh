#!/bin/bash
# This script recursively searches all files and subdirectories for all go.mod files,
# and comments out (inserts a '//' at the start of) any lines that end with '//_LOCAL' that are not already commented out.
# This script is intended to be run as part of a preprocessing stage to prepare the project for deployment.
# Note that this script is only tested on MAC OS X, but should work on Linux also.

find "$(cd ..; pwd)" -name "go.mod" -exec sed -i '' -e 's/^[^//].*\/\/_LOCAL.*$/\/\/&/' {} + ;
