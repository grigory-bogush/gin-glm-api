## gin-glm-api

Basic go http service

## Auth is read from env
    `ADMIN_USERNAME` & `ADMIN_PASSWORD`, add those values to `.env` for running locally

## Available endpoints
 `/game-records` - list all records
 Sample
 ```
    curl localhost:8080/game-records -u admin:admin
 ```

 `/create-game-record` - create a record
 Sample
 ```
    curl -X POST localhost:8080/create-game-record \
        -H "Content-Type: application/json" \
        -u admin:admin \
        -d '{ "name": "Disco Elysium", "platform": "PC", "medal": "Gold", "complete_time_hours": 50, "genre": "CRPG" }'
 ```

 `/game-record/:id` - get a single record
 Sample
 ```
    curl localhost:8080/game-record/5 -u admin:admin
 ```

 `/update-game-record/:id"` - update a record
  Sample
  ```
    curl -X PATCH localhost:8080/update-game-record/5 \
        -u admin:admin \
        -d '{ "medal": "Silver" }'
  ```

 `/delete-game-record/:id` - delete a record
 ```
    curl -X DELETE localhost:8080/delete-game-record/:id  -u admin:admin
 ```

 ## TODO
    - CSV loader
    - JSON loader
