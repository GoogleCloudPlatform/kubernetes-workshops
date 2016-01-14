# Workshop Tests


## To run

Have a working kubectl connected to a cluster, a working go
environment, and run:

```
go test
```

## Tagging Workshops

In each workshop, surround the lines you wish to test with `START
bash` and `END bash`. `START`, `END`, and `bash` must have spaces on
either side and are case sensitive. All lines between will be run,
except those that start with ```.

Example:

    <!-- START bash -->
    ```
    kubectl get pods
    ```
    <!-- END bash -->

You could also put commands inside comments, that you don't want to
show up in the workshop:

```
<!-- START bash
echo "do setup"
END bash -->
```
