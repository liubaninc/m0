#!/bin/bash

chain_id=${1:-testnetA}

docker rm -f $(docker ps -a | grep $chain_id | awk '{print $1}')
docker volume rm $(docker volume ls -q | grep $chain_id)