## gin-glm-api

Basic go http service

## Available endpoints
 `/game-records` - list all records
 Sample
 ```
    curl localhost:8080/game-records
 ```

 `/create-game-record` - create a record
 Sample
 ```
    curl -X POST localhost:8080/create-game-record \
        -H "Content-Type: application/json" \
        -d '{ "name": "Disco Elysium", "platform": "PC", "medal": "Gold", "complete_time_hours": 50, "genre": "CRPG" }'
 ```

 `/game-record/:id` - get a single record
 Sample
 ```
    curl localhost:8080/game-record/5
 ```

 `/update-game-record/:id"` - update a record
  Sample
  ```
    curl -X PATCH localhost:8080/update-game-record/5 \
        -d '{ "medal": "Silver" }'
  ```

 `/delete-game-record/:id` - delete a record
 ```
    curl -X DELETE localhost:8080/delete-game-record/:id
 ```

 ## TODO
    - Auth
    - CSV loader
    - JSON loader
