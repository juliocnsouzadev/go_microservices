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
