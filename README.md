# Simple REST API using Go with MongoDB

## How to 

* Create database and collection 
* Setup config param in config.yaml
* Go get dependency 
* Run app (currently in port 3000)

```
go run app.go
```

Sample POST request 

```
curl --location --request POST 'http://localhost:3000/movies' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "Toy Story 2",
    "cover_image": "http://test_cover_movie2",
    "description": "Kids Movie Too"    
}'

```



credit https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94 by Mohamed Labouardy