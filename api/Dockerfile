# the name of the base image

# FROM ubuntu:latest
FROM --platform=linux alpine:latest

# alpine (or ubuntu) is the base image. You can try different base images.
# after building your docker image, you will see how different base images have different sizes.
# e.g. ubuntu is huge, alpine is stripped back
# Alpine is chosen for its small footprint compared to Ubuntu

COPY ./bin/server /opt/ 
# COPY copies the file server into the /opt/ directory within the image

CMD [ "/opt/server" ]

# CMD will tell Docker what command to execute when our image is used to start a container (aka will start our Hello World server)
# If you comment this out, you have to manually go to the file in the image and run it