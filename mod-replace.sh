#!/bin/bash
# This script recursively searches all files and subdirectories for all *.mod files,
# and removes any 'replace' statements
find "$(cd ..; pwd)" -name "filename" -exec sed -E -i '' 's/gitlab.com\/otis-team\/.*$/ /'


