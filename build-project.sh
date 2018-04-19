#!/bin/bash

#protoc -I ./mcs --go_out=plugins=grpc:./pb ./mcs/*.proto
protoc  ./mcs/rpcserviceproto.proto --go_out=plugins=grpc:./

eval $(minikube docker-env)

docker build -t local/gcd -f Dockerfile.gcd .

kubectl apply -f gcd-deployment.yaml
