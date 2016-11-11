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

        {
            "users": [
                {
                    "name": "Alec Holmes",
                    "email": "alecholmez@me.com",
                    "age": "19"
                },
                {
                    "name": "Other Person",
                    "email": "testing@gmail.com",
                    "age": "24"
                },
                {
                    "name": "Test Person",
                    "email": "test.person@other.com",
                    "age": "57"
                }
            ]
        }

+ Response 404 (application/json)

        Not found

### Create A User [POST]
Creates a user in the database

+ Request (application/json)


+ Response 200 (application/json)


+ Response 400 (application/json)

        Request can not be blank

## User [/user/{id}]

+ Parameters
    + id (required, string) - user id

### Get A User [GET]
Retrieves a crew from the database

+ Response 200 (application/json)


+ Response 404 (application/json)

        user not found

### Delete A User [DELETE]
Retrieves a crew from the database

+ Response 200 (application/json)


+ Response 404 (application/json)

        user not found

### Update A User [PUT]
Retrieves a crew from the database

+ Response 200 (application/json)


+ Response 404 (application/json)

        user not found
