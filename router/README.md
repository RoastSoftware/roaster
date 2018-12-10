./controller
============
This folder contains routes and controllers for each of the web pages.

## File structure
The folder should contain a `router.New` start point for the routes that
initializes all the routes defined in the handler sub package. An example structure:

```
.
├── route
│   ├── avatar.go
│   ├── handler.go
│   ├── roast.go
│   ├── session.go
│   ├── static.go
│   └── user.go
├── router.go
└── router_test.go
```
