#!/bin/bash

# 进到当前sh脚本目录
cd $(cd $(dirname ${BASH_SOURCE:-$0});pwd)

kubectl apply -f ./etcd.yaml

kubectl apply -f ./apisix-gw-config-cm.yaml

kubectl apply -f ./deployment.yaml

kubectl apply -f ./service.yaml

kubectl apply -f ./dashboard.yaml