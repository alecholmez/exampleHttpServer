FORMAT: 1A

FORMAT: 1A

# H2TP

An example boilerplate web-server written in golang

# Group Users

Resources related to the users in the API.

## Users Collection [/users]

### List All Users [GET]
Retrieves all crews from the database

+ Response 200 (application/json)

        [
          {
            "_id": "5825d0abd65d88006988d33c",
            "age": 19,
            "createdAt": "2016-11-11T14:07:39.522Z",
            "email": "alecholmez@me.com",
            "name": "Alec Holmes"
          },
          {
            "_id": "5825d0cdd65d88006988d33d",
            "age": 19,
            "createdAt": "2016-11-11T14:08:13.874Z",
            "email": "alecholmez@me.com",
            "name": "Alec Holmes"
          },
          {
            "_id": "5825d0ced65d88006988d33e",
            "age": 19,
            "createdAt": "2016-11-11T14:08:14.465Z",
            "email": "alecholmez@me.com",
            "name": "Alec Holmes"
          },
          {
            "_id": "5825d0cfd65d88006988d33f",
            "age": 19,
            "createdAt": "2016-11-11T14:08:15.072Z",
            "email": "alecholmez@me.com",
            "name": "Alec Holmes"
          },
          {
            "_id": "5825d0d0d65d88006988d340",
            "age": 19,
            "createdAt": "2016-11-11T14:08:16.205Z",
            "email": "alecholmez@me.com",
            "name": "Alec Holmes"
          }
        ]

+ Response 404 (application/json)

        Not found

### Create A User [POST]
Creates a user in the database

+ Request (application/json)

        {
            "name": "Alec Holmes",
            "email": "alecholmez@me.com",
            "age": 19
        }

+ Response 200 (application/json)

        {
          "_id": "5825d0d0d65d88006988d340",
          "age": 19,
          "createdAt": "2016-11-11T14:08:16.205Z",
          "email": "alecholmez@me.com",
          "name": "Alec Holmes"
        }

+ Response 400 (application/json)

        Request can not be blank

## User [/user/{id}]

+ Parameters
    + id (required, string) - user id

### Get A User [GET]
Retrieves a crew from the database

+ Response 200 (application/json)

        {
          "_id": "5825d0d0d65d88006988d340",
          "age": 19,
          "createdAt": "2016-11-11T14:08:16.205Z",
          "email": "alecholmez@me.com",
          "name": "Alec Holmes"
        }

+ Response 404 (application/json)

        user not found

### Delete A User [DELETE]
Retrieves a crew from the database

+ Response 200 (application/json)

+ Response 404 (application/json)

        user not found

### Update A User [PUT]
Retrieves a crew from the database

+ Request (application/json)

        {
            "name": "Alec Holmez"
        }

+ Response 200 (application/json)

        {
          "_id": "5825d0d0d65d88006988d340",
          "age": 19,
          "createdAt": "2016-11-11T14:08:16.205Z",
          "email": "alecholmez@me.com",
          "name": "Alec Holmez"
        }

+ Response 404 (application/json)

        user not found
