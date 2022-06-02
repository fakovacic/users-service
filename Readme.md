# Users service

## Run

```
make compile
make docker-build
docker-compose up
```

## Routes

### GET - /users

- return list of users

#### query params

- `limit` - define how much rows will be returned, default is 10
- `page` - define offset for paging
- `filters.{field}` - optional, use for filter results `filters.country=XX`


#### Example requests

```
curl -X GET http://localhost:8080/users

curl -X GET http://localhost:8080/users?limit=100&page=2

curl -X GET http://localhost:8080/users?filters.country=XX
```



### POST - /users

- create user

#### Example requests

```
curl -X POST http://localhost:8080/users -d '{"user":{"email":"my@email.test"}}'
```


### PATCH - /users/:id

- update user

#### Example requests

```
curl -X PATCH http://localhost:8080/users/ecf7e1a1-e2e5-4fb6-9c1b-abd1a8084d2b -d '{"fields":["email"],"user":{"email":"my@email.tester"}}'
```


### DELETE - /users/:id

- delete user

#### Example requests

```
curl -X DELETE http://localhost:8080/users/ecf7e1a1-e2e5-4fb6-9c1b-abd1a8084d2b'
```




