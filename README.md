# cqrs-example
CQRS example with Go, MySQL, NATS, ElasticSearch, Kubernetes

## Deployment

Please make sure you installed Docker, Kubernetes before running below commands:

```sh
# Deploy infra services
cd k8s/
kubectl create -f mysql.yaml
kubectl create -f elasticsearch.yaml
kubectl create -f nats.yaml

# Deploy our services
cd ../books
make first_deploy
cd ../book-query
make first_deploy

# Post a book, this should result in data to be stored in MySQL and trigger book-query service to store data into ElasticSearch as well
curl --header "Content-Type: application/json" --request POST --data '{"name":"The Outliers"}' http://192.168.99.100:30305/api/v1/books

# Query the data, this should go into MySQL
curl http://192.168.99.100:30306/api/v1/books
# Search for the data, this should go into ElasticSearch
curl http://192.168.99.100:30306/api/v1/books/search?query="outliers"

```

