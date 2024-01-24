#!/bin/bash

if [ -d "build" ]; then
    echo "build dir existed, deleting..."
    rm -rf build
fi

mkdir build

cd build
service_path="../src/service"

find "$service_path" -type d | while read -r service; do
    if [ $service != $service_path ]; then
      service_name="$(basename "$service")"
      echo "building $service_name"
      go build $service
    fi
done
echo "building web"
go build ../src/web