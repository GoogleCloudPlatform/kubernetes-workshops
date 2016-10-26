# Core Concepts

## Local Docker

This document is for cloud, for local docker see [local.md](local.md).

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster
* Commands assume that you have a local copy of this git repository, 
  and `core-concepts` is the current directory.

## Lab


### Pods

First define a single pod, see [pod.yaml](pod.yaml). Start the
Lobsters app from this pod declaration.

<!-- START bash -->
```
kubectl create -f ./pod.yaml
```
<!-- END bash -->

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
Look at the details

```
kubectl describe pod lobsters
```
```
Name:		lobsters
Namespace:	default
Node:		gke-myclus-default-pool-c6da95a3-ot2d/10.128.0.6
Start Time:	Sun, 15 May 2016 19:54:57 -0700
Labels:		app=lobsters
Status:		Running
IP:		10.0.2.4
Controllers:	<none>
Containers:
  lobsters:
    Container ID:	docker://e860dcd38e802d6b587a46ddfd21a750e2bc454cf0bf131ddaa2649d4d1957eb
    Image:		gcr.io/google-samples/lobsters:1.0
    Image ID:		docker://4e7ae90a0dcf3ca1903ead56199ed1839ed84373479d758c538b9028702672a1
    Port:		3000/TCP
    QoS Tier:
      cpu:	Burstable
      memory:	BestEffort
    Requests:
      cpu:		100m
    State:		Running
      Started:		Sun, 15 May 2016 19:54:58 -0700
    Ready:		True
    Restart Count:	0
    Environment Variables:
Conditions:
  Type		Status
  Ready 	True 
Volumes:
  default-token-oqs86:
    Type:	Secret (a volume populated by a Secret)
    SecretName:	default-token-oqs86
Events:
  FirstSeen	LastSeen	Count	From						SubobjectPath			Type		Reason		Message
  ---------	--------	-----	----						-------------			--------	------		-------
  19s		19s		1	{default-scheduler }								Normal		Scheduled	Successfully assigned lobsters to gke-myclus-default-pool-c6da95a3-ot2d
  18s		18s		1	{kubelet gke-myclus-default-pool-c6da95a3-ot2d}	spec.containers{lobsters}	Normal		Pulled		Container image "gcr.io/google-samples/lobsters:1.0" already present on machine
  18s		18s		1	{kubelet gke-myclus-default-pool-c6da95a3-ot2d}	spec.containers{lobsters}	Normal		Created		Created container with docker id e860dcd38e80
  18s		18s		1	{kubelet gke-myclus-default-pool-c6da95a3-ot2d}	spec.containers{lobsters}	Normal		Started		Started container with docker id e860dcd38e80
```

Delete the pod

<!-- START bash -->
```
kubectl delete pod lobsters
```
<!-- END bash -->
<!-- START bash
while kubectl get pod lobsters; do echo running; sleep 1; done
END bash -->

```
pod "lobsters" deleted
```

The pod is gone forever. The same would be true if the processes in 
the pod terminated for any reason.

### Service

To access Lobsters from outside the cluster, we'll need a service. The
service defined in [service.yaml](service.yaml) will route traffic to
any pod with the label `app: lobsters`, which matches our pod
definition. The service is for port 80, but routes to the port labeled
`web` in our pod definition. The `type: LoadBalancer` creates an IP
external to the cluster in supported environments.

Create the service and pod:

<!-- START bash -->
```
kubectl create -f ./service.yaml,./pod.yaml
```
<!-- END bash -->

```
service "lobsters" created
pod "lobsters" created
```

Get the external IP, this will take a minute. EXTERNAL-IP will be blank until the load balancer is ready.

<!-- START bash -->
```
kubectl get svc lobsters
```
<!-- END bash -->

```
NAME       CLUSTER-IP     EXTERNAL-IP    PORT(S)   AGE
lobsters   10.3.253.158   1.2.3.4        80/TCP    1m
```

Check that it is working by visiting the external IP in your browser.



Delete Pods and Services where the `app` label equals `lobsters`.

<!-- START bash -->
```
kubectl delete pod,svc -l app=lobsters
```
<!-- END bash -->
<!-- START bash
while kubectl get pod lobsters; do echo running; sleep 1; done
while kubectl get service lobsters; do echo running; sleep 1; done
END bash -->

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

<!-- START bash -->
```
kubectl create -f ./rc.yaml,./service.yaml
```
<!-- END bash -->

```
replicationcontroller "lobsters" created
service "lobsters" created
```

Wait for the external IP:

<!-- START bash -->
```
kubectl get svc lobsters
```
<!-- END bash -->

```
NAME       CLUSTER-IP     EXTERNAL-IP    PORT(S)   AGE
lobsters   10.3.253.158   1.2.3.4        80/TCP    1m
```

Check that it is working by visiting the external IP in your browser.


Now, look at the pod

```
kubectl get pods -o wide
```

```
NAME             READY     STATUS    RESTARTS   AGE       NODE
lobsters-jf0xs   1/1       Running   0          2m        gke-myclus-2f1fdf58-node-lfaa
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
NAME             READY     STATUS    RESTARTS   AGE       NODE
lobsters-t1vwk   1/1       Running   0          6s        gke-myclus-2f1fdf58-node-lfaa
```

A new pod was created! It might even be on a different node.

Scaling is as easy as:

<!-- START bash -->
```
kubectl scale --replicas=5 rc lobsters
```
<!-- END bash -->

```
replicationcontroller "lobsters" scaled
```

Check the pods

```
kubectl get pods -o wide
```

```
NAME             READY     STATUS              RESTARTS   AGE       NODE
lobsters-32ona   1/1       Running             0          26s       gke-myclus-2f1fdf58-node-lfaa
lobsters-8twm0   1/1       Running             0          2m        gke-myclus-2f1fdf58-node-lfaa
lobsters-hhves   0/1       ContainerCreating   0          26s       gke-myclus-2f1fdf58-node-kxe4
lobsters-lv5km   0/1       ContainerCreating   0          26s       gke-myclus-2f1fdf58-node-bvxp
lobsters-tlojp   0/1       ContainerCreating   0          26s       gke-myclus-2f1fdf58-node-bvxp
```

Also the RC

<!-- START bash -->
```
kubectl get rc lobsters -o wide
```
<!-- END bash -->

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

Delete Replication Controllers and Services where the `app` label equals `lobsters`.

<!-- START bash -->
```
kubectl delete rc,svc -l app=lobsters
```
<!-- END bash -->
<!-- START bash
while kubectl get rc lobsters; do echo running; sleep 1; done
while kubectl get service lobsters; do echo running; sleep 1; done
while [ -n "$(kubectl get pod -l app=lobsters)" ]; do echo running; done
END bash -->

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

<!-- START bash -->
```
kubectl create -f ./dep.yaml,./service.yaml
```
<!-- END bash -->

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

<!-- START bash -->
```
kubectl apply -f ./dep-2.yaml
```
<!-- END bash -->

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

Deletes everything created in this Lab: resources whose `app` label equals `lobsters`.


<!-- START bash -->
```
kubectl delete pod,rc,svc,deployment -l app=lobsters
```
<!-- END bash -->
<!-- START bash
while kubectl get deployment lobsters; do echo running; sleep 1; done
while kubectl get service lobsters; do echo running; sleep 1; done
while [ -n "$(kubectl get pod -l app=lobsters)" ]; do echo running; done
END bash -->
