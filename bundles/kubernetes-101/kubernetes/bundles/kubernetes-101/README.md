# Kubernetes workshop

In this workshop you will learn how to:

* Provision Kubernetes using [Google Container Engine](https://cloud.google.com/container-engine)
* Deploy and manage Docker containers using kubectl

Kubernetes Version: v1.2.2

All of the code for this workshop was written by Kelsey Hightower.

## Workshop setup

#### Google Cloud Platform account setup
* Navigate to https://console.cloud.google.com and login with your credentials.
* Select your project from the project listing.
* Navigate (using the menu or the search bar) to [Compute Engine](https://console.cloud.google.com/compute/).
* Enable Compute Engine (this may take a few minutes).
* Click the following button to activate [Cloud Shell](https://cloud.google.com/shell/docs),
which is your "command line in the cloud" and will be used complete the labs. ![Cloud Shell Icon](https://cloud.google.com/shell/docs/images/shell_icon.png)

#### Provision Kubernetes using Google Container Engine (GKE)

Kubernetes can be configured with many options and add-ons, but can be time consuming to bootstrap from the ground up. In this section you will bootstrap Kubernetes using [Google Container Engine](https://cloud.google.com/container-engine) (GKE).

In your **Cloud Shell** terminal, issue the following commands (feel free to change the zone or cluster name):

```
gcloud config set compute/zone europe-west1-b
gcloud container clusters create myk8scluster --num-nodes 7
```

#### Clone repository

In your Cloud Shell environment clone the following repository.

```
git clone https://github.com/askcarter/workshop-in-a-box.git
cd workshop-in-a-box/kubernetes101/kubernetes
```

## Labs

Kubernetes is all about applications and in this section you will utilize the Kubernetes API to deploy, manage, and upgrade applications. In this part of the workshop you will use an example application called "app" to complete the labs.

* [Containerizing your application](labs/containerizing-your-application.md)
* [Creating and managing pods](labs/creating-and-managing-pods.md)
* [Monitoring and health checks](labs/monitoring-and-health-checks.md)
* [Managing application configurations and secrets](labs/managing-application-configurations-and-secrets.md)
* [Creating and managing services](labs/creating-and-managing-services.md)
* [Creating and managing deployments](labs/creating-and-managing-deployments.md)
* [Rolling out updates](labs/rolling-out-updates.md)

## Lab Docker images

App is an example 12 Facter application. During this workshop you will be working with the following Docker images:

* [askcarter/monolith](https://hub.docker.com/r/askcarter/monolith) - Monolith includes auth and hello services.
* [askcarter/auth](https://hub.docker.com/r/askcarter/auth) - Auth microservice. Generates JWT tokens for authenticated users.
* [askcarter/hello](https://hub.docker.com/r/askcarter/hello) - Hello microservice. Greets authenticated users.
* [ngnix](https://hub.docker.com/_/nginx) - Frontend to the auth and hello services.

## Links

  * [Kubernetes](http://googlecloudplatform.github.io/kubernetes)
  * [gcloud Tool Guide](https://cloud.google.com/sdk/gcloud)
  * [Docker](https://docs.docker.com)
  * [etcd](https://coreos.com/docs/distributed-configuration/getting-started-with-etcd)
  * [nginx](http://nginx.org)
