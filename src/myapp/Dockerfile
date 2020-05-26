FROM alpine
LABEL maintainer="xuyc@sina.com"

ENV APP_DIR /usr/myapp
RUN mkdir -p $APP_DIR
ADD . $APP_DIR
EXPOSE 8080
# Set the entrypoint
ENTRYPOINT (cd $APP_DIR &&chmod +x myapp && ./myapp)


