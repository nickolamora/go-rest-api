# go-rest-api

Sample REST API with MySQL written in GO.

## Dependencies

* Gorilla mux - router for go lang http requests
    * `go get -u github.com/gorilla/mux`
* Gorm - ORM library for golang
    * `go get -u gorm.io/gorm`
* Gorm mysql connector
    *  `go get -u gorm.io/driver/mysql`

## Paths

Runs in port 9000.

* GET /customer
* GET /customer/{id}
* POST /customer
  * ```json
    {
        "firstName": "Pepito",
        "lastName": "Morales",
        "email": "my@email.com"
    }

    ```
* PUT /customer/{id}
* DELETE /customer/{id}

## Running the application
* `go build`
* `./go-rest-api`
