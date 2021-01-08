.PHONY: startOperator apply delete

startOperator:
	kubectl apply -f operator-prereqs.yaml
	kubectl apply -f operator-deployment.yaml

apply:
	kubectl apply -f deployment.yaml

delete:
	kubectl delete -f deployment.yaml