# Simple E-comerce api with Golang + Gin

### Tech Stack

- Golang 1.23.4
- Gin - Api framework
- PostgreSQL - Relational db
- Gorm - ORM
- JWT - JSON web token authorization
- Docker - Application containerization

### Description

- User Authentication and Registration
  - New users can register to add new products
  - Authorize users with jwt
  - Secure some routes and validate them with jwt provided from cookies
- Products
  - All user can create product and manage them (crud operations)
- API documentation: localhost:4000/swagger/index.html

## Run application

```
# run locally
/> make run start
# run with container (build image, run with compose)
/> docker build -t myapp .
/> docker run --rm -p 4000:4000 myapp
Or
/> docker-compose up
#  docker-compose up -d --force-recreate
/> swag init
```

- Build with vendor (for deployment)

```bash
/> go build -o go-ecom -mod=vendor main.go
```

# Use vendor for used packages

```
/> go mod tidy
/> go mod vendor
/> go build -mod=vendor
/> go run -mod=vendor .
```

## Deployment links

Documentation - https://goecom.35.171.20.164.nip.io/swagger/index.html

# TODO

- add cart feature
- add orders feature
- reverse proxy
