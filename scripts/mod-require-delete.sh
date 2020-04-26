#!/bin/bash
# This script recursively searches all files and subdirectories for all go.mod files,
# and removes any 'replace block' in that file.
# This script is intended to be run before deployment or 'git push' to clear any dependencies
# on local module versions. Note that this script is only tested on MAC OS X, but should work on Linux also.

FILES=$(find "$(cd .; pwd)" -name "go.mod" -type f)
for file in $FILES
do
    awk '/require/,/\)/{sub(/.*/,"")} 1' $file > tmp && mv tmp $file
done