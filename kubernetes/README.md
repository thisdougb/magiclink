# magiclink

### Deploy To Kubernetes
The magiclink service depends on a running redis service.
I have added a k8s deployment yaml file for redis, for testing.

Get the magiclink service and deployment spec:
```
$ curl -sSLO https://raw.githubusercontent.com/thisdougb/kubernetes/magiclink.yaml
$ kubectl apply -f magiclink.yaml    
service/magiclink created
deployment.apps/magiclink created
$
```
Check the pod is running OK:
```
$ kubectl get po
NAME                         READY   STATUS    RESTARTS   AGE
magiclink-7c76857f8c-jghcr   1/1     Running   0          26s
redis-df87ffcd6-vqh2b        1/1     Running   0          14m
$
$ kubectl logs magiclink-7c76857f8c-jghcr
2021/11/07 12:45:55 webserver.Start(): listening on port 8080
$
```
And the service:
```
$ kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
kubernetes   ClusterIP   10.245.0.1      <none>        443/TCP    127m
magiclink    ClusterIP   10.245.87.158   <none>        80/TCP     23s
redis        ClusterIP   10.245.59.174   <none>        6379/TCP   13m
```
You can then check all is OK by port-forwarding, and using curl.
Here we forward from localhost:8080 to magiclink:80 (the service is listening on port 80):
```
$ kubectl port-forward service/magiclink 8080:80
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
```
In another terminal we can curl the magiclink service:
```
$ curl --data '{"email":"me@mydomain.com"}' -X POST http://localhost:8080/send/
OK
```
