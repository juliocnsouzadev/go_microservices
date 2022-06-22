# Microservices with Go
## Building highly available, scalable, resilient distributed applications using Go

### Applications
- Front End service, that just displays web pages;
- Authentication service, with a Postgres database;
- Logging service, with a MongoDB database;
- Listener service, which receives messages from RabbitMQ and acts upon them;
- Broker service, which is an optional single point of entry into the microservice cluster;
- Mail service, which takes a JSON payload, converts into a formatted email, and send it out.

### Deployments
- Docker Swarm
- Kubernetes

### Dependencies
- broker-service
```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
```

- authentication-service
```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
go get golang.org/x/crypto/bcrypt
```