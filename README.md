# Kubernetes Playground

Playground to learn deployment a web service on a cluster.
For more details and motivation, check [my post](https://harunsasmaz.com/en/projects/kubernetes-playground/)

## Before you start

Make sure you have a Google Cloud Platform (GCP) account and then:
> GCP provides 300$ free budget for new users.

- Create a project
- Enable Google Kubernetes Engine (GKE).
- Enable Google Container Registry (GCR).

Then:

- Install [Go](https://go.dev/doc/install)
- Install [Terraform CLI](https://learn.hashicorp.com/tutorials/terraform/install-cli).
- Install [Kubernetes CLI](https://kubernetes.io/releases/download/)
- Install [Google Cloud CLI](https://formulae.brew.sh/cask/google-cloud-sdk)

Login to your GCP account through `gcloud`:

```
gcloud auth application-default login
```

Configure Docker to access private registry:

```
gcloud auth configure-docker
```

## Create a cluster

Before you start, make sure you set variables according to your project in `variables.tf`

```
cd <PROJECT_DIRECTORY>
terraform init
terraform apply
```

This may take a while. After your cluster is created,

```
gcloud container clusters get-credentials <CLUSTER_NAME>
```

This will set kubectl to access your cluster. To verify:

```
kubectl cluster-info
```

### Allow Kubernetes to Pull Images from GCR

```
kubectl create secret docker-registry gcr-access-token \
--docker-server=eu.gcr.io \
--docker-username=oauth3accesstoken \
--docker-password="$(gcloud auth print-access-token)" \
--docker-email=your@email.com
```

## Deploy Hello Service

```
make hello-push-image
kubectl apply -f definitions/hello
```

## Deploy TODO Service

We start with deploying databases that we use, and then deploy TODO service.

### Deploy PostgreSQL

1. Create a persistent volume
```
kubectl apply -f definitions/postgre/pg-pv.yaml
```
2. Create a persistenv volume claim
```
kubectl apply -f definitions/postgre/pg-pvc.yaml
```
3. Create PostreSQL Deployment
```
kubectl apply -f definitions/postgre/pg-deployment.yaml
```
4. Create PostgreSQL Service
```
kubectl apply -f definitions/postgre/pg-service.yaml
```

5. Verify if it works properly:

```
kubectl get pods
```

If the status of the pod with name postgres is running, then it works properly.

6. Use secrets for password protection

```
kubectl create secret generic pg-password \
--from-literal=password=<YOUR_PASSWORD>
```

### Deploy Redis

1. Create config map
```
kubectl apply -f definitions/redis/redis-configmap.yaml
```
2. Create Redis Deployment
```
kubectl apply -f definitions/redis/redis-deployment.yaml
```
3. Create Redis Service
```
kubectl apply -f definitions/redis/redis-service.yaml
```

### Deploy TODO Service

1. Build and push Docker image 
```
make todo-push-image
kubectl apply -f definitions/todo
```
