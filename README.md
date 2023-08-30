# Go service skeleton

### Project structure
```
├── Makefile
├── README.md
├── api                             #Any web server realization
│   ├── main.go
│   ├── middleware
│   │   └── cors.go
│   ├── server
│   │   ├── gin.go
│   │   ├── gin_default_routes.go
│   │   ├── response.go
│   │   └── server.go
│   └── service
│       ├── di.go
│       ├── errors.go
│       ├── registar.go
│       └── route_ping.go
├── build                           #Application builds
│   └── config.yaml
├── cmd                             #Console applications realization
│   └── main.go
├── configs                         #Config files
│   └── config.local.yaml
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal                        #Internal application logic
│   ├── processors                  #Interface for external data sources(cmd/api/etc.)
│   │   ├── di.go
│   │   ├── dto.go
│   │   └── example_controller.go
│   ├── entity                      #Entitits (for internal usage only)
│   │   ├── entity.go
│   │   ├── error.go
│   │   └── example.go
│   ├── internal.go
│   ├── presenters                  #Presenters for extrenal usage (outside internal) (api/cmd/etc.)
│   │   └── interface.go
│   ├── repositories                #Internal data repositories (dbs, for example)
│   │   ├── db_provider_pgsql.go
│   │   └── pgsql
│   │       ├── example_repository.go
│   │       ├── migrations
│   │       │   ├── 202201141633500_example_migration.down.sql
│   │       │   └── 202201141633500_example_migration.up.sql
│   │       ├── pgsql.go
│   │       ├── queries
│   │       │   └── example.sql
│   │       └── requester
│   │           ├── db.go
│   │           ├── example.sql.go
│   │           ├── models.go
│   │           └── querier.go
│   └── usecases                       #Application logic
│       └── example
│           ├── interface.go
│           └── service.go
├── pkg                                #Any external packages
│   ├── argumentsresolver              #Input argument resolver
│   │   └── arguments_resolver.go 
│   ├── config                         #Configuration mapping
│   │   ├── api.go
│   │   ├── config.go
│   │   ├── db.go
│   │   └── logger.go
│   ├── logger                         #Logger interface
│   │   └── logger.go
│   └── pkg.go
├── sqlc.sh
└── sqlc.yaml
```

### Project structure description.

#### Project was designed according to [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and must comply with it.

#### Project has next limitation:
* Data that is transferred from external sources (cmd/api/etc.) must be transferred exclusively using DTOs from
  folders **internal/processors** to processors in same folder
* Data that is returned from the system (**internal** folders) should only be returned using
  **internal/presenters**
* Methods and structures in the **internal/processors** folder are used to transform data from outside the system into the system and
  on the contrary. Logic in these methods should not be implemented
* Main application logic should be placed in the **internal/usecases/{usecase} folders**
* _Database interactions (mongo/pgsql/etc.) are placed in the **internal/repositories** folder. Operating on database outside of this
  folders **forbidden**, any interactions with **internal/usecases** must be bound to entity_
* All elements from the **entity** folder must be used only within internal packages

<br/>

#### cli parameters and project configuration
By default, the project will look for the _config.yaml_ file in the directory where the project is launched, this can be
change by passing the config-file argument (for example _--config-file=/tmp/cfg.yml_), any additional
parameters can be written in the _argumentsresolver_ package. You can connect resolver in any part of the project using
argumentsresolver.ArgumentsInterface interface
