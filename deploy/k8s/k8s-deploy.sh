#!/bin/bash
 
# 进到当前sh脚本目录
echo $(cd $(dirname ${BASH_SOURCE:-$0});pwd)
cd $(cd $(dirname ${BASH_SOURCE:-$0});pwd)

kubectl apply -f ./k8s_rbac.yaml