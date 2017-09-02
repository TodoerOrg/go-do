# go-do
todoer backend implemented in go



# docker experiments 

build the my-golang-app docker image from Dockerfile 
> docker build -t my-golang-app .

run the image inside a new container ( do not exit immediately )
> docker run -dit my-golang-app

list all running docker container ( note the container id )
> docker ps

attach a bash shell to the container 
> docker exec -i -t <CONTAINER_ID> /bin/bash

stop a running container
> docker stop <CONTAINER_ID>

start a golang server inside a docker container 
> docker run --publish 8080:8080 --name docker-server-test --rm my-golang-app

# docker compose experiments

based on the learnings from:

https://hharnisc.github.io/2016/06/19/integration-testing-with-docker-compose.html

we can create a docker-compose that describes api integration testing setup with docker containers.

