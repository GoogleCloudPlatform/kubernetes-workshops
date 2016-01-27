# Kubernetes Quickstart

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster

## Lab

Run the Lobsters app container

<!-- dont-START bash -->
```
kubectl run lobsters --image=gcr.io/google-samples/lobsters:latest
```
<!-- dont-END bash -->

<!-- Test this in the meantime START bash
kubectl run nginx --image=nginx
  END bash -->

Look at the pods
<!-- START bash -->
```
kubectl get pods
```
<!-- END bash -->

TODO: more stuff with pod

## Cleanup

Delete your container

<!-- dont-START bash -->
```
kubectl delete rc lobsters
```
<!-- dont-END bash -->
<!-- Test this in the meantime START bash
kubectl delete rc nginx
  END bash -->
