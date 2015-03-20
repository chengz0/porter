#!/bin/bash

git add .
git commit -a -m "nima"
git tag -a $1 -m "$1" -s
git push origin
