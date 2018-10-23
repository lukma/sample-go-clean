# sample-go

This project explain how to implement rest at golang application project.

## Table of Contents

- [Dependencies](#dependencies)
- [Structure](#structure)
- [How To Run](#how-to-run)

## Dependencies

This project dependencies:

* [godotenv](https://github.com/joho/godotenv) is a dependency to handle multiple env.
* [gin-gionic](https://github.com/gin-gonic/gin) is a dependency to handle request and response http.
* [gin-gionic-cors](https://github.com/gin-contrib/cors) is a dependency as middleware cors for gin gionic.
* [mgo](https://gopkg.in/mgo.v2) is a dependency to handle mongoDB transactional.
* [jwt](https://github.com/dgrijalva/jwt-go) is a dependency to handle auth using Json Web Token.
* [bcrypt](https://golang.org/x/crypto/bcrypt) is a dependency to handle password encryption.

## Structure

This project structure:

```
domain/
data/
services/
```

## How To Run

* Install docker.
* Install and run [mongodb](https://docs.mongodb.com/manual/).
* Build an image `docker build -t sample-go:latest .`.
* Running image to container `docker run --publish 8080:8080 -e ENV='docker' --name sample-go --link mongo:mongo -d sample-go:latest`.