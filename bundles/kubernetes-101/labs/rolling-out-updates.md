# Rolling out Updates

Kubernetes makes it easy to rollout updates to your applications using the builtin rolling update mechanism. In this lab you will learn how to:

* Modify deployments to trigger rolling updates
* Pause and resume an active rolling update
* Rollback a deployment to a previous revision

Update to kubectl 1.3.0

```
wget https://storage.googleapis.com/kubernetes-release/release/v1.3.0/bin/linux/amd64/kubectl
chmod +x kubectl
sudo mv kubectl /usr/local/bin
```

## Tutorial: Rollout a new version of the Auth service

```
/usr/local/bin/kubectl rollout history deployment auth
```

Modify the auth deployment image:

```
vim deployments/auth.yaml
```

```
image: "kelseyhightower/auth:2.0.0"
```

```
/usr/local/bin/kubectl apply -f deployments/auth.yaml --record
```

```
/usr/local/bin/kubectl describe deployments auth
```

```
/usr/local/bin/kubectl get replicasets
```

```
/usr/local/bin/kubectl rollout history deployment auth
```

## Tutorial: Pause and Resume an Active Rollout

```
/usr/local/bin/kubectl rollout history deployment hello
```

Modify the hello deployment image:

```
vim deployments/hello.yaml
```

```
image: "kelseyhightower/hello:2.0.0"
```

```
/usr/local/bin/kubectl apply -f deployments/hello.yaml --record
```

```
/usr/local/bin/kubectl describe deployments hello
```

```
/usr/local/bin/kubectl rollout pause deployment hello
```

```
/usr/local/bin/kubectl rollout resume deployment hello
```

## Exercise: Rollback the Hello service

Use the `kubectl rollout undo` command to rollback to a previous deployment of the Hello service.

## Summary

In this lab you learned how to rollout updates to your applications by modifying deployment objects to trigger rolling updates. You also learned how to pause and resume an active rolling update and rollback it back using the `kubectl rollout` command.
