#!/bin/bash
# This script recursively searches all files and subdirectories for all go.mod files,
# and removes the first two characters at the start of each line that starts with '//' and ends with '//_LOCAL'
# This script is intended to be run after 'git pull' to set up local development environmen to use
# local modules instead of downloading remote versions. Note that this script is only tested on MAC OS X, but should work on Linux also.

find "$(cd .; pwd)" -name "go.mod" -exec sed -i '' -e '/\/\/[[:space:]]*.*\/\/_LOCAL/s/^..//' {} + ;