.PHONY: buildBinary buildImage

buildBinary:
	CGO_ENABLED=0 GOOS=linux go build -o user-service/main ../user-service
	CGO_ENABLED=0 GOOS=linux go build -o user-web/main ../user-web


buildImage:
	docker build -f ./Dockerfile --tag liaotian/user-service:v1.0.4 --target user-service . --no-cache
	docker build -f ./Dockerfile --tag liaotian/user-web:v1.0.4 --target user-web . --no-cache

buildApplyYaml:
	kubectl apply -f k8s_rbac.yaml
	kubectl apply -f user-service/k8s.yaml
	kubectl apply -f user-web/k8s.yaml