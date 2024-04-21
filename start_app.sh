#!/bin/bash

minikube start --driver docker

kubectl apply -f karmadb-config.yaml

kubectl apply -f karmadb-secret.yaml

kubectl apply -f karmadb.yaml

kubectl apply -f mongo-express.yaml

kubectl apply -f karma-api.yaml

