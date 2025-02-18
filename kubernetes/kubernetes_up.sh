# Construir a imagem localmente
docker build -t fastfood-payment-app:latest .

# Carregar a imagem no Minikube
## minikube image load fastfood-payment-app:latest

# Aplicar os arquivos de configuração do Kubernetes
kubectl apply -f kubernetes/postgres-dbinit-configmap.yaml
kubectl apply -f kubernetes/postgres-pv.yaml
kubectl apply -f kubernetes/postgres-pvc.yaml
kubectl apply -f kubernetes/postgres-deploy.yaml
kubectl apply -f kubernetes/postgres-service.yaml
kubectl apply -f kubernetes/fastfood-deployment.yaml
kubectl apply -f kubernetes/fastfood-service.yaml
kubectl apply -f kubernetes/fastfood-hpa.yaml