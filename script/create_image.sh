#!/bin/sh
imageName="myapp:1.0"
cd src/myapp 
docker build -t justware/${imageName} .
#docker save justware/${imageName} > ${imageName}.tar
#microk8s.ctr image import ${imageName}.tar
