

## Build

- This project is developed on linux, with golang version `go1.21.4 linux/amd64`.



### Build on linux

- clone and build with the following command:

```
git clone https://github.com/StLeoX/dockee.git
cd dockee
go build .
```



### Environment

```
Linux ubuntu-linux 4.4.0-142-generic #168~14.04.1-Ubuntu SMP Sat Jan 19 11:26:28 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux
```



### clone and Setup

```
git clone https://github.com/StLeoX/dockee.git
cd  dockee
go build .
```



### Commands

```
USAGE:
   sudo dockee [global options] command [command options] [arguments...]

COMMANDS:
   exec     exec a command into container
   init     Init container process run user's process in container. Do not call it outside
   ps       list all containers
   logs     print logs of a container
   rm       remove unused container
   run      create a container: my-docker run -ti [command]
   stop     stop a container
   network  container network commands
   commit   commit a container to image
   image    image commands
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```



#### run command

```
USAGE:
   sudo dockee run [command options] [arguments...]

OPTIONS:
   --ti              enable tty
   -d                detach container
   --mem value       memory limit
   --cpushare value  cpushare limit
   --cpuset value    cpuset limit
   --name value      container name
   -e value          set environment
   --net value       container network
   -p value          port mapping
```



#### network command

```
USAGE:
   sudo dockee network command [command options] [arguments...]

COMMANDS:
   create  create a container network
   ls      list container network
   rm      remove container network

OPTIONS:
   --help, -h  show help
```



### Images

pull a centos8 images from docker hub

```
docker pull roboxes/centos8:latest

docker run -it --name mycentos roboxes/centos8:latest

docker export - o mycentos.tar 835360ffl6b8 （container ID)

mv mycentos.tar dockee/Images
```


