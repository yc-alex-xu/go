#!/bin/sh
# stop & delete all container
docker stop $(docker ps -q) & docker rm $(docker ps -aq)
imageName="myapp:1.0"
cd src/myapp 
docker rmi justware/${imageName}
rm  ${imageName}  ${imageName}.tar
