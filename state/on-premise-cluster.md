# Storing State

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster
* Commands assume that you have a local copy of this git repository,
  and `state` is the current directory.


## Lab

Create a Service for your app. This will be used for the entire Lab.

```
kubectl create -f ./service-local.yaml
```
```
service "lobsters" created
```

Use the commands you have learned in previous modules to find the node
port and Docker VM IP.

### No Database

Deploy the app that uses a local sqlite file:

```
kubectl create -f ./dep.yaml
```
```
deployment "lobsters" created
```

Because [dep.yaml](dep.yaml) specifies `replicas: 5`, You have 5
Lobsters pods running. The problem is that each one has it's own
separate database. The app and container was not designed to scale
with replicas.

Try visiting the site and logging in (test/test), add a story. Refresh
and it might disappear!

Delete just the deployment:

```
kubectl delete deployment lobsters
```
```
deployment "lobsters" deleted
```

### Run MySQL in Kubernetes

Look through [database.yaml](database.yaml). This config creates a
MySQL Deployment and Service. It uses the standard
[MySQL container](https://hub.docker.com/_/mysql/), and specifies the
password to the container using the `MYSQL_ROOT_PASSWORD` environment
variable, sourcing the value from the Secret you defined above. The
Service name is `lobsters-sql` and will be resolvable as a hostname to
any Pod in the cluster using Kube DNS.

Create the password:
```
kubectl create secret generic db-pass --from-literal=password=mypassword
```
```
secret "db-pass" created
```

Deploy the database:

```
kubectl create -f ./database.yaml
```
```
service "lobsters-sql" created
deployment "lobsters-sql" created
```

We have [frontend-dep.yaml](frontend-dep.yaml) set up to connect to
the database using the `lobsters-sql` hostname of the database
Service. Because the password is sourced from the same Secret, they
will always be in sync. Deploy the frontend:

```
kubectl create -f ./frontend-dep.yaml
```
```
deployment "lobsters" created
```

This time, to run the rake commands we will re-use the `lobsters-db`
container image, but specify an alternate command in the Kubernetes
config. Because we only want this command to run once and exit, we use
the Kubernetes Job object. See [rake-db.yaml](rake-db.yaml) for the
configuration.

```
kubectl create -f ./rake-db.yaml
```
```
job "lobsters-rake" created
```

Check the Job status:

```
kubectl describe job lobsters-rake
```
```
Name:		lobsters-rake
Namespace:	default
Image(s):	gcr.io/google-samples/lobsters-db:1.0
Selector:	controller-uid=f1d48c99-186b-11e6-9e5e-42010af001a5
Parallelism:	1
Completions:	1
Start Time:	Thu, 12 May 2016 11:04:17 -0700
Labels:		app=lobsters,tier=rake
Pods Statuses:	0 Running / 1 Succeeded / 0 Failed
No volumes.
Events:
  FirstSeen	LastSeen	Count	From			SubobjectPath	Type		Reason			Message
  ---------	--------	-----	----			-------------	--------	------			-------
  38s		38s		1	{job-controller }			Normal		SuccessfulCreate	Created pod: lobsters-rake-w112p
```

Succeeded! Now check the app and verify it is working fine.

We still have a problem with this configuration. We are running a
database in Kubernetes, but we have no data persistence. The data is
being stored inside the container, which is inside the Pod. This is
fine as long as the container and Node continue to run. Kubernetes
likes to consider Pods ephemeral and stateless. If the Node were to
crash, or the Pod be deleted, Kubernetes will reschedule a new Pod for
the lobsters-sql Deployment, but the data would be lost.

Delete the database and service, leave the frontend running for the
next step:

```
kubectl delete deployment,svc lobsters-sql
```
```
deployment "lobsters-sql" deleted
service "lobsters-sql" deleted
```

Also delete the Job.

```
kubectl delete job lobsters-rake
```
```
job "lobsters-rake" deleted
```

### Database with Persistent Volume

To run the database inside Kubernetes, but keep the data persistent,
we will use a Persistent Volume (PV), which is a Kubernetes object
that usually refers to an external storage device. This is typically a
resilient cloud disk in clouds, or an NFS or iSCSI disk in on-premise
clusters.

For running locally on your computer, we will use a `hostPath` type of
PV. This maps to a directory on the node. This is only as resilient as
the node, and would not work on a multi-node cluster.

Have the admin create the Persistent Volumes for the cluster. This will
create multiple volumes.

```
kubectl create -f ../local-pvs.yaml
```
```
persistentvolume "db-local-pv-1" created
persistentvolume "db-local-pv-2" created
persistentvolume "db-local-pv-3" created
persistentvolume "db-local-pv-4" created
persistentvolume "db-local-pv-5" created
persistentvolume "db-local-pv-6" created
persistentvolume "db-local-pv-7" created
persistentvolume "db-local-pv-8" created
persistentvolume "db-local-pv-9" created
persistentvolume "db-local-pv-10" created
persistentvolume "db-local-pv-11" created
persistentvolume "db-local-pv-12" created
persistentvolume "db-local-pv-13" created
persistentvolume "db-local-pv-14" created
persistentvolume "db-local-pv-15" created
```
...
```
kubectl get pv
```
```
NAME             CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM     STORAGECLASS   REASON    AGE
db-local-pv-1    20Gi       RWO            Delete           Available                                      1m
db-local-pv-10   20Gi       RWO            Delete           Available                                      1m
db-local-pv-11   20Gi       RWO            Delete           Available                                      1m
db-local-pv-12   20Gi       RWO            Delete           Available                                      1m
db-local-pv-13   20Gi       RWO            Delete           Available                                      1m
db-local-pv-14   20Gi       RWO            Delete           Available                                      1m
db-local-pv-15   20Gi       RWO            Delete           Available                                      1m
db-local-pv-2    20Gi       RWO            Delete           Available                                      1m
db-local-pv-3    20Gi       RWO            Delete           Available                                      1m
db-local-pv-4    20Gi       RWO            Delete           Available                                      1m
db-local-pv-5    20Gi       RWO            Delete           Available                                      1m
db-local-pv-6    20Gi       RWO            Delete           Available                                      1m
db-local-pv-7    20Gi       RWO            Delete           Available                                      1m
db-local-pv-8    20Gi       RWO            Delete           Available                                      1m
db-local-pv-9    20Gi       RWO            Delete           Available                                      1m

```
...
```
kubectl describe pv db-local-pv-1
```
```
Name:            db-local-pv-1
Labels:          type=local
Annotations:     <none>
Finalizers:      []
StorageClass:
Status:          Available
Claim:
Reclaim Policy:  Delete
Access Modes:    RWO
Capacity:        20Gi
Node Affinity:   <none>
Message:
Source:
    Type:          HostPath (bare host directory volume)
    Path:          /tmp/mysql-disk-1
    HostPathType:
Events:            <none>
```

Now take a look at [database-pvc.yaml](database-pvc.yaml). This has a
new Persistent Volume Claim (PVC) object in it. A PVC will claim an
existing PV in the cluster that meets its requirements. The advantage
here is that our database config is independent of the cluster
environment. This PVC can be used to claim one of the local PV we have, and
also a PV in a cloud Kubernetes cluster.

The PVC is named `mysql-pv-claim` and is then referenced in the Pod
specification and mounted to `/var/lib/mysql`. Deploy the new database:

```
kubectl create -f ./database-pvc.yaml
```
```
service "lobsters-sql" created
persistentvolumeclaim "mysql-pv-claim" created
deployment "lobsters-sql" created
```

Now you can see that the PVC is bound to the PV:

```
kubectl get pvc
```
```
NAME             STATUS    VOLUME          CAPACITY   ACCESS MODES   STORAGECLASS   AGE
mysql-pv-claim   Bound     db-local-pv-2   20Gi       RWO                           22s
```
Depending on which persistent volume was assigned to you, issue the command:
...
```
kubectl get pv {your-assigned-pv}
```
```
NAME            CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS    CLAIM                 STORAGECLASS   REASON    AGE
db-local-pv-2   20Gi       RWO            Delete           Bound     ns15/mysql-pv-claim                            5m
```

Re run the Job, as we have a new DB:

```
kubectl create -f ./rake-db.yaml
```
```
job "lobsters-rake" created
```

Now visit the site and make sure it works.

## Cleanup

```
kubectl delete svc,deployment,job,pvc -l app=lobsters
```
```
service "lobsters" deleted
service "lobsters-sql" deleted
deployment "lobsters" deleted
deployment "lobsters-sql" deleted
job "lobsters-rake" deleted
persistentvolumeclaim "mysql-pv-claim" deleted
```

We didn't label the PV, as it is of general use.

```
kubectl get pv
```
```
NAME             CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM     STORAGECLASS   REASON    AGE
db-local-pv-1    20Gi       RWO            Delete           Available                                      6m
db-local-pv-10   20Gi       RWO            Delete           Available                                      6m
db-local-pv-11   20Gi       RWO            Delete           Available                                      6m
db-local-pv-12   20Gi       RWO            Delete           Available                                      6m
db-local-pv-13   20Gi       RWO            Delete           Available                                      6m
db-local-pv-14   20Gi       RWO            Delete           Available                                      6m
db-local-pv-15   20Gi       RWO            Delete           Available                                      6m
db-local-pv-3    20Gi       RWO            Delete           Available                                      6m
db-local-pv-4    20Gi       RWO            Delete           Available                                      6m
db-local-pv-5    20Gi       RWO            Delete           Available                                      6m
db-local-pv-6    20Gi       RWO            Delete           Available                                      6m
db-local-pv-7    20Gi       RWO            Delete           Available                                      6m
db-local-pv-8    20Gi       RWO            Delete           Available                                      6m
db-local-pv-9    20Gi       RWO            Delete           Available                                      6m
```

You can see it is now released.

...
```
kubectl delete secret db-pass
```
```
secret "db-pass" deleted
```
