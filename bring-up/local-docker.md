# Local Docker Cluster Bring-Up

## Prereqs

Docker installed and
configured. [Docker Toolbox](https://docs.docker.com/toolbox/overview/)
recommended.

## Start Kubernetes

```
export K8S_VERSION=v1.2.4
```

```
docker run \
    --volume=/:/rootfs:ro \
    --volume=/sys:/sys:ro \
    --volume=/var/lib/docker/:/var/lib/docker:rw \
    --volume=/var/lib/kubelet/:/var/lib/kubelet:rw \
    --volume=/var/run:/var/run:rw \
    --net=host \
    --pid=host \
    --privileged=true \
    --name=kubelet \
    -d \
    gcr.io/google_containers/hyperkube-amd64:${K8S_VERSION} \
    /hyperkube kubelet \
        --containerized \
        --hostname-override="127.0.0.1" \
        --address="0.0.0.0" \
        --api-servers=http://localhost:8080 \
        --config=/etc/kubernetes/manifests \
        --cluster-dns=10.0.0.10 \
        --cluster-domain=cluster.local \
        --allow-privileged=true --v=2
```

## Setup Kubectl

Download `kubectl` to your local directory, and add your local
directory to your path:

Mac OS X
```
curl -O http://storage.googleapis.com/kubernetes-release/release/${K8S_VERSION}/bin/darwin/amd64/kubectl
chmod 755 kubectl
PATH=$PATH:$(pwd)
```

Linux
```
wget http://storage.googleapis.com/kubernetes-release/release/${K8S_VERSION}/bin/linux/amd64/kubectl
chmod 755 kubectl
PATH=$PATH:$(pwd)
```

Windows (with bash)
```
wget http://storage.googleapis.com/kubernetes-release/release/${K8S_VERSION}/bin/windows/amd64/kubectl.exe
chmod 755 kubectl
PATH=$PATH:$(pwd)
```
## Windows / Mac OS X

In another terminal, create an ssh tunnel for kubectl to talk to your cluster. Leave this running.

```
docker-machine ssh $(docker-machine active) -N -L 8080:localhost:8080
```
## Create kubernetes cluster configuration (Optional)

```
kubectl config set-cluster test-doc --server=http://localhost:8080
kubectl config set-context test-doc --cluster=test-doc
kubectl config use-context test-doc
```

## Check that your cluster is running

```
kubectl get nodes
```
```
NAME        STATUS    AGE
127.0.0.1   Ready     2s
```
## Start DNS

```
kubectl create -f ./skydns.yaml
```
```
replicationcontroller "kube-dns-v11" created
service "kube-dns" created
```
