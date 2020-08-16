#!/bin/bash

# 进到当前sh脚本目录
cd $(cd $(dirname ${BASH_SOURCE:-$0});pwd)

kubectl apply -f ./k8s_rbac.yaml

source ../traefik/k8s-deploy.sh

source ../user-service/k8s-deploy.sh

source ../user-web/k8s-deploy.sh