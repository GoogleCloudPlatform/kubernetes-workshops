# Kubernetes Quickstart

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster

## Lab

Run the Lobsters app container

<!-- START Bash -->
```
kubectl run gcr.io/google-samples/lobsters:latest
```
<!-- END Bash -->

TODO: more stuff with pod

## Cleanup

Delete your container

<!-- START Bash -->
```
kubectl delete lobsters
```
<!-- END Bash -->
