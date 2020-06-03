#!/bin/sh
#there is godep issue in China
#so need execute: godep save firstly
cd src/myapp && CGO_ENABLED=0 godep go build -ldflags '-d -w -s'