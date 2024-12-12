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
```

### Description
- User Authentication and Registration
  - New users can register to add new products
  - Authorize users with jwt
  - Secure some routes and validate them with jwt provided from cookies
- Products
  - All user can create product and manage them (crud operations)

## Deployment links
...

# TODO 
- image upload
- add cart feature
- add orders feature
- reverse proxy
- swagger documentation
