.PHONY: apply delete

apply:
	helm3 install prometheus-operator -n monitoring ./prometheus-operator
	helm3 install prometheus -n monitoring --set serviceMonitorsSelector.app=prometheus --set ruleSelector.app=prometheus ./prometheus
#	helm install --name alertmanager --namespace=monitoring stable/alertmanager
#	kubectl patch svc prometheus-operator-grafana -p '{"spec":{"type":"NodePort"}}' -n monitoring
#	kubectl patch svc prometheus-operator-prometheus -p '{"spec":{"type":"NodePort"}}' -n monitoring

delete:
	helm3 uninstall prometheus -n monitoring
	helm3 uninstall prometheus-operator -n monitoring
	kubectl delete crd alertmanagers.monitoring.coreos.com podmonitors.monitoring.coreos.com prometheuses.monitoring.coreos.com prometheusrules.monitoring.coreos.com servicemonitors.monitoring.coreos.com thanosrulers.monitoring.coreos.com
