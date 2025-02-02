# Simple E-comerce api with Golang + Gin

# Run application
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

### Description
- User Authentication and Registration
  - New users can register to add new products
  - Authorize users with jwt
  - Secure some routes and validate them with jwt provided from cookies
- Products
  - All user can create product and manage them (crud operations)
- API documentation: localhost:4000/swagger/index.html


## Deployment links
...

# TODO
- add cart feature
- add orders feature
- reverse proxy
- swagger documentation
