# grpc_microservices_comms_golang

## gRPC Commands
Install:
```shell
pip install grpcio-tools
```
Test out:
```shell
python -m grpc_tools.protoc --version
```
Compiling a proto file:
```shell
python -m grpc_tools.protoc --proto_path=proto --python_out=microservice_comms --grpc_python_out=microservice_comms proto/microservice_comms.proto
```
