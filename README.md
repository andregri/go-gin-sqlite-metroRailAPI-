# go-gin-sqlite-metroRailAPI-
Restful API for a metro rail usin Go Gin and SQLite3


## Tests

**Create** a new station:
```
curl -X POST http://localhost:8000/v1/stations \
    -H 'cache-control: no-cache' \
    -H 'content-type: application/json' \
    -d '{"name":"Brooklyn", "opening_time":"8:12:00", "closing_time":"18:23:00"}'

{"result":{"id":1,"name":"Brooklyn","opening_time":"8:12:00","closing_time":"18:23:00"}}
```

**Get** the newly created station:
```
curl -X GET http://localhost:8000/v1/stations/1

{"result":{"id":1,"name":"Brooklyn","opening_time":"8:12:00","closing_time":"18:23:00"}}
```

**Delete** the station
```
curl -X DELETE http://localhost:8000/v1/stations/1

{"result":"1"}
```