#!/bin/bash
# Loads a .env file and exports all variables set in that file to the system environment

if [ "${0##*/}" != "bash"  ]; then
  echo "Warning: File should be sourced. Usage: source ./dotenv.sh <path_to_env_file>"
fi


if [ "$#" -ne 1 ]; then
  echo "Usage: source ./dotenv.sh <path_to_env_file>"
else
  set -a
  source "$1"
  set +a
  echo "$1 was loaded into the system environment"
fi