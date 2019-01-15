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
+ Response 400 (application/json)
    Bad request with either invalid/missing fields or unparsable data structure.
    + Attributes (Error)
+ Response 409 (application/json)
    Email or username is conflicting with an already registered user.
    + Attributes (Error)

### View/Handle Specific User [/user/{username}]
#### Change User [PATCH]
+ Request change user (application/json)
    + Body
    ```json
    {
        "email": "email@example.com",
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
+ Response 400 (application/json)
    Bad request with either invalid/missing fields or unparsable data structure.
    + Attributes (Error)
+ Response 409 (application/json)
    Email is conflicting with an already registered user.
    + Attributes (Error)

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
+ Response 400 (application/json)
    Bad request with missing username parameter.
    + Attributes (Error)
+ Response 404 (application/json)
    Requested user does not exist.
    + Attributes (Error)

### User Score [/user/{username}/score]
#### Retrieve User Score [GET]
+ Request user score
+ Response 200 (application/json)
    + Body
    ```json
    {
        "score": 123
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "score": {
                "type": "number"
            }
        }
    }
    ```
+ Response 404 (application/json)
    Requested user does not exist.
    + Attributes (Error)

#### Remove User [DELETE]
+ Request remove user
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
+ Response 501 (application/json)
    Not implemented.
    + Attributes (Error)

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
+ Response 401 (application/json)
    Unauthorized request, must be logged in to upload avatar.
    + Attributes (Error)
+ Response 400 (application/json)
    Invalid Content-Type or unreadable data in request.
    + Attributes (Error)
+ Response 413 (application/json)
    Too large file sent.
    + Attributes (Error)
+ Response 415 (application/json)
    Invalid image type, supported formats are: JPEG, PNG and GIF.
    + Attributes (Error)

#### Retrieve User Avatar [GET]
+ Request user avatar
+ Response 200 (image/png)
+ Response 404 (application/json)
    Requested user does not exist.
    + Attributes (Error)

### Followees for User [/user/{username}/followees]
#### Add New Followee [POST]
+ Request add new followee (application/json)
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
    + Body
    ```json
    {
        "followee": "InterestingFollowee"
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "followee": {
                "type": "string",
                "required": true
            }
        }
    }
    ```
+ Response 200
+ Response 400 (application/json)
    Missing followee field or unparsable data structure.
    + Attributes (Error)
+ Response 404 (application/json)
    Requested followee does not exist.
    + Attributes (Error)

#### Retrieve User Followees [GET]
+ Request user followees
+ Response 200 (application/json)
    + Body
    ```json
    {
        "followees": [
            {
                "username":"MyAwesomeFollowee",
                "createTime":"2018-03-04:00:00:00Z"
            }
            {
                "username":"MyAwesomeFollowee2",
                "createTime":"2018-07-04:00:00:00Z"
            }
        ]
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "followees":
            {
                "type": "array",
                "items": {
                    "type": "string"
                }
            }
        }
    }
    ```
+ Response 204
+ Response 400 (application/json)
    Missing username in request.
    + Attributes (Error)
+ Response 404 (application/json)
    Requested user does not exist.
    + Attributes (Error)

#### Remove User Followee [DELETE]
+ Request remove user followee
    + Headers
    ```
    Cookie: roaster_auth=AB32DEAC21A91DE123[...]
    ```
+ Response 200
+ Response 400 (application/json)
    Missing username for lookup.
    + Attribute (Error)

#### Retrieve User Followers [GET]
+ Request user followers
+ Response 200 (application/json)
    + Body
    ```json
    {
        "followers": [
            {
                "username":"MyAwesomeFollower",
                "createTime":"2018-03-04:00:00:00Z"
            }
            {
                "username":"MyAwesomeFollower2",
                "createTime":"2018-07-04:00:00:00Z"
            }
        ]
    }
    ```
    + Schema
    ```json
    {
        "type": "object",
        "properties": {
            "followers":
            {
                "type": "array",
                "items": {
                    "type": "string"
                }
            }
        }
    }
    ```
+ Response 204
+ Response 400 (application/json)
    Missing username for lookup.
    + Attribute (Error)

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
+ Response 400 (application/json)
    Empty username or password field or unparsable data structure.
    + Attribute (Error)
+ Response 401 (application/json)
    Invalid username or password.
    + Attribute (Error)

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
+ Response 204

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
        "username": "willeponken",
        "language": "python3",
        "code": "print('I´m Roastin´ Roger... Roger.', invalid = 'Bönor är ändå rätt gött'",
        "createTime": "2018-12-17T23:00:00Z",
        "score": 12,
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
        ],
        "statistics": {
            "numberOfErrors": 1,
            "numberOfWarnings": 1,
            "linesOfCode": 1
        }
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
            "score": {
                "type": "number"
            },
            "language": {
                "type": "string"
            },
            "code": {
                "type": "string"
            },
            "createTime": {
                "type": "string"
            },
            "statistics": {
            "type": "object",
            "properties": {
                "linesOfCode": {
                    "type": "number"
                },
                "numberOfWarnings": {
                    "type": "number"
                },
                "numberOfErrors": {
                    "type": "number"
                }
            }
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
+ Response 400 (application/json)
    Invalid Roast data sent.
    + Attribute (Error)

## Feed [/feed{?user}{?followees}{?page}{?page-size}]
### Get Feed [GET]
+ Parameters
    + page: 0 (number) - Page of feed to return.
    + `page-size`: 25 (number, optional) - Page size for each page.
        + Default: 25
    + user: willeponken (string, optional) - Display only for user.
    + followees: true (string, optional) - Display only user followees (excluding user).
+ Request a feed (application/json)
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
+ Response 400 (application/json)
    + Attribute (Error)

## Statistics [/statistics]
### Roast Statistics Timeseries [/statistics/roast/timeseries{?start}{?end}{?interval}{?user}{?followees}]
#### Get Statistics for Roasts as Time series [GET]
+ Parameters
    + start: `2018-12-17T23:00:00Z` (string) - Start date according to RFC3339.
    + end: `2018-12-17T23:10:00Z` (string) - End date according to RFC3339.
    + interval: 10m (string) - Interval for each data point in time unit, such as 'm', 'h' etc.
    + user: willeponken (string, optional) - Show statistics for specific user.
    + followees: true (string, optional) - Show statistics for specific user followees (excluding user).
+ Request Roast count statistics as time series
+ Response 200 (application/json)
    + Body
    ```json
    [
 {
  "timestamp": "2018-12-17T23:10:00Z",
  "count": 123,
  "numberOfErrors": 231,
  "numberOfWarnings": 210,
  "linesOfCode": 5901
 },
 {
  "timestamp": "2018-12-17T23:00:00Z",
  "count": 32,
  "numberOfErrors": 5,
  "numberOfWarnings": 7,
  "linesOfCode": 512
 }
    ]
    ```
+ Response 400 (application/json)
    + Attribute (Error)

### Roast Statistics Count [/statistics/roast/count{?user}{?followees}]
#### Get Number of Roasts [GET]
+ Parameters
    + user: willeponken (string, optional) - Show count for specific user.
    + followees: true (string, optional) - Show count for specific user followees (excluding user).
+ Request Number of Roasts
+ Response 200 (application/json)
    + Body
    ```json
 {
  "count": 321
 }
    ```
+ Response 400 (application/json)
    + Attribute (Error)

### Roast Statistics Lines of Code [/statistics/roast/lines{?user}{?followees}]
#### Get Lines of Code Analyzed [GET]
+ Parameters
    + user: willeponken (string, optional) - Show lines of code for specific user.
    + followees: true (string, optional) - Show lines of code for specific user followees (excluding user).
+ Request Lines of code analyzed
+ Response 200 (application/json)
    + Body
    ```json
 {
  "lines": 321
 }
    ```
+ Response 400 (application/json)
    + Attribute (Error)

### Search [/search{?query}]
#### Search for content [GET]
+ Parameters
    + query: willeponken%2Fcool+guy (string) - Search query (percent-encoded).
+ Request search results for query
+ Response 200 (application/json)
    + Body
    ```json
    [
    {
        "category": 0,
        "title": "User willeponken",
        "description": "A really cool guy, named willeponken",
        "url": "/user/willeponken"
    },
    {
        "category": 0,
        "title": "User coolguy",
        "description": "A really willeponken, named coolguy",
        "url": "/user/coolguy"
    }
    ]
    ```
    + Schema
    ```json
    {
        "type": "array",
        "items": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
    ```
+ Response 204
+ Response 400 (application/json)
    + Attribute (Error)

# Data Structures
## Error (object)
+ message: `An error message` (string)
