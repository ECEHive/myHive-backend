#!/bin/bash
set -e
if [ -z "$1" ]
    then
          echo "No argument supplied"
              exit
fi
echo "building version $1"
docker build . -t "ecehive/myhive_backend:$1"
docker tag "ecehive/myhive_backend:$1" "ecehive/myhive_backend:latest"
docker push "ecehive/myhive_backend:latest"
git tag "v$1"
git push --tags
