[![Build Status](https://travis-ci.org/GoogleCloudPlatform/kubernetes-workshops.svg?branch=master)](https://travis-ci.org/GoogleCloudPlatform/kubernetes-workshops)

# Kubernetes Workshops

This repository contains both complete workshops, as well as
segmented workshop modules that can be combined to create Kubernetes
workshops of various lengths and focus.

Modules contain a README.md that walks through the module. If code,
configuration, or scrips are needed, it is included and tested.

This is not an official Google product.

## Full Workshops
| workshop  | version | description |
| --- | --- | --- |
| [Kubernetes 101](bundles/kubernetes-101) | v1.2.0 | Covers the basics of using Kubernetes to manage applications at scale.  In this workshop, you'll take an app, build it into a docker container, then use Kubernetes to deploy, scale, and update it. This workshop comes in multiple versions:  A video course, a codelab with an accompanying talk, or a set of workshop material with slides. |


## Individual Modules

Name | Slides | Level | Time Estimate | Completion Status
------------- | ------------- | ------------- | ------------ | ------------
[Cluster Bring Up](bring-up) | [Link](https://docs.google.com/presentation/d/1AZSJi4wl1ALfMNuW8X2hoN6DZZVs35dl5lPDtDuWT_U/edit?usp=sharing) | Beginner | 1 hour | Ready
[Quickstart](quickstart) | [Link](https://docs.google.com/presentation/d/1nH88mgUhcGtuyCD9W1k3blvxrTkcqC3FfcRkiX10JR8/edit?usp=sharing) | Beginner | 1 hour | Ready
[Core Kubernetes Concepts](core-concepts) | [Link](https://docs.google.com/presentation/d/1JP6-utzrocigFpVyd9IFoZmjPV5vxGhiOONT36XBF_o/edit?usp=sharing) | Beginner | 4 hours | Ready
[Storing State](state) | [Link](https://docs.google.com/presentation/d/1av0gZl90NS2oPm2u5utht7fZums76fvsh-hqt_JMzJE/edit?usp=sharing) | Intermediate | 2 hours | Ready
[Advanced Concepts](advanced) | [Link](https://docs.google.com/presentation/d/1_mWY3fTavAYjD9twABOOEE0dRVz8YeMpeJ0s1GgvHpk/edit?usp=sharing) | Intermediate | 2 hours | Ready
[Dockerize an App](dockerize) | [Link]() | Intermediate | 2 hours | Defer
[Networking](networking) | [Link]() | Intermediate | 2 hours | Defer
[Troubleshooting](troubleshooting) | [Link]() | Intermediate | 2 hours | Delay
[Putting it all together](combine) | [Link]() | Advanced | 2 hours | Delay

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
* move everything under a deployment
* update to new versions of our app, quick rolling update
  * lightweight here - more detail in "Advanced" module

### Storing State

* Deploy an app with MySQL
* multiple iterations where to store the data, how it goes away
  * start with host volume, end at persistent disk
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
  see how we get into the cluster
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
