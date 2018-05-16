# Advanced Concepts Local

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster

## Lab

For this module, we will be interacting with the frontend pods. You
can deploy the database and use it for the whole module:

Create the password secret
```
kubectl create secret generic db-pass --from-literal=password=supersecret
```
or
```
kubectl create secret generic db-pass --from-file=password=pass.txt
```
```
secret "db-pass" created
```

Have the admin provision the persistent volumes for the database.

Deploy the database

```
kubectl create -f database-pvc.yaml
```

```
service "lobsters-sql" created
persistentvolumeclaim "mysql-pv-claim" created
deployment "lobsters-sql" created
```

Use previous commands you've learned to check the status of the
database (`describe pod`, `logs`)

With a fresh database, we need to run the rake commands to setup the database:

```
kubectl create -f ./rake-db.yaml
```
```
job "lobsters-rake" created
```

Check that it's done:

```
kubectl describe job lobsters-rake
```
```
Name:		lobsters-rake
Namespace:	default
Image(s):	gcr.io/google-samples/lobsters-db:1.0
Selector:	controller-uid=a61b758c-16c6-11e6-9e5e-42010af001a5
Parallelism:	1
Completions:	1
Start Time:	Tue, 10 May 2016 08:48:32 -0700
Labels:		app=lobsters,tier=rake
Pods Statuses:	0 Running / 1 Succeeded / 0 Failed
No volumes.
Events:
  FirstSeen	LastSeen	Count	From			SubobjectPath	Type		Reason			Message
  ---------	--------	-----	----			-------------	--------	------			-------
  1m		1m		1	{job-controller }			Normal		SuccessfulCreate	Created pod: lobsters-rake-g8nzr
```

Also, the frontend service can be created and used for the whole module:

```
kubectl create -f ./service-local.yaml
```
```
service "lobsters" created
```

Again, use commands you've learned previously to find the node port,
and access the site through `localhoast` (linux) or your Docker Machine
VM IP (Mac/Win).

### Advanced Pod Patters

Pods can contain multiple containers. Here, the frontend Rails app is
sharing it's log directory with a sidecar logger container. The
sidecar logger will send logs to Logstash.

> Note: our Logstash pod is not putting the logs in Kibana or anywhere else.

Deploy logstash:

```
kubectl create -f ./logstash.yaml
```
```
service "logstash" created
deployment "logstash" created
```

Inspect [frontend-sidecar.yaml](frontend-sidecar.yaml), and view the
`containers:` array. There are two containers and they both share the
volume `logdir` which is defined further up. The `logdir` volume is an
`emptyDir` type. This is ephemeral and will only live as long as the
pod. Deploy the frontend:

```
kubectl create -f ./frontend-sidecar.yaml
```
```
deployment "lobsters" created
```

The code for the logger container is under the [filebeat/](filebeat/)
directory.

Now without modifying our Lobsters app or container configuration, we
added a sidecar container to the pod and are aggregating the logs in
logstash. Check the name of the logstash container and view its stdout
to see the aggregated logs.

```
kubectl get pods
```
```
NAME                            READY     STATUS    RESTARTS   AGE
lobsters-1346140349-ok74t       2/2       Running   0          7m
lobsters-sql-3710543743-d7unn   1/1       Running   0          1h
logstash-1334877011-jhzbg       1/1       Running   0          1h
```

This one is `logstash-1334877011-jhzbg`, yours will be different.

```
kubectl logs logstash-1334877011-jhzbg
```
```
<Rails Logs>
```

### Canary, Rollback

We can create a second Deployment, and using the power of labels, the
frontend service will route traffic to both deployments. The
[frontend-beta.yaml](frontend-beta.yaml) config will create a second
deployment using the version `2.0` of the Lobsters container. Notice
the `channel: beta` label, so that the Deployments (and ReplicaSets)
can separately manage their own replicas.

```
kubectl create -f ./frontend-beta.yaml
```
```
deployment "lobsters-beta" created
```

The Service is still using the selector `app=lobsters,tier=frontend`,
without the `channel` key. This will match pods from both Deployments.

