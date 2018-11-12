./model
=======
This folder contains models that interact with the PostgreSQL database.

The `model.go` file contains the code that opens and initializes the PostgreSQL
database. Each model category is declared in their respective Go file directly
in this folder. The `forwardengineer` folder holds the PostgreSQL schema and a
generated Go file that holds a copy of all queries in the schema which can be
used by `model.go` to forwardengineer the database. An example file structure:

```
.
├── forwardengineer
│   ├── schema.go
│   └── schema.sql
├── model.go
└── user.go
```
