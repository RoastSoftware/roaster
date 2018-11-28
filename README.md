Project Roaster
===============
> Automatically grades _your_ code!

[Roast](https://roast.software) analyzes and grades your code. The user can
register and follow their personal programming progress and compare it with
friends! You could compare it to music scrobbling, but for code.

Personal statistics is generated by running static code analysis on the uploaded
code - data such as number of errors, what type of errors and even code style
problems is collected. The statistics can then be viewed in a feed where both
your own and your friends progress is listed.

Global statistics like most common errors and number of rows analyzed is
displayed on a page that can be viewed by everyone.

## Project status
| Build status | Test coverage |
|:------------:|:-------------:|
| [![Build Status](https://travis-ci.org/LuleaUniversityOfTechnology/2018-project-roaster.svg?branch=master)](https://travis-ci.org/LuleaUniversityOfTechnology/2018-project-roaster) | [![Coverage Status](https://coveralls.io/repos/github/LuleaUniversityOfTechnology/2018-project-roaster/badge.svg)](https://coveralls.io/github/LuleaUniversityOfTechnology/2018-project-roaster) |

## Set up Roaster (for developers)
### Prerequisites
 * Go >= v1.11
 * NodeJS w/npm
 * Docker

### Initial setup
We use PostgreSQL and Redis as our databases. The simplest way to set them up is
using the official Docker images.

Run the PostgreSQL Docker image with (exposed at port: `5432`), replace "<db password>" with any password you want:
```
docker run --name roaster-postgresql -e POSTGRES_PASSWORD=<db password> -p 5432:5432 -d postgres
```

And run the Redis Docker image with (exposed at port: `6379`, non-persistent, w/o password):
```
docker run -it --rm -p 6379:6379 --name roaster-redis -d redis
```

The PostgreSQL database can be configured with pgAdmin4:
```
docker pull dpage/pgadmin4
docker run -p 80:80 \
	--name roaster-pgadmin4 \
	-e "PGADMIN_DEFAULT_EMAIL=user@domain.com" \
	-e "PGADMIN_DEFAULT_PASSWORD=SuperSecret" \
	-d dpage/pgadmin4 \
	--link roaster-pgadmin4:roaster-postgresql
```

Access the pgAdmin4 website at: http://localhost

If you are using windows, the easiest thing to do is to install pgAdmin4 through the installer on their website.
Start the application, choose new server, under the connection tab, type the ip that is shown when running the command "docker-machine ip" in your docker terminal. Use the port "5432", name the server whatever you like, and the password is the one you set in the first step in this guide.

The PostgreSQL database should be setup with a new database with a schema called
`roaster` and a user who only has privilege to that schema.

Clone the Roaster repository:
```
git clone git@github.com:LuleaUniversityOfTechnology/2018-project-roaster.git
```

And enter the directory:
```
cd 2018-project-roaster
```

### Backend (roasterd)
The backend is written in Go and therefore requires that Go is installed
correctly on your machine. Also, a version >= v1.11 is required for Go modules
which is used in this project.

Make sure you enable Go modules if you are using Go v1.11:
```
export GO111MODULE=on
```

#### Requirements (see Initial setup above)
 * A PostgreSQL database must be set up with a scheme named `roaster` that the
provided user has access to.
 * A Redis server on localhost without any password listening on port `6379`.

#### Start web server (development mode)
To run the `roasterd` server in development mode, simply run:
```
go run github.com/LuleaUniversityOfTechnology/2018-project-roaster/cmd/roasterd \
	-dev-mode \
	-database-source="user=<db user> dbname=<db name> password=<db password> sslmode=disable"
```

#### Start web server (production mode)
The `roasterd` server requires two cryptographically secure random keys. They
can be generated and set to their environment variables with:
```
export SESSION_KEY=$(LC_ALL=C tr -dc '[:alnum:]' < /dev/urandom | head -c32)
export CSRF_KEY=$(LC_ALL=C tr -dc '[:alnum:]' < /dev/urandom | head -c32)
```
This will produce two unique keys with 32 bytes of random characters.
Make sure that you do **NOT** use the same key for both environment variables.

Then run the `roasterd` server with:
```
go run github.com/LuleaUniversityOfTechnology/2018-project-roaster/cmd/roasterd \
	-database-source="user=<db user> dbname=<db name> password=<db password> sslmode=disable"
```

If you do not set the environment variables above you can expect a crash from
the server telling you that you do not have enough amount of bytes for the two
keys.

#### A note about `sslmode`
The `sslmode` key can be set to one of:
 * `require` - Use SSL/TLS w/o verification
 * `verify_ca` - Verify CA for SSL/TLS, but not the hostname
 * `verify_full` - Verify both CA and hostname for SSL/TLS
 * `disable` - No SSL/TLS

### Frontend
The frontend is written in TypeScript, HTML and CSS. Everything is packaged and
compiled using webpack.

First you have to be located in the `www/` folder:
```
cd www/
```

Then, install the required dependencies with:
```
npm install
```

Finally, start the frontend autobuild with:
```
npm start
```

Now, everytime you make any change to the frontend, everything will
automatically recompile and can be accessed from: `http://localhost:5000`
(hosted by the `roasterd` backend).
