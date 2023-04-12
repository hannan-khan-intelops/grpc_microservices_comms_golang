install_helm:
	helm install service-2-grpc-server charts/service-2-grpc-server/ --values charts/service-2-grpc-server/values.yaml
	helm install service-1-grpc-client charts/service-1-grpc-client/ --values charts/service-1-grpc-client/values.yaml

uninstall_helm:
	helm uninstall service-1-grpc-client
	helm uninstall service-2-grpc-server

reinstall: uninstall_helm install_helm