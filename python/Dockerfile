# syntax=docker/dockerfile:1

# base python image for custom image
FROM python:3.9.13-slim-buster

# create working directory and install pip dependencies
WORKDIR /mazaipy
RUN pip install bottle
RUN apt-get update && apt-get install -y curl

# copy python project files from local to /mazaipy image working directory
COPY . .

# run server  
CMD [ "python3", "server.py"]

