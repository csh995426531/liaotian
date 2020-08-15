#!/bin/bash
DOCKER_IMAGE_HOST="registry.cn-hangzhou.aliyuncs.com"
DOCKER_IMAGE_NAMESPACE="fumi_fm"
DOCKER_IMAGE_HUB="user-web"

IMAGE_TAG="v1.0"

WORK_PATH=$(dirname $0)

# 当前位置跳到脚本位置
cd ./${WORK_PATH}

# 取到脚本目录
WORK_PATH=$(pwd)

mkdir bin

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct

# 跨平台 编译Linux 需要交叉编译
# CGO_ENABLED=0 GOOS=linux go build -o ${WORK_PATH}/bin/main ${WORK_PATH}/../user-web

go build -o ${WORK_PATH}/bin/main ${WORK_PATH}/../user-web

echo -e "\033[32m编译完成: \033[0m ${WORK_PATH}/bin/main"

# 容器制作
docker build -t ${DOCKER_IMAGE_HOST}/${DOCKER_IMAGE_NAMESPACE}/${DOCKER_IMAGE_HUB}:${IMAGE_TAG} -f ./Dockerfile .

echo -e "\033[32m镜像打包完成，请推送: \033[0m ${DOCKER_IMAGE_HOST}/${DOCKER_IMAGE_NAMESPACE}/${DOCKER_IMAGE_HUB}:${IMAGE_TAG}\n"

# 删除原二进制文件以及所在目录
rm -rf bin

echo -e "\033[32m残留二进制文件清理成功"

# 推送
docker push ${DOCKER_IMAGE_HOST}/${DOCKER_IMAGE_NAMESPACE}/${DOCKER_IMAGE_HUB}:${IMAGE_TAG}

echo -e "\033[32m镜像推送完成: \033[0m ${DOCKER_IMAGE_HOST}/${DOCKER_IMAGE_NAMESPACE}/${DOCKER_IMAGE_HUB}:${IMAGE_TAG} \n"