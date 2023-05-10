install_helm:
	helm install service-2-grpc-server charts/service-2-grpc-server/ --values charts/service-2-grpc-server/values.yaml
	helm install service-1-grpc-client charts/service-1-grpc-client/ --values charts/service-1-grpc-client/values.yaml

uninstall_helm:
	helm uninstall service-1-grpc-client
	helm uninstall service-2-grpc-server

define update_client
	cd $(1) && rm -f go.mod go.sum && \
	go mod init example.com/$(strip $(1)) && \
	go mod edit --replace example.com/microservice=../microservice && \
	go get example.com/microservice && \
	go mod tidy
endef

generate_go_proto:
	mkdir -p microservice
	protoc --go-grpc_out=. --go_out=. proto/microservice.proto
	cd microservice && rm -f go.mod go.sum && go mod init example.com/microservice && go get -u && go mod tidy
	@$(call update_client, "service-1-grpc-client")
	@$(call update_client, "service-2-grpc-server")

build_images:
	cd service-1-grpc-client && docker build -t hannankhanintelops/service-1-grpc-client-image .
	cd service-2-grpc-server && docker build -t hannankhanintelops/service-2-grpc-server-image .

run_images:
	cd service-1-grpc-client && docker run -t service-1-grpc-client-container --network=host hannankhanintelops/service-1-grpc-client-image
	cd service-2-grpc-server && docker run -t service-2-grpc-server-container --network=host hannankhanintelops/service-2-grpc-server-image

serve_docker_locally: build_images run_images

reinstall: uninstall_helm install_helm