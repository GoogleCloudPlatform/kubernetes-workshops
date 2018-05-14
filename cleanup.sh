for i in {1..15}; do
  kubectl delete ns ns${i}
done
sleep 10
kubectl delete -f advanced/local-pvs.yaml