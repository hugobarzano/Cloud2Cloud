---
apiVersion: "extensions/v1beta1"
kind: "Deployment"
metadata:
  name: "nginx-1"
  namespace: "default"
  labels:
    app: "nginx-1"
spec:
  replicas: 3
  selector:
    matchLabels:
      app: "nginx-1"
  template:
    metadata:
      labels:
        app: "nginx-1"
    spec:
      containers:
      - name: "nginx"
        image: "nginx:latest"
---
apiVersion: "autoscaling/v1"
kind: "HorizontalPodAutoscaler"
metadata:
  name: "nginx-1-hpa"
  namespace: "default"
  labels:
    app: "nginx-1"
spec:
  scaleTargetRef:
    kind: "Deployment"
    name: "nginx-1"
    apiVersion: "apps/v1beta1"
  minReplicas: 1
  maxReplicas: 5
  targetCPUUtilizationPercentage: 80






  ---
  apiVersion: "v1"
  kind: "Service"
  metadata:
    name: "nginx-1-service"
    namespace: "default"
    labels:
      app: "nginx-1"
  spec:
    ports:
    - protocol: "TCP"
      port: 80
    selector:
      app: "nginx-1"
    type: "LoadBalancer"
    loadBalancerIP: ""
