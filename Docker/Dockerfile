# goal is to write the dockerfile for redis

FROM alpine

RUN apk add --update redis
COPY ./ ./
COPY <source it will be from local> <destination inside container>

CMD ["redis-server"]


WORKDIR
docker run -p < port -> route incoming request to this port on localhost> :<port inside the container> <image-id>
docker run -p 8080:8080 <image-id>
docker run -p 8080:5000 <image-id>
The above command will make that whenever there is request to localhost:8080, please forward it to container:5050

If we dont specify the WORKDIR in dockerfile, the files gets copied to root directory.
if there are same files in root directory these will actually override those files. 
so it is better to create a new directory for the application.
we can do that using 
WORKDIR /app

this actually create a  app directory and copy the files into that directory.