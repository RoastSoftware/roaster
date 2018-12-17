FORMAT: 1A

HOST: https://roast.software/

# Roaster

The Roaster API let's people validate their code using statical code analysis.
Users can register and track their progress, also allowing friends to follow
your progress.

# Roaster API Root [/]
This resource returns the base `index.html`. The index contains JavaScript that
implements a SPA (Single Page Application) that is able to communicate with the
API.

## Retrieve the roast.software SPA [GET]
+ Response 200 (text/html)
```html
<!DOCTYPE html>
<html [...]>
[... The roast.software SPA ...]
</html>
```

## User [/user]
### Create User [POST]
+ Request a user registration (application/json)
    + Body
    ```json
    {
        "email": "email@example.com",
        "username": "roastin_roger",
        "password": "MyTr3me3d0usPassw0rd!",
        "fullname": "Roger Roger"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "email": {
                "type": "string"
            },
            "username": {
                "type": "string"
            },
            "password": {
                "type": "string"
            },
            "fullname": {
                "type": "string",
                "required": false
            }
        }
    }
    ```
+ Response 200 (application/json)
    + Headers
    ```
    Set-Cookie: roaster_auth=AB32DEAC21A91DE123[...]; Expires=Wed, 21 Oct 2015 07:28:00 GMT;
    ```
    + Body
    ```json
    {
        "email": "email@example.com",
        "username": "roastin_roger",
        "fullname": "Roger Roger"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "email": {
                "type": "string"
            },
            "username": {
                "type": "string"
            },
            "fullname": {
                "type": "string"
            }
        }
    }
    ```

### View/Handle Specific User [/user/{username}]
#### Change User [PATCH]
+ Request change user (application/json)
    + Body
    ```json
    {
        "email": "email@example.com",
        "password": "My3venM0reTr3mend0usP@ssw0rd!",
        "fullname": "Sir. Roger Roger"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "email": {
                "type": "string",
                "required": false
            },
            "password": {
                "type": "string",
                "required": false
            },
            "fullname": {
                "type": "string",
                "required": false
            }
        }
    }
    ```
+ Response 200 (application/json)
    + Body
    ```json
    {
        "email": "email@example.com",
        "username": "roastin_roger",
        "fullname": "Sir. Roger Roger"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "email": {
                "type": "string",
                "required": false
            },
            "username": {
                "type": "string"
            },
            "fullname": {
                "type": "string",
                "required": false
            }
        }
    }
    ```

#### Retrieve User Information [GET]
+ Request user information
+ Response 200 (application/json)
    + Body
    ```json
    {
        "email": "email@example.com",
        "username": "roastin_roger",
        "fullname": "Roger Roger"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "email": {
                "type": "string"
            },
            "username": {
                "type": "string"
            },
            "fullname": {
                "type": "string"
            }
        }
    }
    ```

#### Remove User [DELETE]
+ Request remove user
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
+ Response 200
    + Headers
    ```
    Set-Cookie: roaster_auth=deleted; Expires=Thu, 01 Jan 1970 00:00:00 GMT;
    ```

### Avatar For User [/user/{username}/avatar]
#### Upload User Avatar [PUT]
+ Request image to upload (multipart/form-data; boundary=---BOUNDARY)
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
    + Body
    ```
    --{boundary value}
    Content-Disposition: form-data; name="file"; filename="image.jpg"
    Content-Type: image/jpeg (according to the type of the uploaded file)
    {file content}
    --{boundary value}
    ```
+ Response 204

#### Retrieve User Avatar [GET]
+ Request user avatar
+ Response 200 (image/png)

### Friends for User [/user/{username}/friend]
#### Add New Friend [POST]
+ Request add new friend (application/json)
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
    + Body
    ```json
    {
        "friend": "MyAwesomeFriend"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "friend": {
                "type": "string",
                "required": true
            }
        }
    }
    ```
+ Response 200

#### Retrieve User Friends [GET]
+ Request user friends
+ Response 200 (application/json)
    + Body
    ```json
    {
        "friends": ["MyAwesomeFriend", "MyNotSoAwesomeFriend"]
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "friends": {
                "type": "array",
                "items": {
                    "type": "string"
                }
            }
        }
    }
    ```

#### Remove User Friend [DELETE]
+ Request remove user friend
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
+ Response 200

## Session [/session]
### Authenticate for New Session (sign in) [POST]
+ Request a session authentication (application/json)
    + Body
    ```json
    {
        "username": "roastin_roger",
        "password": "MyTr3me3d0usPassw0rd!",
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "username": {
                "type": "string"
            },
            "password": {
                "type": "string"
            }

        }
    }
    ```
