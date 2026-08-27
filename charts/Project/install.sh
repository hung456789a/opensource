docker pull mongodb/mongodb-enterprise-server:latest

docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-enterprise-server:latest

docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-enterprise-server:5.0-ubuntu2004

docker container ls
