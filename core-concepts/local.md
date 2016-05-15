# Core Concepts

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster
* Commands assume that you have a local copy of this git repository, 
  and `core-concepts` is the current directory.

## Lab

### Mac / Windows

You'll need to know the IP of the `docker-machine` vm that is your docker host:

```
docker-machine ip $(docker-machine active)
```

Use this when browsing to a node IP in place of `localhost` below.

### Pods

First define a single pod, see [pod.yaml](pod.yaml). Start the
Lobsters app from this pod declaration.

```
kubectl create -f ./pod.yaml
```

```
pod "lobsters" created
```

Check for pods

```
kubectl get pods
```

```
NAME       READY     STATUS    RESTARTS   AGE
lobsters   1/1       Running   0          1m
```

Delete the pod

```
kubectl delete pod lobsters
```

```
pod "lobsters" deleted
```

The pod is gone forever

### Service

To access Lobsters from outside the cluster, we'll need a service. The
service defined in [service-local.yaml](service-local.yaml) will route
traffic to any pod with the label `app: lobsters`, which matches our
pod definition. The service routs to the port labeled `web` in our pod
definition. The `type: NodePort` line allows traffic on a particular
port of each node to be routed to the service.

Create the service and pod:

```
kubectl create -f ./service-local.yaml,./pod.yaml
```

```
service "lobsters" created
pod "lobsters" created
```

Check the service's node port, yours will be different:

```
kubectl get svc lobsters -o yaml | grep nodePort
```

```
  - nodePort: 31618
```

Check that it is working by visiting the node IP with the port you
found `http://localhost:31618/`


Delete

```
kubectl delete pod,svc -l app=lobsters
```

```
pod "lobsters" deleted
service "lobsters" deleted
```

### Replication Controller

When we delete the pod above, it stays deleted. Your pod can also
disappear if your cluster node fails, or if the app crashes and can't
be restarted. The Kubernetes solution to this is a Replication
Controller, or RC for short. An RC will make sure a pod or number of
pods is always running somewhere in the cluster. In fact, it is almost
never appropriate to create individual pods as we did above.

See [rc.yaml](rc.yaml) for the RC definition. It is mostly the same as
the pod definition, but wrapped in an RC.

Start lobsters using an RC, use the same service definition:

```
kubectl create -f ./rc.yaml,./service-local.yaml
```

```
replicationcontroller "lobsters" created
service "lobsters" created
```

Check the service's node port, yours will be different:

```
kubectl get svc lobsters -o yaml | grep nodePort
```

```
  - nodePort: 31618
```

Check that it is working by visiting the node IP with the port you
found `http://localhost:31618/`


Now, look at the pod

```
kubectl get pods -o wide
```

```
NAME                   READY     STATUS    RESTARTS   AGE       NODE
lobsters-tx1sa         1/1       Running   0          21s       127.0.0.1
```

This pod was created by the replication controller. Try deleting the
pod, use the exact pod name for your pod:

```
kubectl delete pod lobsters-jf0xs
```

```
pod "lobsters-jf0xs" deleted
```

Now check again:

```
kubectl get pods -o wide
```

```
NAME                   READY     STATUS    RESTARTS   AGE       NODE
lobsters-l5fq3         1/1       Running   0          1s        127.0.0.1
```

A new pod was created!

Scaling is as easy as:

```
kubectl scale --replicas=5 rc lobsters
```

```
replicationcontroller "lobsters" scaled
```

Check the pods

```
kubectl get pods -o wide
```

```
NAME                   READY     STATUS    RESTARTS   AGE       NODE
lobsters-9ijsi         1/1       Running   0          6s        127.0.0.1
lobsters-l5fq3         1/1       Running   0          36s       127.0.0.1
lobsters-pfnlj         1/1       Running   0          6s        127.0.0.1
lobsters-sceuy         1/1       Running   0          6s        127.0.0.1
lobsters-txgwb         1/1       Running   0          6s        127.0.0.1
```

Also the RC

```
kubectl get rc lobsters -o wide
```

```
NAME       DESIRED   CURRENT   AGE       CONTAINER(S)   IMAGE(S)                             SELECTOR
lobsters   5         5         3m        lobsters       gcr.io/google-samples/lobsters:1.0   app=lobsters
```

The lobsters service will now route incoming traffic to any one of the
5 pods that match the selector app=lobsters.

We have a problem though. Each one of those replicas is using a local
SQLite file inside the container. You could post a new link to the
site, refresh and hit a different replica! The module
[Storing State](../state) will cover solutions.

Delete

```
kubectl delete rc,svc -l app=lobsters
```

```
replicationcontroller "lobsters" deleted
service "lobsters" deleted
```

You don't need to delete the pods, deleting the RC will take care of it.

### Deployments

A Deployment is very similar to an RC, the difference is apparent when
you change the configuration. All changes to RCs are instantaneous,
while change in Deployments are controlled.

Start up Lobsters using the Deployment declaration in
[dep.yaml](dep.yaml). You'll notice that it is almost identical to an
RC declaration.

```
kubectl create -f ./dep.yaml,./service-local.yaml
```

```
deployment "lobsters" created
service "lobsters" created
```

Deployments control and create Replica Sets (like RCs). Check the RS:

```
kubectl get rs -o wide
```

```
NAME                  DESIRED   CURRENT   AGE       CONTAINER(S)   IMAGE(S)                             SELECTOR
lobsters-1901432027   5         5         47s       lobsters       gcr.io/google-samples/lobsters:1.0   app=lobsters,pod-template-hash=1901432027
```

Load up the lobsters site in your browser using the same way as the
previous steps.

We will now change the version of the version of the
running Lobsters app. Where version 1.0 has `Example Lobsters`,
version 2.0 has `Example Lobsters 2.0!`. When we update the Deployment
to specify the 2.0 image, it will create a new RS and slowly increase
the replicas on the new RS, while decreasing the number of replicas on
the old RS, this will result in a smooth transition from version 1.0
to 2.0 while keeping around 5 total replicas running at all times.

```
kubectl apply -f ./dep-2.yaml
```

```
deployment "lobsters" configured
```

Now quickly:

```
kubectl get rs -o wide
NAME                  DESIRED   CURRENT   AGE       CONTAINER(S)   IMAGE(S)                             SELECTOR
lobsters-1901432027   2         2         16m       lobsters       gcr.io/google-samples/lobsters:1.0   app=lobsters,pod-template-hash=1901432027
lobsters-1980468444   4         4         8m        lobsters       gcr.io/google-samples/lobsters:2.0   app=lobsters,pod-template-hash=1980468444
```

You will see the the new RS created, and scaled up to 5
replicas. Also, `kubectl get pods` will show 5 pods that started
recently.

Refresh your browser and you will see the new version. Feel free to
use `kubectl apply` to switch between the two version and observe.


## Cleanup

Deletes everything created in this Lab

```
kubectl delete pod,rc,svc,deployment -l app=lobsters
```
