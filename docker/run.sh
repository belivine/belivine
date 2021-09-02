#!/bin/bash

# goto basedir
cd $(dirname "$0")

# data
mkdir ./data/pgsql -p

# log
mkdir ./logs/go -p
mkdir ./logs/pgsql -p

# docker
docker-compose pull
docker-compose stop
docker-compose down
docker-compose up --build -d
