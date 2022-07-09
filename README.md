# About this project

I wanted to experiment with multiple things

- [golang](https://go.dev/), language that I recently discovered and really appreciate
- [GRPC](https://grpc.io/) nice replacement for JSON if you ask me, perfect for inter-service communication
- microservices architecture because I never worked with them

# Goal

Create a platform where users can create shows (some kind of room) and post comments in it.

I will not bother with registration and authentication.

# What is missing

- For now, I didn't make a service in charge of the shows, the gateway is doing all of the work.

- Comments do no check if a show exists.

- Everything is currenly in local state. This needs to be moved into some database in order to be able to scale the servers horizontally.

# Services

## Redis

Used for pubsub. It is shared by all services but can be instancied for each service independently.

## Comments

Microservice handling show comments. Very simple, only exposed through GRPC.

## Gateway

Service exposing a GraphQL HTTP+WS endpoint and a GraphiQL playground.
Communicates with other services (**Comments** in this case) using GRPC.

# Technical choices

I chose to do inter-service communication using GRPC.  
Service internal communication are done using redis pubsub.

Another solution would be to use pubsub for inter-service communication. In this scenario, the gateway would not use streaming GRPC but would subscribe to pubsub instead.
