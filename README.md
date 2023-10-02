# x-tention-crew-mico-services

# To Run Project

#### Clone The Repository
```
git clone https://github.com/nikhilnarayanan623/x-tention-crew-mico-services
```
#### Checkout To Project Directory
```
cd ./x-tention-crew-mico-services
```
## Run using docker-compose

##### Setup env file for docker-compose (look up the .env.example for your reference)
create .env file on the project root dir and add the below envs on it
```.env
#example
SERVICE2_REST_PORT="8001"

USER_DB_HOST="postgresdb"
USER_DB_PORT="5432"
USER_DB_USER="myuser"
USER_DB_PASSWORD="password"
USER_DB_NAME="mydb"

USER_SERVICE_REST_PORT="8000"
USER_SERVICE_GRPC_HOST="app1"
USER_SERVICE_GRPC_PORT="50051"

REDIS_HOST="redisdb"
REDIS_PORT="6379"

POSTGRES_USER="myuser"
POSTGRES_PASSWORD="password"
POSTGRES_DB="mydb"
```
#### Run Docker Containers
```
docker compose up
```
#### To See API Documetnation For User-Service
http://localhost:{USER-SERVICE-REST-PORT}/swagger/index.html#/

#### To See API Documetnation For Service-2
http://localhost:{SERVICE-2-REST-PORT}/swagger/index.html#/
