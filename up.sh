#!/bin/bash
tag=${DOCKER_REGISTRY}/$(basename ${PWD})
set -ex
docker build -t ${tag} .
docker push ${tag}
docker images --digests ${tag}
