install_helm:
	helm install service-2-grpc-server charts/service-2-grpc-server/ --values charts/service-2-grpc-server/values.yaml
	helm install service-1-grpc-client charts/service-1-grpc-client/ --values charts/service-1-grpc-client/values.yaml

uninstall_helm:
	helm uninstall service-1-grpc-client
	helm uninstall service-2-grpc-server

generate_go_proto:
	mkdir -p microservice
	protoc --go-grpc_out=. --go_out=. proto/microservice.proto
	cd microservice && rm -f go.mod go.sum && go mod init example.com/microservice && go get -u && go mod tidy
	cd service-1-grpc-client && rm -f go.mod go.sum && \
	go mod init example.com/service-1-grpc-client && \
	go mod edit --replace example.com/microservice=../microservice && \
	go get example.com/microservice && \
	go mod tidy

reinstall: uninstall_helm install_helm