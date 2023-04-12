FROM python:3

WORKDIR /usr/src/client

COPY service-1-grpc-client/requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt && mkdir microservice_comms

COPY ./service-1-grpc-client service-1-grpc-client/
COPY ./microservice_comms microservice_comms/

# add our proto module to the python path.
ENV PYTHONPATH="$PYTHONPATH:/usr/src/client/microservice_comms"

# we must expose the port that will be running our flask app.
EXPOSE 5000

# we use python as an entrypoint (consider it an executable)
ENTRYPOINT [ "python" ]

CMD [ "./service-1-grpc-client/main.py" ]