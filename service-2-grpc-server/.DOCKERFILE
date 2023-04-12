FROM python:3

WORKDIR /usr/src/server

COPY service-2-grpc-server/requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY ./service-2-grpc-server service-2-grpc-server/
COPY ./microservice_comms microservice_comms/

# add our proto module to the python path.
ENV PYTHONPATH="$PYTHONPATH:/usr/src/server/microservice_comms"

# we expose the listening port.
ENV PORT 4317
EXPOSE 4317

CMD [ "python", "./service-2-grpc-server/main.py" ]