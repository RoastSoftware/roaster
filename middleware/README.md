./middleware
============
This folder contains middleware that is applied to every request or specific
routes.

There should only be sub packages that holds different modular middlewares.
An example file structure:

```
.
└── auth
    ├── auth.go
    └── auth_test.go
```
