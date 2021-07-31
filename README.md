# demo-crud-api-rest-go
Api persona CRUD Go

Docker:
* docker build -t demo-crud-api-rest-go:v1 .
* docker network create gonetwork
* docker run -p 3306:3306 --name go-mysql8 --network gonetwork -e MYSQL_ROOT_PASSWORD=mysql -e MYSQL_DATABASE=goSchema -d mysql:8
* docker run -p 9000:9000 --network gonetwork --name demo-crud-apirest-go -d demo-crud-api-rest-go:v1