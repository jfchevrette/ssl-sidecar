#!/bin/bash

GIT_HASH=$(git describe --always --dirty)

CGO_ENABLED=0 GOOS=linux go build github.com/brendandburns/ssl-sidecar
docker build -t brendanburns/ssl-sidecar:${GIT_HASH} .
docker push brendanburns/ssl-sidecar:${GIT_HASH}

rm ssl-sidecar
