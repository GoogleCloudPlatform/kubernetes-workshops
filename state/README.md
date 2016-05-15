# Storing State

## Local Docker

This document is for cloud, for local docker see [local.md](local.md).

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster
* Commands assume that you have a local copy of this git repository, 
  and `state` is the current directory.

## Lab

Create a Service for your app. This will be used for the entire Lab.

```
kubectl create -f ./service-cloud.yaml
```
```
service "lobsters" created
```

Use the commands you have learned in previous modules to find the external IP.

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

### External Database

An easy way to move an existing app to Kubernetes, is to leave the
database where it is. To try this out, we can start up a Lobsters pod
that talks to an external MySQL database.

We have a new container
`gcr.io/google-samples/lobsters-db:1.0`. Source is under
[lobsters-db/](lobsters-db/). This container is configured to use the
environment variables `DB_HOST` and `DB_PASSWORD` to connect to a
MySQL database server. The username and port are hard coded to "root"
and 3306.

Edit [frontend-external.yaml](frontend-external.yaml) in your favorite
editor. Notice how we are setting the environment variables in the
`env:` section. Change the value of the `DB_HOST` variable to an
existing MySQL server, and save the file. Leave the configuration of
the password variable alone, It is set up to use a Kubernetes Secret
to retrieve the password. This is more secure than specifying the
password in the Deployment configuration.

Create the Secret using the below command. Change `mypassword` to be
the actual password to the external MySQL server.

```
kubectl create secret generic db-pass --from-literal=password=mypassword
```
```
secret "db-pass" created
```

> An alternate way to create the secret is to put the password in a
> file. We'll use `pass.txt`. Caution: your editor may put a trailing
> newline at the end of the file, which will not create the correct
> password.
> 
> ```
> kubectl create secret generic db-pass --from-file=password=pass.txt
> ```
> ```
> secret "db-pass" created
> ```

Now you can create the Lobsters deployment that connects to the
external MySQL server:

```
kubectl create -f ./frontend-external.yaml
```
```
deployment "lobsters" created
```

About Rails: it needs the database set up using a few `rake`
commands. These need to be run using the app code. To do this, we will
exec a bash shell inside one of our frontend pods. First, find the
name of any one of your frontend pods:

```
kubectl get pods
```
```
NAME                            READY     STATUS    RESTARTS   AGE
lobsters-3566082729-3j2mv       1/1       Running   0          3m
lobsters-3566082729-4wup5       1/1       Running   0          3m
lobsters-3566082729-8cxp0       1/1       Running   0          3m
lobsters-3566082729-add2d       1/1       Running   0          3m
lobsters-3566082729-sepki       1/1       Running   0          3m
```

Then, exec the shell.

```
kubectl exec -it lobsters-3566082729-3j2mv -- /bin/bash
```
```
root@lobsters-3566082729-3j2mv:/app#
```

Now inside the `/app` dir, run the rake command:

```
bundle exec rake db:create db:schema:load db:seed
```
```
<db output>
```

Then exit the shell.

Check the site, now you can log in and create stories, and no matter
which replica you visit, you'll see the same data.


Delete the deployment to move to the next step:

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
any Pod in the cluster using Kube DNS. Deploy the database:

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

Delete the database and service, leave the frontend running for the next step:

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
we will use a Persistent Volume, which is a Kubernetes object that
refers to an external storage device. Clouds provide resilient disks
that can be easily be attached and detached to cluster nodes.

> Note: For AWS see the docs to create an EBS disk and
> PersistentVolume object:
> http://kubernetes.io/docs/user-guide/volumes/#awselasticblockstore
> http://kubernetes.io/docs/user-guide/persistent-volumes/#persistent-volumes

Crate a cloud disk, make sure you use the same zone as your cluster.

```
gcloud compute disks create mysql-disk --size 20GiB
```
```
Created [https://www.googleapis.com/compute/v1/projects/myproject/zones/us-central1-c/disks/mysql-disk].
NAME          ZONE           SIZE_GB  TYPE         STATUS
mysql-disk    us-central1-c  20       pd-standard  READY
```

The [cloud-pv.yaml](cloud-pv.yaml) config will create a Persistent
Volume object of the `gcePersistentDisk` type that refers to the
`mysql-disk' you created above.

```
kubectl create -f ./cloud-pv.yaml
```
```
persistentvolume "pv-1" created
```
...
```
kubectl get pv
```
```
NAME      CAPACITY   ACCESSMODES   STATUS      CLAIM     REASON    AGE
pv-1      20Gi       RWO           Available                       6s
```
...
```
kubectl describe pv pv-1
```
```
Name:		pv-1
Labels:		<none>
Status:		Available
Claim:		
Reclaim Policy:	Retain
Access Modes:	RWO
Capacity:	20Gi
Message:	
Source:
    Type:	GCEPersistentDisk (a Persistent Disk resource in Google Compute Engine)
    PDName:	mysql-disk
    FSType:	ext4
    Partition:	0
    ReadOnly:	false
```

Now take a look at [database-pvc.yaml](database-pvc.yaml). This has a
new Persistent Volume Claim (PVC) object in it. A PVC will claim an
existing PV in the cluster that meets its requirements. The advantage
here is that our database config is independent of the cluster
environment. If we used cloud disks for PVs in our cloud cluster, and
iSCSI PVs in our on-prem cluster, the database config would be the
same.

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
kubectl get pv
```
```
NAME      CAPACITY   ACCESSMODES   STATUS    CLAIM                    REASON    AGE
pv-1      20Gi       RWO           Bound     default/mysql-pv-claim             11m
```
...
```
kubectl get pvc
```
```
NAME             STATUS    VOLUME    CAPACITY   ACCESSMODES   AGE
mysql-pv-claim   Bound     pv-1      20Gi       RWO           4m
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
NAME      CAPACITY   ACCESSMODES   STATUS     CLAIM                    REASON    AGE
pv-1      20Gi       RWO           Released   default/mysql-pv-claim             22m
```

You can see it is now released.

```
kubectl delete pv pv-1
```
```
persistentvolume "pv-1" deleted
```
...
```
kubectl delete secret db-pass
```
```
secret "db-pass" deleted
```

Save your cloud disk for the next lab, but when you are ready to delete it:
```
gcloud compute disks delete mysql-disk
```
```
The following disks will be deleted. Deleting a disk is irreversible
and any data on the disk will be lost.
 - [mysql-disk] in [us-central1-c]

Do you want to continue (Y/n)?  y

Deleted [https://www.googleapis.com/compute/v1/projects/myproject/zones/us-central1-c/disks/mysql-disk].
```
