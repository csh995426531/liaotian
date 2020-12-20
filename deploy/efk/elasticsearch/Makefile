.PHONY: apply delete

apply:
	kubectl apply -f elasticsearch-svc.yaml
	kubectl apply -f elasticsearch-statefulset.yaml
delete:
	kubectl delete -f elasticsearch-svc.yaml
	kubectl delete -f elasticsearch-statefulset.yaml