#!/bin/bash

cd ./outputs || exit

for file in *; do
    aws s3 cp ${file} s3://bucket/path/to/${file}
done
