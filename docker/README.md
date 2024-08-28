## docker

## useful docker commands:

1. search for a specific image from Docker Hub
```
$ docker search image_name
```

2. download an image from a registry
```
$ docker pull image_name
```

3. if you want to see images on your system
```
$ docker images
```

4. run a container:
```
$ docker run [OPTIONS] IMAGE [COMMAND] [ARGS...]

Example:
$ docker run -it ubuntu /bin/bash -c "echo 'hello $USER'"

$ docker run -it ubuntu /bin/bash
```

5. if you are stuck
```
$ docker --help
```
6. see all running containers
```
$ docker ps --all
$ docker ps -a
```

7. see the last N created containers, N is a number by default N is -1
```
$ docker ps --last N
$ docker ps -n N
```

8. to remove an image
```
$ docker rmi image_name
```

9. 