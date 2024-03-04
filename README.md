# API Using Golang

Final Task of BTPN Syariah Fullstack Developer Project-Based Internship Program by Fikri Naufal Hamdi.

## Program Description

This program is an API developed using the GoLang programming language, utilizing the Gin Gonic framework and Gorm ORM. Based on the given case, this API is created for a mobile banking application to enhance user engagement by providing a feature to upload and delete user profile pictures. It includes endpoints for user and photo management with employment of JWT authentication system to secure sensitive access.

## Requirements

1. Golang
2. PostgreSQL
3. [*Optional*] CompileDaemon
4. Several Modules
    - "github.com/gin-gonic/gin"
    - "gorm.io/gorm"
    - "github.com/golang-jwt/jwt/v5"
    - "github.com/asaskevich/govalidator"
    - "gorm.io/driver/postgres"
    - "golang.org/x/crypto/bcrypt"
    - "github.com/nedpals/supabase-go"
    - "github.com/githubnemo/CompileDaemon"
    
## How to Run Project

1. Clone this repository
    ```
    git clone https://github.com/fiknaufalh/task-5-pbi-btpns-fikrinaufalhamdi.git
    ```

2. Rename `.env-example` to `.env` for using my personal database. If you want to use your config, adjust the environment variables.


3. Go to the terminal and navigate to this project directory
    ```
    cd task-5-pbi-btpns-fikrinaufalhamdi
    ```

4. Install the dependency libraries provided in `go.mod`
   ```
   go mod download
   ```

5. Start the server
   ```
   go run main.go
   ```

   [_Optional_] If you want to develop this project further and automatically run the server after change in the codebase, you can using CompileDaemon

   ```
   go install github.com/githubnemo/CompileDaemon
   ```

   And start the server using this command
   ```
   compiledaemon --command="./profile-picture-api"
   ```


## API Documentation

### Root 
* Method: `GET`
* Endpoint: `/`
* Response: 
    - Status Code: `200 OK`
    - Body:
        ```
        {
            "author": "Fikri Naufal Hamdi",
            "message": "Welcome to Profile Picture API!"
        }
        ```

### User Registration
* Method: `POST`
* Endpoint: `/users/register`
* Request: 
    - Body:
        ```
        {
            "username": "username",
            "email": "user@example.com",
            "password": "password"
        }
        ```
* Response: 
    - Status Code: `201 Created`
    - Body:
        ```
        {
            "message": "User created",
            "status": "Succeed"
        }
        ```

### User Login
* Method: `POST`
* Endpoint: `/users/login`
* Request: 
    - Body:
        ```
        {
            "email": "user@example.com",
            "password": "password"
        }
        ```
* Response: 
    - Status Code: `202 Accepted`
    - Body:
        ```
        {
            "message": "Login Success",
            "status": "Succeed"
        }
        ```
    - Cookie
        - Authorization (token): Bearer Token Value

### User Logout
* Method: `POST`
* Endpoint: `/users/logout`
* Response: 
    - Status Code: `200 OK`
    - Body:
        ```
        {
            "message": "Logout Success",
            "status": "Succeed"
        }
        ```

### Get User by ID
* Method: `GET`
* Endpoint: `/users/{userID}`
* Authorization
    - Bearer token
* Response: 
    - Status Code: `200 OK`
    - Body:
        ```
        {
            "data": {
                "id": 2,
                "username": "username",
                "email": "user@example.com",
                "password": "$2a$10$f4PzOPNgogyrKCLovynXtemkSlkgS8Fxywv6LofROFprbbBzQlBD2",
                "photos": null,
                "created_at": "2024-03-04T17:04:22.38024+07:00",
                "updated_at": "2024-03-04T17:04:22.38024+07:00"
            },
            "message": "Fetch a user",
            "status": "Succeed"
        }
        ```

### Update User by ID
* Method: `POST`
* Endpoint: `/users/{userID}`
* Authorization
    - Bearer token
* Request: 
    - Body:
        ```
        {
            "username": "username",
            "email": "user@example.com",
            "password": "new_password"
        }
        ```
* Response: 
    - Status Code: `200 OK`
    - Body:
        ```
        {
            "message": "Success update data",
            "status": "Succeed"
        }
        ```

### Delete User by ID
* Method: `DELETE`
* Endpoint: `/users/{userID}`
* Authorization
    - Bearer token
* Request: 
    - Body:
        ```
        {
            "username": "username",
            "email": "user@example.com",
            "password": "new_password"
        }
        ```
* Response: 
    - Status Code: `200 OK`
    - Body:
        ```
        {
            "message": "Success delete data",
            "status": "Succeed"
        }
        ```

### Get All Photos
* Method: `GET`
* Endpoint: `/photos`
* Authorization
    - Bearer token
* Response: 
    - Status Code: `200 OK`
    - Body:
        ```
        {
            "data": [
                {
                    "id": 1,
                    "title": "Sample Photo",
                    "caption": "This is a test photo.",
                    "photo_url": "https://example.com/sample.jpg",
                    "user_id": 1,
                    "created_at": "2024-03-03T00:58:13.857809+07:00",
                    "updated_at": "2024-03-03T00:58:13.857809+07:00"
                }
            ],
            "message": "Fetch all user's photos",
            "status": "Succeed"
        }
        ```

### Create Photo
* Method: `POST`
* Endpoint: `/photos`
* Authorization
    - Bearer token
* Request: 
    - Body:
        ```
        {
            "title": "Rakamin Profile",
            "caption": "As the blue as the sky",
            "photo_url": "bit.ly/RakaminProfile"
        }
        ``` 
* Response: 
    - Status Code: `201 Created`
    - Body:
        ```
        {
            "message": "Success create photo",
            "status": "Succeed"
        }
        ```

### Update Photo
* Method: `PUT`
* Endpoint: `/photos/{photoID}`
* Authorization
    - Bearer token
* Request: 
    - Body:
        ```
        {
            "title": "Rakamin Profile",
            "caption": "As the blue as the sky, right!",
            "photo_url": "bit.ly/RakaminProfile"
        }
        ``` 
* Response: 
    - Status Code: `200 OK`
    - Body:
        ```
        {
            "message": "Success update photo",
            "status": "Succeed"
        }
        ```

### Delete Photo
* Method: `DELETE`
* Endpoint: `/photos/{photoID}`
* Authorization
    - Bearer token
* Response: 
    - Status Code: `200 OK`
    - Body:
        ```
        {
            "message": "Success delete photo",
            "status": "Succeed"
        }
        ```

---
***<p style="text-align: center;">Made by Fikri Naufal Hamdi</p>***