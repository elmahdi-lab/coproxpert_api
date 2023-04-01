#!/usr/bin/env bash

# output to console
set -x


# prune docker
docker system prune -a -f

docker rmi coproxpert_api

# build docker
docker-compose up --build