#We will be using the latest golang image.
FROM golang:latest AS BUILD

# We set the working directory.
WORKDIR /usr/src/app

# We copy our entire project over to the container image.
COPY . .

# We copy the main directories go mod and go sum to app/.
# These required packages are then installed by the container.
RUN cd service-1-grpc-client && go mod download && cd ..

EXPOSE 5000

# We navigate to the main package, build the go executable for execution in alpine linux (thats what all those extra
# things are before go build), and delete unnecessary directories.
RUN cd ./service-1-grpc-client && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../script_artifact && cd .. && rm -rf main go.mod go.sum

# to get processes: ps -aef --forest
# we would use the below command to create a sleep container for our app, for dev purposes.
ENTRYPOINT [ "tail", "-f", "/dev/null" ]

#FROM python:3
#
#WORKDIR /usr/src/client
#
#COPY service-1-grpc-client/requirements.txt ./
#
#RUN pip install --no-cache-dir -r requirements.txt && mkdir microservice_comms
#
#COPY ./service-1-grpc-client service-1-grpc-client/
#COPY ./microservice_comms microservice_comms/
#
## add our proto module to the python path.
#ENV PYTHONPATH="$PYTHONPATH:/usr/src/client/microservice_comms"
#
## we must expose the port that will be running our flask app.
#EXPOSE 5000
#
## we use python as an entrypoint (consider it an executable)
#ENTRYPOINT [ "python" ]
#
#CMD [ "./service-1-grpc-client/main.py" ]