Hello. This is my testing project name to do list using docker and no need to setup db ,please enjoy :)

#runing script
docker build -t to-do-list .
docker run -p 1112:1112 --name to-do-list-project -d to-do-list:latest

#stop and remove
docker stop to-do-list-project
docker rm to-do-list-project