+ Response 200 (application/json)
    + Headers
    ```
    Set-Cookie: roaster_auth=AB32DEAC21A91DE123[...]; Expires=Wed, 21 Oct 2015 07:28:00 GMT;
    ```
    + Body
    ```json
    {
        "email": "email@example.com",
        "username": "roastin_roger",
        "fullname": "Roger Roger"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "email": {
                "type": "string"
            },
            "username": {
                "type": "string"
            },
            "fullname": {
                "type": "string"
            }
        }
    }
    ```

### Get Current Session (resume login session) [GET]
+ Request current session
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
+ Response 200 (application/json)
    + Body
    ```json
    {
        "email": "email@example.com",
        "username": "roastin_roger",
        "fullname": "Roger Roger"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "email": {
                "type": "string"
            },
            "username": {
                "type": "string"
            },
            "fullname": {
                "type": "string"
            }
        }
    }
    ```

### Remove Current Session (sign out) [DELETE]
+ Request remove current session
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
+ Response 200
    + Headers
    ```
    Set-Cookie: roaster_auth=deleted; Expires=Thu, 01 Jan 1970 00:00:00 GMT;
    ```
## Roast [/roast]
### Analyze code snippet [POST]
+ Request a static analysation of provided code snippet (application/json)
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...] (string, optional) - Should be sent if user has an active session.
    ```
    + Body
    ```json
    {
        "language": "python3",
        "code": "print('I´m Roastin´ Roger... Roger.', invalid = 'Bönor är ändå rätt gött')",
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "language": {
                "type": "string"
            },
            "code": {
                "type": "string"
            }

        }
    }
    ```
+ Response 200 (application/json)
    + Body
    ```json
    {
        "grade": 12,
        "warnings": [
            {
                "row": 1,
                "column": 48,
                "engine": "pycodestyle",
                "name": "E251",
                "description": "no spaces around keyword / parameter equals"
            }
        ],
        "errors": [
            {
                "row": 1,
                "column": 41,
                "engine": "pyflakes",
                "name": "TypeError",
                "description": "print() got an unexpected keyword argument 'invalid'"
            }
        ]
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "grade": {
                "type": "number"
            },
            "warning": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "row": {
                            "type": "number"
                        },
                        "column": {
                            "type": "number"
                        },
                        "engine": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        },
                        "description": {
                            "type": "string"
                        }
                    }
                }
            },
            "errors": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "row": {
                            "type": "number"
                        },
                        "column": {
                            "type": "number"
                        },
                        "engine": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        },
                        "description": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
    ```

## Global Feed [/feed{?page}]
### Get Global Feed [GET]
+ Request a global feed (application/json)
    + Parameters
        + page: 0 (number) - Page of global feed to return.
+ Response 200 (application/json)
    + Body
    ```json
    {
        "items": [
            {
                "category": 0,
                "username": "JustARegularUsername",
                "title": "Woah, this user failed!",
                "description": "Something happened [...]",
                "time": "Mon Dec  3 10:51:52 CET 2018"
            }, {
                "category": 3,
                "username": "JustAnotherRegularUsername",
                "title": "Woah, this user also failed!",
                "description": "Something happened [...]",
                "time": "Mon Dec  3 09:48:12 CET 2018"
            }
        ]
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "items": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "category": {
                            "type": "number"
                        },
                        "username": {
                            "type": "string"
                        },
                        "title": {
                            "type": "string"
                        },
                        "description": {
                            "type": "string"
                        },
                        "time": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
    ```

### User Feed [/feed/{username}{?page}]
#### Get User Feed [GET]
+ Request a user feed (application/json)
    + Parameters
        + page: 0 (number) - Page of user feed to return.
+ Response 200 (application/json)
    + Body
    ```json
    {
        "items": [
            {
                "category": 0,
                "username": "JustARegularUsername",
                "title": "Woah, this user failed!",
                "description": "Something happened [...]",
                "time": "Mon Dec  3 10:51:52 CET 2018"
            }, {
                "category": 3,
                "username": "JustAnotherRegularUsername",
                "title": "Woah, this user also failed!",
                "description": "Something happened [...]",
                "time": "Mon Dec  3 09:48:12 CET 2018"
            }
        ]
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "items": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "category": {
                            "type": "number"
                        },
                        "username": {
                            "type": "string"
                        },
                        "title": {
                            "type": "string"
                        },
                        "description": {
                            "type": "string"
                        },
                        "time": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
    ```

## Statistics [/statistic]
### Global Roast Count Statistics [/statistic/roast/count{?start}{?end}{?interval}]
#### Get Count of Roasts as Time series [GET]
+ Request Roast count statistics as time series
    + Parameters
        + start: "Mon Dec  3 10:31:52 CET 2018" (string) - Start date.
        + end: "Mon Dec  3 10:51:52 CET 2018" (string) - End date.
        + interval: 10m (string) - Interval for each data point in time unit, such as 'm', 'h' etc.
+ Response 200 (application/json)
    + Body
    ```json
    [
 {
  "timestamp": "2018-12-17T23:10:00Z",
  "count": 123
 },
 {
  "timestamp": "2018-12-17T23:00:00Z",
  "count": 32
 }
    ]
    ```

### User Specific Roast Count Statistics [/statistic/{username}/roast/count{?start}{?end}{?interval}]
#### Get Count of Roasts as Time series for Specific User [GET]
+ Request Roast count statistics as time series for specific user
    + Parameters
        + start: "2018-12-17T23:00:00Z" (string) - Start date.
        + end: "2018-12-17T23:10:00Z" (string) - End date.
        + interval: 10m (string) - Interval for each data point in time unit, such as 'm', 'h' etc.
+ Response 200 (application/json)
    + Body
    ```json
    [
 {
  "timestamp": "2018-12-17T23:10:00Z",
  "count": 123
 },
 {
  "timestamp": "2018-12-17T23:00:00Z",
  "count": 32
 }
    ]
    ```

### Global Language Statistics [/statistic/language{?start}{?end}{?interval}]
#### Get For All Languages [GET]
+ Request statistics for all languages
    + Parameters
        + start: "Mon Dec  3 10:31:52 CET 2018" (string) - Start date.
        + end: "Mon Dec  3 10:51:52 CET 2018" (string) - End date.
        + interval: 600 (number) - Interval for each data point in seconds.
+ Response 200 (application/json)
    + Body
    ```json
    {
        "python": [
            {
                "timestamp": "Mon Dec  3 10:41:52 CET 2018",
                "errors": 5,
                "warnings": 7,
                "rows": 512
            },
            {
                "timestamp": "Mon Dec  3 10:31:52 CET 2018",
                "errors": 5,
                "warnings": 7,
                "rows": 512
            }
        ],
        "fortran": [
            {
                "timestamp": "Mon Dec  3 10:41:52 CET 2018",
                "errors": 1125,
                "warnings": 127,
                "rows": 5
            },
            {
                "timestamp": "Mon Dec  3 10:31:52 CET 2018",
                "errors": 5943,
                "warnings": 712,
                "rows": 3
            }
        ]
    }
    ```


### Global Language Specific Statistics [/statistic/language/{language}{?start}{?end}{?interval}]
#### Get For Specific Languages [GET]
+ Request statistics for specific languages
    + Parameters
        + start: "Mon Dec  3 10:31:52 CET 2018" (string) - Start date.
        + end: "Mon Dec  3 10:51:52 CET 2018" (string) - End date.
        + interval: 600 (number) - Interval for each data point in seconds.
+ Response 200 (application/json)
    + Body
    ```json
    [
        {
            "timestamp": "Mon Dec  3 10:41:52 CET 2018",
            "errors": 5,
            "warnings": 7,
            "rows": 512
        },
        {
            "timestamp": "Mon Dec  3 10:31:52 CET 2018",
            "errors": 5,
            "warnings": 7,
            "rows": 512
        }
    ]
    ```

### User Language Statistics [/statistic/{username}/language{?start}{?end}{?interval}]
#### Get For All Languages For User [GET]
+ Request statistics for all languages for user
    + Parameters
        + start: "Mon Dec  3 10:31:52 CET 2018" (string) - Start date.
        + end: "Mon Dec  3 10:51:52 CET 2018" (string) - End date.
        + Interval: 600 (number) - Interval for each data point in seconds.
+ Response 200 (application/json)
    + Body
    ```json
    {
        "python": [
            {
                "timestamp": "Mon Dec  3 10:41:52 CET 2018",
                "errors": 5,
                "warnings": 7,
                "rows": 512
            },
            {
                "timestamp": "Mon Dec  3 10:31:52 CET 2018",
                "errors": 5,
                "warnings": 7,
                "rows": 512
            }
        ],
        "fortran": [
            {
                "timestamp": "Mon Dec  3 10:41:52 CET 2018",
                "errors": 1125,
                "warnings": 127,
                "rows": 5
            },
            {
                "timestamp": "Mon Dec  3 10:31:52 CET 2018",
                "errors": 5943,
                "warnings": 712,
                "rows": 3
            }
        ]
    }
    ```

### User Language Specific Statistics [/statistic/{username}/language/{language}{?start}{?end}{?interval}]
#### Get For Specific Languages For User [GET]
+ Request statistics for specific languages for user
    + Parameters
        + start: "Mon Dec  3 10:31:52 CET 2018" (string) - Start date.
        + end: "Mon Dec  3 10:51:52 CET 2018" (string) - End date.
        + interval: 600 (number) - Interval for each data point in seconds.
+ Response 200 (application/json)
    + Body
    ```json
    [
        {
            "timestamp": "Mon Dec  3 10:31:52 CET 2018",
            "errors": 5,
            "warnings": 7,
            "rows": 512
        },
        {
            "timestamp": "Mon Dec  3 10:41:52 CET 2018",
            "errors": 5,
            "warnings": 7,
            "rows": 512
        }
    ]
    ```
