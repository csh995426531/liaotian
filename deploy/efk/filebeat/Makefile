.PHONY: apply delete

apply:
	helm3 install -n logging filebeat .

delete:
	helm3 uninstall -n logging filebeat