Now, due to the service's even load balancing and that we have 5
stable frontend pods and 2 beta frontend pods, we have the following
traffic split:

```
5/(5+2) = 71% stable
2/(5+2) = 29% beta
```

Visit the Lobsters site and refresh a few times, how often do you see
"Example Lobsters 2.0!!!"?

To raise the split to 50/50, scale one or the other. We'll scale the stable down to 2

```
kubectl scale deployment lobsters --replicas=2
```
```
deployment "lobsters" scaled
```
...
```
kubectl get deployments
```
```
NAME            DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
lobsters        2         2         2            2           24m
lobsters-beta   2         2         2            2           32m
lobsters-sql    1         1         1            1           50m
logstash        1         1         1            1           43m
```

Now that we have determined 2.0 is a good Lobsters version, lets
update the stable channel to 2.0. Edit `frontend-sidecar.yaml` in your
favorite editor and change the image from version 1.0 to 2.0: `image:
gcr.io/google-samples/lobsters-db:2.0` Now apply the new version:

```
kubectl apply -f ./frontend-sidecar.yaml
```
```
deployment "lobsters" configured
```

Now we can remove the beta deployment:

```
kubectl delete deployment lobsters-beta
```
```
deployment "lobsters-beta" deleted
```

Check:
```
kubectl get deployments
```
```
NAME           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
lobsters       5         5         5            5           30m
lobsters-sql   1         1         1            1           56m
logstash       1         1         1            1           50m
```

Lobsters is back up to 5 replicas, as that is what we had specified in
our frontend-sidecar.yaml. Visit the site, and you will only see
version 2.0 now.

Now let's say you found a critical bug in version 2.0. Kubernetes
deployments save their revision history so rollback is easy. You can
check the revisions with:

```
kubectl rollout history deployment lobsters
```
```
deployments "lobsters":
REVISION	CHANGE-CAUSE
1		<none>
2		<none>
```

And check each versions' config:

```
kubectl rollout history deployment lobsters --revision=1
```
```
deployments "lobsters" revision 1
Labels:		app=lobsters,channel=stable,pod-template-hash=1346140349,tier=frontend
Annotations:	kubernetes.io/change-cause=
Image(s):	gcr.io/google-samples/lobsters-db:1.0,gcr.io/google-samples/k8s-filebeat:1.0
Volumes:
  logdir:
    Type:	EmptyDir (a temporary directory that shares a pod's lifetime)
    Medium:
```

Ah yes, revision 1 is the one we want. You can specify it with the
`--to-revision` option, or just leave it off to roll back to the
previous version:

```
kubectl rollout undo deployment lobsters
```
```
deployment "lobsters" rolled back
```

Now visit the site and we are back to Lobsters version 1.0.

### Auto-scale

Horizontal pod autoscaling requires the Heapster and Metrics server cluster add-on.
Those were already setup in the clster. The below docs describe how autoscaling works, and how to configure Heapster:
* http://kubernetes.io/docs/user-guide/horizontal-pod-autoscaling/
* https://github.com/kubernetes/kubernetes/blob/release-1.2/docs/design/horizontal-pod-autoscaler.md#autoscaling-algorithm
* https://github.com/kubernetes/heapster/blob/master/docs/influxdb.md

Autoscaling requires resource requests and limits to be set:

```
kubectl appy -f frontend-sidecar-with-resources.yaml
```

Once everything is setup, pod auto-scaling is as simple as:

```
kubectl autoscale deployment lobsters --min=1 --max=6
```
```
deployment "lobsters" autoscaled
```

Now in a separate terminal window, watch the number of pods with:

```
kubectl get pods -w
```

Use a bash shell on your local machine or Docker VM to generate load:

```
while true ; do curl http://cjdv-k8-master.ep.esp.local:<node-port>/; done
```

Watch your other terminal with the pod list, the auto-scaler checks
every 30s and scales appropriately.

Hit `Ctrl-C` to stop the load.

## Cleanup

```
kubectl delete deployments,jobs,svc,pvc -l app=lobsters
```
