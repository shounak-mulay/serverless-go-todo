#!/bin/bash

# Reading root folder to start the build process
if [ -z "$FUNCTIONS_ROOT" ]; then
  echo ""
  echo "FUNCTIONS_ROOT not set, setting to functions"
  export FUNCTIONS_ROOT="functions"
fi

echo ""
echo -e "\033[0;32m-------------------------- Building functions in /$FUNCTIONS_ROOT --------------------------\033[0m"
echo ""

# Loop through all folders with a go file in root
for file in $FUNCTIONS_ROOT/*/*.go; do
  dir=$(dirname $file)
  export GO111MODULE=on

  # Read .env file if present, else use default values
  if [ -f $dir/.env ]; then
      export $(cat $dir/.env | xargs)
  else
      export GOARCH=amd64 GOOS=linux FILE=main.go
  fi
    
  # Use dir name as the file build file name and call go build
  function_name=$(echo $dir | awk -F'/' '{print $NF}')
  echo "Building function: $function_name for $GOARCH/$GOOS"
  go build -ldflags="-s -w" -o "bin/$function_name" "$dir/$FILE"
  
  echo ""
  echo -e "\033[0;32m----------------------------------------------------\033[0m"
  echo ""
done
