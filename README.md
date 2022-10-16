# go-microservices

## 注册用户
curl --request POST \
--url http://localhost:5001/auth/register \
--header 'Content-Type: application/json' \
--data '{
"email":"malen.ma@gmail.com",
"password":"123456"
}'

## 用户登录
curl --request POST \
--url http://localhost:5001/auth/login \
--header 'Content-Type: application/json' \
--data '{
"email":"malen.ma@gmail.com",
"password":"123456"
}'


## 创建商品
curl --request POST \
--url http://localhost:5001/product/ \
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoLXN2YyIsImV4cCI6MTY5NzQ5NzY3OSwiSWQiOjEsIkVtYWlsIjoibWFsZW4ubWFAZ21haWwuY29tIn0.CNlUDKQFY8NCHrbuaj5ZbtVu9KNI9CpepzR58TrAsHg' \
--header 'Content-Type: application/json' \
--data '{
 "name": "Product A",
 "stock": 5,
 "price": 15
}'

## 查找商品
curl --request GET \
--url http://localhost:5001/product/1 \
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoLXN2YyIsImV4cCI6MTY5NzQ5NzY3OSwiSWQiOjEsIkVtYWlsIjoibWFsZW4ubWFAZ21haWwuY29tIn0.CNlUDKQFY8NCHrbuaj5ZbtVu9KNI9CpepzR58TrAsHg'


## 创建订单
curl --request POST \
  --url http://localhost:5001/order/ \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoLXN2YyIsImV4cCI6MTY5NzQ5NzY3OSwiSWQiOjEsIkVtYWlsIjoibWFsZW4ubWFAZ21haWwuY29tIn0.CNlUDKQFY8NCHrbuaj5ZbtVu9KNI9CpepzR58TrAsHg' \
  --header 'Content-Type: application/json' \
  --data '{
 "productId": 1,
 "quantity": 1
}'