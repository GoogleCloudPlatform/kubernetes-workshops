# Cloud Cluster Bring-Up

## Prerequisites

A cloud account with a supported cloud provider.

* [GCP](http://cloud.google.com/)
* AWS

## Lab

Download the `kubernetes.tar.gz` latest binary release:

```
curl -LO https://github.com/kubernetes/kubernetes/releases/download/v1.2.4/kubernetes.tar.gz
```

Extract

```
tar -xzvf kubernetes.tar.gz
cd kubernetes/
```

## Configure your cloud defaults

### GCP

Edit `cluster/gce/config-default.sh` in your favorite text editor. The
defaults are all fine, but you may wish to customize your cluster:

| Variable | Description |
| --- | --- |
| ZONE | Set this to whichever zone you wish. |
| NODE_SIZE | VM size of individual nodes. |
| MASTER_SIZE | VM size of the master. |
| NUM_NODES | Number of node VMs to create. |

> Note: These can be set by environment variable as well.

Also, specify that you want kubernetes to start on GCP, we'll use the
environment variable here:

```
export KUBERNETES_PROVIDER=gce
```

### AWS

Edit `cluster/aws/config-default.sh` in your favorite text editor. By
default, the script will provision a new VPC and a 4 node k8s cluster
in us-west-2a (Oregon) with EC2 instances running on Ubuntu.

Also, specify that you want kubernetes to start on AWS, we'll use the
environment variable here:

```
export KUBERNETES_PROVIDER=aws
```

## Configure your cloud-specific prerequisites

### GCP

1. Go to the
   [Google Cloud Platform Console](https://console.cloud.google.com/project/_/compute/instances?_ga=1.92147801.233469832.1449873262).
   When prompted, select an existing project or create a new project.
2. Follow the prompts to set up billing. If you are new to Google
   Cloud Platform, you have
   [free trial](https://cloud.google.com/free-trial/) credit to pay
   for your instances.

   > Note: If you are using Cloud Shell, you don't need to do the below steps.

3. Download and extract the
   [Google Cloud SDK](https://cloud.google.com/sdk/), which includes
   the `gcloud` command line tool. It will be extracted to the
   `google-cloud-sdk` directory.
4. Run the install script to add SDK tools to your path, enable
   command completion in your bash shell, and/or and enable usage
   reporting.

   ```
   ./google-cloud-sdk/install.sh
   ```

5. Run gcloud init to initialize the SDK:


   ```
   gcloud init
   ```

### AWS

1. You need an AWS account. Visit
   [http://aws.amazon.com](http://aws.amazon.com) to get started
2. Install and configure
   [AWS Command Line Interface](http://aws.amazon.com/cli)
3. You need an AWS
   [instance profile and role](http://docs.aws.amazon.com/IAM/latest/UserGuide/instance-profiles.html)
   with EC2 full access.

This script use the 'default' AWS profile by default.  You may
explicitly set AWS profile to use using the `AWS_DEFAULT_PROFILE`
environment variable:

```
export AWS_DEFAULT_PROFILE=myawsprofile
```

## Start the cluster

Now everything should be set up to start your cluster

```
cluster/kube-up.sh
```

Cluster bring-up can take a while, once all the VMs are created, the
script will wait for the cluster to start up and self provision, you
will see:

```
Waiting up to 300 seconds for cluster initialization.

  This will continually check to see if the API for kubernetes is reachable.
  This may time out if there was some uncaught error during start up.
```

And then, once successful communication and validation has taken place, look for:

```
Cluster validation succeeded
```

## Setup kubectl

Great! Now `kubectl` is configured with authentication to communicate
with your cluster. This is set up for you in your `~/.kube/config`
file. 

> Note: If you are using GCP and have `gcloud` installed, you will
> have the latest release of `kubectl` installed in your path. 
> You can skip the rest of this section. 

The release tarball contains various builds of `kubectl` as
`platforms/<os>/<arch>/kubectl`. There is also a helper script under
`cluster/kubectl.sh` which will identify and run the proper version. Do
one of the following:

* Add the proper `platforms/<os>/<arch>/kubectl` binary to your path.
* Use `cluster/kubectl.sh` in lieu of `kubectl`.

Run `kubectl` or `cluster/kubectl.sh` and verify communication with the 
cluster:

```
kubectl cluster-info
```

You should see a list of url endpoints for the cluster. 


## Cleanup

Don't do it yet, but shutting down the cluster is simply:

```
cluster/kube-down.sh
```
