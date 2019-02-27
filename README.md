# user-service
[![Build Status](https://travis-ci.org/davidseyboldblog/user-service.svg?branch=master)](https://travis-ci.org/davidseyboldblog/user-service)

```
Request:
GET /user-service/user/:id

Response:
200 OK
{
  "id": 1,
  "name": "First Last",
  "email": "first.last@example.com",
  "phone": "1231231234"
}
```

``` 
Request:
POST /user-service/user
{
  "name": "First Last",
  "email": "first.last@example.com",
  "phone": "1231231234"
}

Response:
201 CREATED
{
  "id": 1
}
```

```
Request:
PUT /user-service/:id
{
  "name": "First Last",
  "email": "first.last@example.com",
  "phone": "1231231234"
}

Response:
204 NO CONTENT
```