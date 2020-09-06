.PHONY: apply delete
apply:
	kubectl apply -f ./apisix-gw-config-cm.yaml
	kubectl apply -f ./deployment.yaml
	kubectl apply -f ./service.yaml
	kubectl apply -f ./dashboard.yaml
	kubectl apply -f ./service-monitor-for-prometheus.yaml

delete:
	kubectl delete -f ./apisix-gw-config-cm.yaml
	kubectl delete -f ./deployment.yaml
	kubectl delete -f ./service.yaml
	kubectl delete -f ./dashboard.yaml
	kubectl delete -f ./service-monitor-for-prometheus.yaml