cd /Users/sorin.buliarca/projects/kubernetes-workshops/advanced/
for i in {1..12}; do
  kubectl create ns ns${i}
  kubectl create secret generic db-pass --from-literal=password=supersecret -n ns${i}
  kubectl apply -f logstash.yaml,service-local.yaml,database-pvc.yaml,rake-db.yaml,frontend-sidecar-with-resources.yaml -n ns${i}
  kubectl autoscale deployment lobsters --cpu-percent=10 --min=1 --max=5 -n ns${i}
done

sleep 120
for i in {1..12}; do
  port=$(kc get svc lobsters --no-headers -n ns${i} | awk '{print $5}' | sed "s/.*:\(.*\)\/.*/\1/g");
  nohup siege -t180s -c150 "http://cjdv-k8-master.ep.esp.local:${port}/" &
done