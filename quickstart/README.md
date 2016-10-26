# Kubernetes Quickstart

## Prerequisites

* Have a cluster running and a `kubectl` binary configured to talk to
  that cluster

## Lab

The container config of [Lobsters](https://github.com/jcs/lobsters) is
[here](lobsters/). It is a test configuration with a local sqlite db
inside the single container. The built container image has been pushed
to `gcr.io/google-samples/lobsters:1.0`.

Run the Lobsters app container

<!-- START bash -->
```
kubectl run lobsters --image=gcr.io/google-samples/lobsters:1.0
```
<!-- END bash -->

```
deployment "lobsters" created
```

Get deployments

```
kubectl get deployments
```

```
NAME       DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
lobsters   1         1         1            1           1m
```

Look at the pods
<!-- START bash -->
```
kubectl get pods
```
<!-- END bash -->

```
NAME                        READY     STATUS    RESTARTS   AGE
lobsters-3295177660-xg5t2   1/1       Running   0          55s
```

Check the log, use your specific pod name from above.

```
kubectl logs lobsters-3295177660-xg5t2
```

```
[2016-04-07 17:36:37] INFO  WEBrick 1.3.1
[2016-04-07 17:36:37] INFO  ruby 2.3.0 (2015-12-25) [x86_64-linux]
[2016-04-07 17:36:37] INFO  WEBrick::HTTPServer#start: pid=1 port=3000
```

Let's see if lobsters is up and running, first we will do a private
port forward from our machine to the container.

> Note: If you're running kubectl on a remote machine (or VM), skip
> this step. ("localhost" is remote so it won't work)

> Note: For Mac OS / Windows, use a port other than 8080.

```
kubectl port-forward lobsters-3295177660-xg5t2 8080:3000
```

```
I0407 10:41:41.872146   15115 portforward.go:213] Forwarding from 127.0.0.1:8080 -> 3000
I0407 10:41:41.872238   15115 portforward.go:213] Forwarding from [::1]:8080 -> 3000
I0407 10:41:50.548148   15115 portforward.go:247] Handling connection for 8080
I0407 10:41:50.553497   15115 portforward.go:247] Handling connection for 8080
I0407 10:41:51.068845   15115 portforward.go:247] Handling connection for 8080
...
```

Visit `http://localhost:8080/` in your browser. Ctrl-C to cancel.

Good, now we will open up our Lobsters to the internet.

<!-- START bash -->
```
kubectl expose deployment lobsters --port=80 --target-port=3000 --type=LoadBalancer
```
<!-- END bash -->

```
service "lobsters" exposed
```

Get the external IP, this will take a minute. EXTERNAL-IP will be blank until the load balancer is ready.

```
kubectl get service lobsters
```

```
NAME       CLUSTER-IP    EXTERNAL-IP      PORT(S)   AGE
lobsters   10.3.241.32   1.2.3.4          80/TCP    1m
```

Visit the live site in your browser: `http://<external-ip>`

Congrats! Your Lobsters is live

> Non-LoadBalancer
>
> ```
> kubectl expose deployment lobsters --port=3000 --type=NodePort
> kubectl get svc lobsters -o yaml | grep nodePort
> ```
>
> Access by visiting the IP of any one of your cluster nodes on the
> port shown by the second command.

## Cleanup

Delete your resources

<!-- START bash -->
```
kubectl delete deployment,service lobsters
```
<!-- END bash -->
<!-- START bash
while kubectl get deployment lobsters; do echo running; sleep 1; done
while kubectl get service lobsters; do echo running; sleep 1; done
while [ -n "$(kubectl get pod -l run=lobsters)" ]; do echo running; done
END bash -->
