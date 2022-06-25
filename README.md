# golang-api-layered
Example REST API implementation with layered architecture by golang.

## Application Architecture
![Application Architecture](https://user-images.githubusercontent.com/2717971/175773583-50be1143-0aa7-4b2c-b420-37d8c5a4f9cf.png)
### Application Architecture
Seperated as three layers by interface.

### Testing
Test is written for each layer using each layer's mock.

## Directory structure
```
## Root directories
├── build/ -- Dockerfile
├── cmd/ -- Dommand file
├── env -- Env file
├── main.go 
├── mock/ -- Mock json files used by testing

## Application directories
└── src
    └── app
        ├── application -- Application logic(usecase, converter, etc)
        ├── config -- Config files
        ├── domain -- Service layer(model, service, api client, etc)
        ├── interfaces -- Client interface layer(API endpoints, middleware, converting to json response, etc)
        └── library -- Common util logic 
```

## External API reference
### Github API
https://docs.github.com/en/rest/reference


