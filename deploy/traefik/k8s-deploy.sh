#!/bin/bash

# 进到当前sh脚本目录
cd $(cd $(dirname ${BASH_SOURCE:-$0});pwd)
kubectl apply -f .