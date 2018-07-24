#!/bin/bash

pushd frontend
GOOS=linux go build
popd

pushd backend
GOOS=linux go build
popd
# docker build -t simple-service
# docker push simple-service
