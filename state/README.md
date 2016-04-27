# Storing State

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster

## Lab

Rough notes here:

---

(skip this step if coming from core concepts)
From core concepts, uses a sqlite file on the local filesystem inside the container:
either service-cloud.yaml or service-local.yaml
dep.yaml

ok, delete now, can leave service for next step.

---

Lets move to a proper database:

kubectl create secret generic db-pass --from-literal=password=supersecret
or
kubectl create secret generic db-pass --from-file=password=pass.txt

deploy:
database.yaml

run rake:
rake-db.yaml

deploy:
service-cloud.yaml or service-local.yaml

frontend-dep.yaml

creates a two tier repilcated frontend with single mysql backend,

still same problem! data is stored in the db's ephemeral container

----

now delete the database, and deploy with a persistant disk

(for gce: gcloud compute disks create mysql-disk --size 20GiB)

either local-pv.yaml or cloud-pv.yaml

database-pvc.yaml

then re-run rake: rake-db.yaml (delete old job first)

---
end..

## Cleanup
