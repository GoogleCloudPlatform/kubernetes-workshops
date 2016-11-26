## Workshop setup

#### Google Cloud Platform account setup
* Navigate to https://console.cloud.google.com and login with your credentials.
* Select your project from the project listing.
* Navigate (using the menu or the search bar) to [Compute Engine](https://console.cloud.google.com/compute/).
* Enable Compute Engine (this may take a few minutes).
* Click the following button to activate [Cloud Shell](https://cloud.google.com/shell/docs),
which is your "command line in the cloud" and will be used complete the labs. ![Cloud Shell Icon](https://cloud.google.com/shell/docs/images/shell_icon.png)
* Open the menu icon on the top left.                                             
![Menu](https://codelabs.developers.google.com/codelabs/cloud-speech-intro/img/742dc285f86cdd1f.png)
* Select **API Manager** from the drop down.      
![API-Manager](https://codelabs.developers.google.com/codelabs/cloud-speech-intro/img/4cafd05ec8d75ebf.png).
* select **ENABLE API**.  
 ![Enable-API](https://codelabs.developers.google.com/codelabs/cloud-speech-intro/img/24185da15bfb437f.png)
* Under Google Cloud APIs, click on **Container Engine API**. If you need help finding the API, use the search field.
* Click **ENABLE**.  
![Enable](https://codelabs.developers.google.com/codelabs/cloud-speech-intro/img/985398850889c886.png)
* Wait for a few seconds for it to enable.

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
mkdir -p ~/go/src/github.com/GoogleCloudPlatform
cd ~/go/src/github.com/GoogleCloudPlatform
git clone https://github.com/GoogleCloudPlatform/kubernetes-workshops.git
cd kubernetes-workshops/bundles/kubernetes-101/workshop
```
