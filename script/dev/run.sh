#!/bin/sh
docker run -it --rm --name="hello" -p 8080:8080 \
   -v /home/alex/base/go:/go/src/app  hello:1.0