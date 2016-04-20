[![Build Status](https://travis-ci.org/GoogleCloudPlatform/kubernetes-workshops.svg?branch=master)](https://travis-ci.org/GoogleCloudPlatform/kubernetes-workshops)

# Kubernetes Workshops

This repository contains various modules that can be combined to
create Kubernetes workshops of various lengths and focus. Each
directory is a module and contains a README.md that walks through the
module and gopresents slides that accompany the module. If code,
configuration, or scrips are needed, it is included and tested.

This is not an official Google product.

## Table of Modules

Name | Level | Time Estimate | Completion Status
------------- | ------------- | ------------ | ------------
[Quickstart](quickstart) | Beginner | 1 hour | Draft
[Cluster Bring Up](bring-up) | Beginner | 1 hour | In Progress (Jeff)
[Core Kubernetes Concepts](core-concepts) | Beginner | 4 hours | Draft
[Storing State](state) | Intermediate | 2 hours | Not Started
[Dockerize an App](dockerize) | Intermediate | 2 hours | Not Started
[Advanced Concepts](advanced) | Intermediate | 2 hours | Not Started
[Networking](networking) | Intermediate | 2 hours | Not Started
[Troubleshooting](troubleshooting) | Intermediate | 2 hours | Not Started
[Putting it all together](combine) | Advanced | 2 hours | Not Started

Status: Not Started --> In Progress --> Draft --> Ready

## Overview of Modules

> This is incomplete. These are just brainstorming / rough notes.

### Quickstart

* Quick
* Not complex, uses `kubectl run`, `kubectl expose`.
* Demonstrates the ease of using Kubernetes without learning all the
  concepts and config files up front.

### Cluster Bring Up

* Uses the open source kubernetes release with `cluster/kube-up.sh`
  for cloud bringup
* Option for local docker

### Core Kubernetes Concepts

* Introduce one concept at a time, and then use that concept
  * Order: pod, service, rc, deployment
* go over a declarative pod representation of quickstart app
  * contains 1 pod
* logs, exec, port forwarding
* introduce service
* overview pods, labels, selectors, and services
* change pod to RC
* discuss RC
* scale pod up
* introduce deployments
* move everything under a deplyoment
* update to new versions of our app, quick rolling update
  * lightweight here - more detail in "Advanced" module

### Storing State

* Deploy an app with MySQL
* multiple iterations where to store the data, how it goes away
  * start with host voulme, end at persistant disk
* More of a lecture module, slides discuss state in greater length

### Dockerize an App

* Start with an app
* Write the Dockerfile
* build
* push to registry

### Advanced Concepts

* Lecture
  * App/container patterns
    http://blog.kubernetes.io/2015/06/the-distributed-system-toolkit-patterns.html
  * mapping non-containerized apps

* Hands on:
  * A/B deployment
  * Canary patterns
  * Rolling Deployments
  * Autoscaling

### Networking

* More of a lecture module, slides discuss networking in greater length

* Types of external services VIP/nodeport, run service with each and
  see how we get into the cluister
* discuss subnets, explore on running nodes
* How K8s networking works
* Setting up an external load balancer - Nginx
* Ideas on how to plug into your environment

### Troubleshooting

* Logging & monitoring
* Troubleshooting / Debugging

### Putting it all together

* deploy a production ready app

* Use all the above
* Build up a significant realistic app
* ( not so much lecture, just deploy all this stuff: )
* web frontends, caching, backend jobs, datastore, load testing
* Logging & monitoring
* Troubleshooting
* Autoscaling

## Contributing changes

* See [CONTRIBUTING.md](CONTRIBUTING.md)

## Licensing

* See [LICENSE](LICENSE)
