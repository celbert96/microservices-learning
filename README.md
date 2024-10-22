# Microservices Learning


## How to Run Locally in a K8s Cluster with Minikube

### Start minikube
```minikube start```

### Use minikube docker registry
```eval $(minikube docker-env)```

### Build Docker image
```docker build --tag go-application ./go-application```

### Create the namespaces
```kubectl create -f ./kubernetes/namespace.yaml```

### Create the deployment
```kubectl create -f ./kubernetes/deployment.yaml -n {namespace}```

### Create the LoadBalancer service
```kubectl create -f ./kubernetes/service.yaml -n {namespace}```

### List pods
```kubectl get pods -n {namespace}```

### Port forward 
```kubectl port-forward {podname} 8080:8080 -n {namespace}```