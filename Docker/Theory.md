Docker makes it really easy to install and run software without worrying about setup and  dependencies.

Docker has many components:
1. Docker Machine
2. docker server (docker daemon):Tool that is responsible for creating images, running containers etc
3. docker Images
4. Docker compose
5. Docker Client(docker CLI): tool that we are going to issue the commands
6. docker Hub

example command:

docker run -it redis
with the above command, first the docker cli connects to docker hub and it will download the image. Image will have have all the dependencies. and we will use this image to create container.
container is an running instance of an image.


container Lifecycle:

