.PHONY: apply delete

apply:
	helm install --name prometheus-operator --namespace monitoring ./prometheus-operator
	helm install --name prometheus --namespace monitoring --set serviceMonitorsSelector.app=prometheus --set ruleSelector.app=prometheus ./prometheus
#	helm install --name alertmanager --namespace=monitoring stable/alertmanager
#	kubectl patch svc prometheus-operator-grafana -p '{"spec":{"type":"NodePort"}}' -n monitoring
#	kubectl patch svc prometheus-operator-prometheus -p '{"spec":{"type":"NodePort"}}' -n monitoring

delete:
	helm delete prometheus
	helm delete prometheus-operator
	kubectl delete crd alertmanagers.monitoring.coreos.com podmonitors.monitoring.coreos.com prometheuses.monitoring.coreos.com prometheusrules.monitoring.coreos.com servicemonitors.monitoring.coreos.com thanosrulers.monitoring.coreos.com
	helm del --purge prometheus
	helm del --purge prometheus-operator
