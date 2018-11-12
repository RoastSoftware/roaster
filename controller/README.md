./controller
============
This folder contains routes and controllers for each of the web pages.

## File structure
The folder should contain a `controller.go` start point for the routes that
initializes all the routes defined in the sub packages. An example structure:

```
.
├── controller.go
├── controller_test.go
├── login
│   ├── login.go
│   └── login_test.go
├── logout
│   ├── logout.go
│   └── logout_test.go
├── static
│   ├── static.go
│   └── static_test.go
└── user
    ├── user.go
    └── user_test.go
```
