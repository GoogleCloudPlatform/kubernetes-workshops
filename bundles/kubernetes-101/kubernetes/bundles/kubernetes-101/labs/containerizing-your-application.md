# Containerizing your application

In this lab you will build your application, run it locally, and then package it into a container. Containers are more lightweight than a full virtual machine for your app and they package your application so that it runs the same across different environments such as development, QA, and production.

## Get the application code

The first hurdle to is the application itself.  How do you write it?  How do you deploy it?

First set up the code and the Go build environment. For this lab we'll be using an example app from GitHub: [kelseyhightower/app](https://github.com/kelseyhightower/app).
```bash
export GOPATH=~/go
mkdir -p $GOPATH/src/github.com/kelseyhightower
cd ~/go/src/github.com/kelseyhightower
git clone https://github.com/kelseyhightower/app
```

## Build the app

Now let's build the app as a static binary and test its functionality.
```bash
cd app/monolith
# main.go contains our app entry point
go build -tags netgo -ldflags "-extldflags '-lm -lstdc++ -static'" .
./monolith --http :10180 --health :10181 &
curl http://127.0.0.1:10180
curl http://127.0.0.1:10180/secure
# Yes the password for the next step is 'password'
TOKEN=$(curl http://127.0.0.1:10180/login -u user | jq -r '.token')
echo $TOKEN
curl -H "Authorization: Bearer $TOKEN" http://127.0.0.1:10180/secure
```

Once we have a binary, we can use Docker to package and distribute it.
```bash
docker build -t askcarter/monolith:1.0.0 .
# Optionally push to Docker Hub using docker push <your_repo>/monolith:1.0.0
docker run -d askcarter/monolith:1.0.0
docker ps
docker inspect <container-id>
curl http://<docker-ip>
docker stop <container-id>
docker rm <container-id>
docker rmi askcarter/monolith:1.0.0
```

## Title here

The next hurdle is the infrastructure needed to run manage in production. We'll use Kubernetes (and GKE) to handle that for us.

```bash
cd ~/kubernetes-workshop/kubernetes
kubectl run monolith --image askcarter/monolith:1.0.0
kubectl expose deployment monolith --port 80 --type LoadBalancer
kubectl scale deployment monolith --replicas 3
kubectl get service monolith
curl http://<External-IP>
kubectl delete services monolith
kubectl delete deployment monolith
```
