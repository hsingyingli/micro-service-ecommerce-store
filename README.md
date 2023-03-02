# micro-service-ecommerce-store
--- 
Simple ecommerce store app built with microservices-based architecture. Technical Stack includes golang(gin), react(nextjs, tailwindCSS), rabbitMQ, GRPC, and docker.

## Table of contents
* [Technical Stack](#technical-stack)
* [Event flow](#event-flow)
* [How to use](#how-to-use)
* [Todo List](#todo)

## Technical Stack 
* [Next.js](https://nextjs.org/): Next.js enables you to create full-stack web applications by extending the latest React features, and integrating powerful Rust-based JavaScript tooling for the fastest builds.
* [TailwindCSS](https://tailwindcss.com/): A utility-first CSS framework packed with classes
* [Golang-gin](https://github.com/gin-gonic/gin): Gin is a web framework written in Go. It features a martini-like API with performance that is up to 40 times faster.
* [RabbitMQ](https://www.rabbitmq.com/): RabbitMQ is the most widely deployed open source message broker.
* [GRPC](https://grpc.io/): A high performance, open source universal RPC framework
* [Docker](https://www.docker.com/): Accelerate how you build, share, and run modern applications.

## Event flow 
![Event flow](./public/eventflow.png)

## How to use
***require docker and make installed***

1. git clone 

```sh
git clone https://github.com/hsingyingli/micro-service-ecommerce-store.git
cd ./micro-service-ecommerce-store
```
2. Start all the services defined in a Docker Compose file.

```sh
docker-compose up -d --build
```

3. Init required DB
***Need to wait until all database has been init (docker ps)***

```
make init_db
```

4. Open client service and start
```
http://localhost:3000
```

## TODO
- [x] Dev
    - [x] create Makefile
        - [x] init db

    - [x] create docker compose file
        - [x] auth service and db
        - [x] product service and db
        - [x] cart service and db
        - [x] rabbitMQ 
        - [x] declare container name

- [ ] Client
    - [x] login Page
    - [x] SignUp Page
    - [ ] Layout
        - [ ] Header
            - [x] Account Menu
            - [ ] Cart (if login      )
            - [x] Own product (if login)
    - [ ] Home Page (Practice infinite scroll)
    - [ ] Sell Page
        - [x] create product 
        - [x] list product 
        - [x] filter product
        - [ ] edit product
        - [x] delete product 
        - [x] link to product page
    - [ ] Product Page 
        

- [ ] Auth Service 
    - [x] Login User API
        - [x] PASETO Token Maker
        - [x] Create Access and Refresh Token
        
    - [x] Logout User API
        - [x] Clear refresh token in cookie

    - [x] Renew Access Token

    - [x] GRPC for auth other micro service
        - [x] define proto file
        - [x] listen and serve grpc server

    - [x] Setup Rabbit MQ
        - [x] Publish User.* Message

- [ ] Product Service
    - [x] Auth middleware which ask auth service for verify access token via GRPC 
    - [x] List Own product 
    - [x] delete own product
    - [x] list all product 
    - [ ] Setup Rabbit MQ
        - [x] Publish Product.* Message

- [ ] Cart Service 
    - [x] Setup Rabbit MQ
        - [x] Receive User.* message from auth serivce via rabbit MQ
        - [x] Receive Product.* message from product serivce via rabbit MQ

    - [ ] Create Cart API
