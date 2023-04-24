## Description
The goal was to implement a Go web server emulating the backend of the CRM application. The application handles the basic GET, POST, PUT, and DELETE operations for the customer database represented with struct and saved in a map data type.

## Main bits of the project
* main
* handlers
* modules

## Structure of the project
```
.
crmbackend/
├── README.md
├── go.mod
├── go.sum
├── handlers
│   └── handlers.go
├── main.go
├── main_test.go
└── models
    └── model.go
```

## Features
The project represents the backend of a customer relationship management (CRM) web application. The server supports the following functionalities as users interact with the app via some user interface. User can make HTTP requests using utilities like [Postman](https://www.postman.com/) or [cURL](https://curl.se/)

| Operation                             | Endpoint          |
| ------------------------------------- | ----------------- |
| Adding a customer                     | `/customers`      |
| Getting data for a single customer    | `/customers/{id}` |
| Getting a list of all customers       | `/customers`      |
| Updating a customer's information     | `/customers/{id}` |
| Removing a customer                   | `/customers/{id}` |

The application runs in memory, and the customer's data is not persistent. After stopping the server, all the data is lost. The application stores the customers' data in a Go map data type, which serves as the database of the CRM. The application is a backend API with no front end but a static welcome page.

## Requirements
To run the project and tests you will need
* [Go Programming language](https://go.dev/doc/install)
* [GOOGLE/UUID](github.com/google/uuid)
* [GORILLA/MUX](github.com/gorilla/mux)

GOOGLE/UUID and GORILLA/MUX are third-party Go packages that need to be installed.

To install them, open a terminal and navigate to the root directory of the CRM Backend project and run the following commands.

```
go get github.com/gorilla/mux
go get github.com/google/uuid
```

## Starting and Using the API
To start the application, open a terminal and navigate to the root directory of the CRM Backend project, where the main.go file is located. Then, run the following command:

```
go run main.go
```

This will start the server on port 3000, the default hardcoded port for the server. A port can be changed in the main.go file.

### Run Unit tests
To run application unit tests, navigate to the root directory and run the following command.

```
go test -v
```
