# Dockee

## Build

### Clone and build

- clone and build with the following command:

```
git clone https://github.com/StLeoX/dockee.git
cd dockee
go build
```

## Use

### Commands

```

USAGE:
sudo dockee [global options] command [command options] [arguments...]

COMMANDS:
exec exec a command into container
init Init container process run user's process in container. Do not call it
outside
ps list all containers
logs print logs of a container
rm remove unused container
run create a container: my-docker run -ti [command]
stop stop a container
network container network commands
commit commit a container to image
image image commands
help, h Shows a list of commands or help for one command

GLOBAL OPTIONS:
--help, -h show help

```

#### run command

```
NAME:
   dockee - A docker-like container CLI application

USAGE:
dockee [global options] command [command options] [arguments...]

COMMANDs:
run      create a container
commit   commit a container into images
ps       list all the containers
logs     print logs of a container
exec     exec a command into container
start    start a container
stop     stop a container
rm       remove unused containers
network  container network commands
help, h  Shows a list of commands or help for one command

OPTIONs:
--it            enable tty
-d              detach container
-m value        memory limit
--cpu value     cpu limit
--cpuset value  cpuset limit
-v value        volume
--image value   the image name used to build the container
--name value    container name
-e value        set environment
--net value     container network
-p value        port mapping


GLOBAL OPTIONs:
--debug enable debug mode

EXAMPLEs:
sudo ./dockee run -it --image busybox /bin/sh
sudo ./dockee run -d --image busybox /bin/ps
sudo ./dockee network create --driver bridge --subnet 192.168.33.0/24 bridge33
sudo ./dockee run -it --image busybox --net bridge33 '/bin/ip a'
```

#### network command

```
NAME:
   dockee network - container network commands

USAGE:
   dockee network command [command options] [arguments...]

COMMANDs:
   create  create a container network
   list    list container network
   remove  remove container network
```

### Images

pull a centos8 images from docker hub

```

docker pull roboxes/centos8:latest

docker run -it --name mycentos roboxes/centos8:latest

docker export - o mycentos.tar 835360ffl6b8 # 835360ffl6b8 is docker's container id

mv mycentos.tar dockee/Images

```


