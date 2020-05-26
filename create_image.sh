#!/bin/sh
imageName="myapp:1.0"
cd src/myapp && CGO_ENABLED=0 godep go build -ldflags '-d -w -s'
#there is godep issue in China
#so need execute: godep save firstly
docker build -t ${imageName} .
docker save ${imageName} > ${imageName}.tar
docker rmi ${imageName}
microk8s.ctr image import ${imageName}.tar
rm  ${imageName}  ${imageName}.tar


