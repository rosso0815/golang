# DOCKER

```
export DOCKER_HOST="ssh://pfistera@192.168.0.110"

# lets build
docker build -t hello .

# lets run
docker rm hello
docker run -d -p8080:8080 --name hello hello

```


# microk8s

```
microk8s start
microk8s status --wait-ready
microk8s enable dashboard dns registry istio
```