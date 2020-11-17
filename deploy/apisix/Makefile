.PHONY: apply delete

apply:
	kubectl create -n liaotian configmap apisix-gw-config.yaml --from-file=config.yaml
	kubectl apply -n liaotian -f deployment.yaml
	kubectl apply -n liaotian -f service.yaml
	kubectl apply -n liaotian -f dashboard.yaml

delete:
	kubectl delete -n liaotian cm apisix-gw-config.yaml
	kubectl delete -n liaotian -f deployment.yaml
	kubectl delete -n liaotian -f service.yaml
	kubectl delete -n liaotian -f dashboard.yaml