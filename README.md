This is a task provided by ADmyBRAND India.

# REST-API-go
A REST api based on Go lang.
This is a example of user data store REST api where one can get user details.
For **v1** API fake db is made with slice(like array). There are 5 basic operation(mainly CURD) in this example. One can get all user list, Create a new user, Update a existing user by id, Read a user by id, and Delete a user by id.

## Get user list
---
Call a Get request to **/v1/users**. 

Client will get a list/array  of json object like below.

_Server response_

```
[
    {
        "id": "23",
        "name": "test",
        "dob": "020689",
        "address": "India",
        "description": "test descrption",
        "createdAt": ""
    },
    {
        "id": "57",
        "name": "test",
        "dob": "020689",
        "address": "India",
        "description": "test descrption",
        "createdAt": ""
    }
]
```

## Create a new user 
---
Call a POST request to **/v1/user** with below json object in body.

```
{
        
        "name": "testNew",
        "dob": "020689",
        "address": "India",
        "description": "test descrption",
        "createdAt": ""
}
```

__Note:__ id is not necessary to create a new user as it will genereate a randoom user id and you can check it via server response. 

_Server response_

```
{
    "id": "847",
    "name": "testNew",
    "dob": "020689",
    "address": "India",
    "description": "test descrption",
    "createdAt": ""
}
```

## Update a existing user
---
Call a PUT request to **/v1/user{id}** with below json object in body.

```
  {
        
        "name": "testupdate",
        "dob": "020689",
        "address": "India",
        "description": "test descrption",
        "createdAt": ""
    }
```


**Note:**  Partial update is not valid. Id will not be updated.




_Server response_
```
{
    "id": "23",
    "name": "testupdate",
    "dob": "020689",
    "address": "India",
    "description": "test descrption",
    "createdAt": ""
}
```
## Read a existing user details
---
Call a Get request to **/v1/user/{id}** (here id is user id). 

Client will get a json object like below.
```
{
    "id": "23",
    "name": "testupdate",
    "dob": "020689",
    "address": "India",
    "description": "test descrption",
    "createdAt": ""
}
```

## Delete a existing user
---
To delete a user details , a Delete request should be called with id to  **/v1/user/{id}** (here id is user id).

Server will response back with remaining users list.

_Server response_
```
[
    {
        "id": "57",
        "name": "test",
        "dob": "020689",
        "address": "India",
        "description": "test descrption",
        "createdAt": ""
    },
    {
        "id": "81",
        "name": "test2",
        "dob": "020689",
        "address": "India",
        "description": "test descrption",
        "createdAt": ""
    }
]
```