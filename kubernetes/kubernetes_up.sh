# Construir a imagem localmente
docker build -t fastfood-payment-app:latest .
kubectl apply -f kubernetes/fastfood-payment-secrets.yaml

# Carregar a imagem no Minikube
## minikube image load fastfood-payment-app:latest

# Aplicar os arquivos de configuração do Kubernetes
# kubectl apply -f kubernetes/postgres-dbinit-configmap.yaml
# kubectl apply -f kubernetes/postgres-pv.yaml
# kubectl apply -f kubernetes/postgres-pvc.yaml
# kubectl apply -f kubernetes/postgres-deploy.yaml
# kubectl apply -f kubernetes/postgres-service.yaml

kubectl create secret generic aws-secrets \
  --from-literal=access-key-id=$(aws configure get aws_access_key_id) \
  --from-literal=secret-access-key=$(aws configure get aws_secret_access_key) \
  --from-literal=access-session-token=$(aws configure get aws_session_token)

kubectl apply -f kubernetes/fastfood-payment-deployment.yaml
kubectl apply -f kubernetes/fastfood-payment-service.yaml
kubectl apply -f kubernetes/fastfood-payment-hpa.yaml