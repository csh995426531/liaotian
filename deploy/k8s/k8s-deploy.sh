#!/bin/bash
kubectl apply -f ./k8s_rbac.yaml

source ../traefik/k8s-deploy.sh

source ../user-service/k8s-deploy.sh

source ../user-web/k8s-deploy.sh