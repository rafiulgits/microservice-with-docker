# node image
FROM node:12

# working directory inside docker 
WORKDIR /app

# copy package file into docker to install required packages
COPY package.json /app

# installing required packages
RUN npm install

# copy all other source files into docker
COPY . /app

# installing required packages inside docker

# running node server
CMD node index.js

# exposing node server port with host machine
EXPOSE 3000