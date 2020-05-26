FROM alpine

ENV APP_DIR /usr/myapp
RUN mkdir -p $APP_DIR
ADD . $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./myapp)

EXPOSE 80
