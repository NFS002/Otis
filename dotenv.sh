#!/bin/bash
# Loads a .env file and exports all variables set in that file to the system environment

if [ "$#" -ne 1 ]; then
  echo "Usage: ./dotenv.sh <path_to_env_file>"
else
  set -a
  source "$1"
  set +a
fi