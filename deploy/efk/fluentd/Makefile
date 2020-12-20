.PHONY: apply delete

apply:
	kubectl apply -f fluentd-configmap.yaml
	kubectl apply -f fluentd-daemonset.yaml

delete:
	kubectl delete -f fluentd-configmap.yaml
	kubectl delete -f fluentd-daemonset.yaml