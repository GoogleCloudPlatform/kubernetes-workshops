# Kubernetes 101 Content Bundle

In this workshop you will learn how to:

* Provision Kubernetes using [Google Container Engine](https://cloud.google.com/container-engine)
* Deploy and manage Docker containers using kubectl

Kubernetes Version: v1.2.2

All of the code for this workshop was written by [Kelsey Hightower](https://twitter.com/kelseyhightower).

## Formats

### Course with Codelab

| Resource  | duration | description | 
| --- | --- | --- |
| [Intro video](https://www.youtube.com/watch?v=T59RtLov9E0) | 5m | Course intro video. |
| [Scalable Microservices with Kubernetes](https://www.udacity.com/course/scalable-microservices-with-kubernetes--ud615) | 180m | This is a detailed introductory Kubernetes course with industry experts like Kelsey Hightower, Adrian Cockcroft, and Carter Morgan. Just play the videos and follow along. |

### Kubernetes 101 Talk with Codelab, and Slides with speaker notes

| Resource  | duration | description | 
| --- | --- | --- |
| [Talk](https://www.youtube.com/watch?v=21hXNReWsUU) | 30m | Course intro video. |
| [Codelab](https://codelabs-preview.appspot.com/?file_id=13RVMEz5EWmG6-2ZeQIR_K14LWVipqfM-Bjex2wcFdP4#0) | 180m | Step by step codelab to walk through. |
| [Slides](https://docs.google.com/presentation/d/13SsyxNXnb2pB05LOdjtgBNjARD_qw9Dl0FLZeAlQbKA/edit?usp=sharing) | n/a | Slides that go along with the talk. |

### Kubernetes 101 Workshop with Slides and Labs

| Resource  | duration | description | 
| --- | --- | --- |
| [Slides](https://docs.google.com/presentation/d/1n3avmL5GCYCYJEr8pLFBKe0wzvoOiUV2vxyW_pYFL5s/edit?usp=sharing) | n/a | Slides with Speakernotes. |
| Workshop  | 180m | Kelsey Hightower's craftcon workshop -- the basis for the above talk and course.  The labs are below.  |

## Labs

Kubernetes is all about applications and in this section you will utilize the Kubernetes API to deploy, manage, and upgrade applications. In this part of the workshop you will use an example application called "app" to complete the labs.

* [Workshop Setup](labs/workshop-setup.md)
* [Containerizing your application](labs/containerizing-your-application.md)
* [Creating and managing pods](labs/creating-and-managing-pods.md)
* [Monitoring and health checks](labs/monitoring-and-health-checks.md)
* [Managing application configurations and secrets](labs/managing-application-configurations-and-secrets.md)
* [Creating and managing services](labs/creating-and-managing-services.md)
* [Creating and managing deployments](labs/creating-and-managing-deployments.md)
* [Rolling out updates](labs/rolling-out-updates.md)

## Lab Docker images

App is an example 12 Factor application. During this workshop you will be working with the following Docker images:

* [askcarter/monolith](https://hub.docker.com/r/askcarter/monolith) - Monolith includes auth and hello services.
* [askcarter/auth](https://hub.docker.com/r/askcarter/auth) - Auth microservice. Generates JWT tokens for authenticated users.
* [askcarter/hello](https://hub.docker.com/r/askcarter/hello) - Hello microservice. Greets authenticated users.
* [ngnix](https://hub.docker.com/_/nginx) - Frontend to the auth and hello services.

## Links

  * [Kubernetes](https://www.kubernetes.io)
  * [Docker](https://docs.docker.com)
  * [etcd](https://coreos.com/docs/distributed-configuration/getting-started-with-etcd)
  * [nginx](http://nginx.org)
