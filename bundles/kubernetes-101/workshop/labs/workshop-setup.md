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

```bash
mkdir -p ~/go/src/github.com/askcarter
cd ~/go/src/github.com/askcarter
git clone https://github.com/askcarter/workshop-in-a-box.git
cd workshop-in-a-box/kubernetes-101
```
