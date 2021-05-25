# rest-demo

This Golang application is REST API.

Frameworks used:

- Echo to build the website: https://github.com/labstack/echo
- XORM to populate the database: https://xorm.io/


## Setup

### Configuration file

Create a rest_task.ini configuration file and place it wherever you want. 
Configuration should look like this
```
[database]
host = localhost
name = postgres
password =postgres

[main]
secret_key = "RANDOMKEYGOESHERE"
```

### Export config path

put this line into ~/.bashrc file

```
export REST_TASK_SETTINGS=path/to/config/reset_task.ini 
```

**change** ``path/to/config/`` with your config path.

### Run migration

in order to run migration you need to install [golang-migrate/migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

after installing run this command from project's directory
```
 migrate -path ./contrib/migrations/ -database "postgres://USERNAME:@localhost:5432/DBNAME?sslmode=disable" up
```

**change** USERNAME and DBNAME with you username and dbname

### Build 
Run make command it will take care of rest
```
make
```

### Help

``./bin/rest --help``

### Serve api

``./bin/rest serve_api``

