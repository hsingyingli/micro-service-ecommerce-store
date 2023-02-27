# micro-service-ecommerce-store

## TODO

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

- [ ] Product Service
    - [x] Auth middleware which ask auth service for verify access token via GRPC 
    - [x] List Own product 
    - [x] delete own product
    - [x] list all product 




