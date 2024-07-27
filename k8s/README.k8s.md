# Kubernetes

Yes, you can certainly test the Kubernetes Deployment manually before integrating it with Argo CD. This will help ensure that your application is configured correctly and works as expected. Here are the steps to do this:

### Step 1: Create Kubernetes Manifests

First, create the necessary Kubernetes manifests (e.g., Deployment and Service) locally.

#### Example `deployment.yaml`:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-app
  template:
    metadata:
      labels:
        app: go-app
    spec:
      containers:
        - name: go-app
          image: robertocamp/argocd-proj2:latest
          ports:
            - containerPort: 8080
```

#### Example `service.yaml`:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: go-app-service
spec:
  selector:
    app: go-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

### Step 2: Deploy the Application Manually

1. **Apply the Deployment:**
   ```sh
   kubectl apply -f deployment.yaml
   ```

2. **Apply the Service:**
   ```sh
   kubectl apply -f service.yaml
   ```

### Step 3: Verify the Deployment

1. **Check the Pods:**
   ```sh
   kubectl get pods
   ```
   Ensure that the pod is running without issues.

2. **Check the Service:**
   ```sh
   kubectl get svc
   ```
   Verify that the service is created and has an external IP (if applicable).

3. **Access the Application:**
   - If you are using a LoadBalancer service type, you can access the application using the external IP.
   - If you are using a NodePort service type, you can access the application using the node IP and the node port.

### Step 4: Revoke the Deployment

Once you have verified that the application is running correctly, you can delete the deployment and service to clean up the environment before integrating with Argo CD.

1. **Delete the Deployment:**
   ```sh
   kubectl delete -f deployment.yaml
   ```

2. **Delete the Service:**
   ```sh
   kubectl delete -f service.yaml
   ```

### Step 5: Proceed with Argo CD Integration

#### port forwarding

To deploy your application into a specific namespace, you need to create the namespace if it doesn't exist and then specify the namespace in your deployment and service YAML file. Here's how you can do it:

### Step 1: Create the Namespace

First, create the namespace (e.g., `argocd-proj2`):

```sh
kubectl create namespace argocd-proj2
```

### Step 2: Update the YAML File to Include the Namespace

Update your consolidated YAML file to include the namespace:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
  namespace: argocd-proj2
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-app
  template:
    metadata:
      labels:
        app: go-app
    spec:
      containers:
        - name: go-app
          image: robertocamp/argocd-proj2:latest
          ports:
            - containerPort: 8000
          env:
            - name: PORT
              value: "8000"
---
apiVersion: v1
kind: Service
metadata:
  name: go-app-service
  namespace: argocd-proj2
spec:
  selector:
    app: go-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8000
```

### Step 3: Apply the YAML File to the Specified Namespace

Apply the YAML file to the `argocd-proj2` namespace:

```sh
kubectl apply -f go-app-deployment.yaml
```

### Step 4: Port-Forward to the Service in the Specified Namespace

Set up port-forwarding to the service in the specified namespace:

```sh
kubectl port-forward svc/go-app-service -n argocd-proj2 8000:80
```

### Verification

1. **Check the Pods in the Namespace**:
   ```sh
   kubectl get pods -n argocd-proj2
   ```
   Ensure that the pod is running in the `argocd-proj2` namespace.

2. **Check the Service in the Namespace**:
   ```sh
   kubectl get svc -n argocd-proj2
   ```
   Ensure that the service is created and is exposing port 80 in the `argocd-proj2` namespace.

3. **Port-Forward**:
   ```sh
   kubectl port-forward svc/go-app-service -n argocd-proj2 8000:80
   ```
   Ensure that the port-forwarding is working.

4. **Access the Application**:
   Open your browser and go to `http://localhost:8000`. You should see the message "Hello, Docker! <3".

By specifying the namespace in your YAML file and using the `-n` flag with `kubectl` commands, you can deploy and manage your application in the desired namespace.