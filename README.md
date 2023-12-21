API : 
GET    http://localhost:8080/:path

Request URL
```
http://localhost:8080/XTWZQCPUMEZ5
```
Redirect to http://www.spacex.com

------------------------------------------------------------------------
GET    http://localhost:8080/v1/health -> status check
------------------------------------------------------------------------
POST   http://localhost:8080/v1/addurl

Request Body:
```
{
    "url":"www.spacex.com"
}
```
Response Body :
```
http://localhost:8080/XTWZQCPUMEZ5
```


------------------------------------------------------------------------
GET    http://localhost:8080/v1/mostvisit/:count

Request URL
```
curl --location 'http://localhost:8080/v1/mostvisit/4'
```

Response Body: 
```
{
    "http://https://github.com": 1,
    "http://www.facebook.com": 2,
    "http://www.google.com": 3,
    "http://www.yahoo.com": 2
}
```

------------------------------------------------------------------------
Docker image :- https://hub.docker.com/r/sakibcoolz/zlink_service
or run this dierctly
```
docker run -e SERVICEHOST=localhost -e SERVICEPORT=10000 -p 10000:10000 sakibcoolz/zlink_service
```

