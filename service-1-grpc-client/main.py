from flask import Flask

import grpc
from microservice_comms import microservice_comms_pb2, microservice_comms_pb2_grpc

app = Flask('docker_grpc_example_client')


# easy function to just run the client and wait for a response.
@app.route('/')
def run_client():
    # create an insecure channel at port 4317
    try:
        # create that channel on your local host.
        # the ip address here should match the IP that your server has printed out.
        # Ex. 'Server located at: 172.17.0.3'
        with grpc.insecure_channel('service-2-grpc-server.default.svc.cluster.local:4317') as channel:
            # we create a stub that uses that channel.
            stub = microservice_comms_pb2_grpc.HandlerStub(channel)
            # we send a request and get a response from the stub.
            response = stub.HandleRequest(microservice_comms_pb2.microserviceCommsRequest(col_name="toy_example"))
        # we print the response as a table in html. Thus, we have to replace '/n' with '<br>'
        print("Client received the following response:")
        # return response as cleaned html code
        response = bytes(str(response), "utf-8").decode("unicode_escape").replace('df: "', '').replace('"', '')
        return "<h1>Client Received:</h1>" + response
    except Exception as e:
        return "<h1>ERROR</h1>" + str(e)


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=5000)
