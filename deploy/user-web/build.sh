#!/bin/bash
DOCKER_IMAGE_HOST="registry.cn-hangzhou.aliyuncs.com"
DOCKER_IMAGE_NAMESPACE="liaotian_csh"
DOCKER_IMAGE_HUB="user-web"
IMAGE_TAG="v1.0.4"

# 进到当前sh脚本目录
cd $(cd $(dirname ${BASH_SOURCE:-$0});pwd)

# 容器制作
docker build -t $DOCKER_IMAGE_HOST/$DOCKER_IMAGE_NAMESPACE/$DOCKER_IMAGE_HUB:$IMAGE_TAG -f ./Dockerfile ../../

echo -e "\033[32m镜像打包完成，请推送: \033[0m $DOCKER_IMAGE_HOST/$DOCKER_IMAGE_NAMESPACE/$DOCKER_IMAGE_HUB:$IMAGE_TAG\n"

# 登录镜像服务
docker login --username=崔徐徐ok $DOCKER_IMAGE_HOST -p a123456789A

# 推送
docker push $DOCKER_IMAGE_HOST/$DOCKER_IMAGE_NAMESPACE/$DOCKER_IMAGE_HUB:$IMAGE_TAG

echo -e "\033[32m镜像推送完成: \033[0m $DOCKER_IMAGE_HOST/$DOCKER_IMAGE_NAMESPACE/$DOCKER_IMAGE_HUB:$IMAGE_TAG \n"