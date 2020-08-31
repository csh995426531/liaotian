#!/bin/bash
DOCKER_IMAGE_NAMESPACE="liaotian_csh"
DOCKER_IMAGE_HUB="friend-service"
IMAGE_TAG="v1.0"

# 进到当前sh脚本目录
cd $(cd $(dirname ${BASH_SOURCE:-$0});pwd)

kubectl apply -f ./k8s.yaml
# 缩容扩容
kubectl scale --replicas=3 deployment/$DOCKER_IMAGE_HUB -n liaotian && kubectl scale --replicas=2 deployment/$DOCKER_IMAGE_HUB -n liaotian