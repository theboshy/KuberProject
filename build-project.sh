#!/bin/bash

protoc -I ./mcs --go_out=plugins=grpc:./pb ./mcs/*.proto

docker build -t local/gcd -f Dockerfile.gcd .

kubectl apply -f gcd-deployment.yaml