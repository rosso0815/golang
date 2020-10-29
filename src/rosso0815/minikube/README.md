

https://medium.com/@rakateja/deploying-go-app-on-minikube-using-local-docker-image-81dad14dc7bd

docker system prune -a

docker build --rm -t apf-minikube .

docker images

docker run -p 8080:3000 apf-minikube:latest

docker export apf-minikube > apf-minikube.tar

minikube delete

minikube config set cpus 4
minikube config set memory 8192
minikube config set disk-size 20g
minikube config set vm-driver hyperkit

minikube start

eval $(minikube docker-env)

kubectl apply -f minikubeDeployment.yaml

kubectl get pods

kubectl get deployment apf-minikube

kubectl expose deployment apf-minikube --type=LoadBalancer --name=apf-minikube-


 kubectl get services apf-minikube-service
