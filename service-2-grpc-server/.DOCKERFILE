#We will be using the latest golang image.
FROM golang:latest

# We set the working directory.
WORKDIR /usr/src/app

# We copy our entire project over to the container image.
COPY . .

# We copy the main directories go mod and go sum to app/.
# These required packages are then installed by the container.
RUN cd service-2-grpc-server && go mod download && cd ..

EXPOSE 50005

# We navigate to the main package, build the go executable for execution in alpine linux (thats what all those extra
# things are before go build), and delete unnecessary directories.
RUN cd ./service-2-grpc-server && \
    env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../script_artifact main.go && \
    cd .. && rm -rf main go.mod go.sum && \
    chmod +x script_artifact

# to get processes: ps -aef --forest
# we would use the below command to create a sleep container for our app, for dev purposes.
#ENTRYPOINT [ "tail", "-f", "/dev/null" ]
CMD ["./script_artifact"]

#FROM python:3
#
#WORKDIR /usr/src/server
#
#COPY service-2-grpc-server/requirements.txt ./
#
#RUN pip install --no-cache-dir -r requirements.txt
#
#COPY ./service-2-grpc-server service-2-grpc-server/
#COPY ./microservice_comms microservice_comms/
#
## add our proto module to the python path.
#ENV PYTHONPATH="$PYTHONPATH:/usr/src/server/microservice_comms"
#
## we expose the listening port.
#ENV PORT 4317
#EXPOSE 4317
#
#CMD [ "python", "./service-2-grpc-server/main.py" ]