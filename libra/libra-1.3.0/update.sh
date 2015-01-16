#!/bin/bash

if [ $# -ne 1 ]; then
	echo "run as ./update.sh <version_number>"
	exit
fi

#docker build -t 115.28.62.181:5000/libra:$1 .
#docker push 115.28.62.181:5000/libra:$1

#docker build -t 10.167.15.204:5000/libra:$1 .
#docker push 10.167.15.204:5000/libra:$1

docker build -t 192.168.2.103:5000/libra:$1 .
docker push 192.168.2.103:5000/libra:$1
