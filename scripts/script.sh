#!/bin/bash

type=$1
user=$2

echo "Args: ${type}, ${user}"  > /app/output.txt
echo "File created!"

exit 0