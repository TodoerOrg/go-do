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

