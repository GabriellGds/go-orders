# Orders API

This is a straightforward orders API developed in Golang. Go-Chi is utilized as the API router, SQLX for database operations, POSTGRES as the database, and Swagger for API documentation and testing. 
The project incorporates a Dockerfile to simplify its deployment in containers.


## This project was developed with the following technologies

- [Golang](https://go.dev/): The programming language used in this project.
- [Go-Chi](https://github.com/go-chi/chi): Used for route management.
- [Swagger](https://swagger.io/): Integrated for API documentation and testing.
- [DockerTest](https://github.com/ory/dockertest): A Go library utilized to automate the configuration of the testing environment, specifically for conducting integration tests.
- [SQLX](https://github.com/jmoiron/sqlx): A Go library that simplifies interaction with SQL databases, providing enhanced data binding capabilities.
- [Mockgen](https://github.com/uber-go/mock): A library in Go used to automate the creation of mocks in unit tests
- [Postgres](https://www.postgresql.org/): Database management system used in the project.
- [Docker](https://www.docker.com/products/docker-desktop/): Used to easily run the application, allowing it to run seamlessly across different environments.

## Prerequisites

- [Go](https://go.dev/): The Go programming language.
- [Docker](https://www.docker.com/products/docker-desktop/): Docker is used to build and run the application in a containerized environment.

## How to run the application
 **Clone the repository** 
  ```
  git clone https://github.com/GabriellGds/go-orders.git
```

**Navigate to the project directory**
```
 cd go-orders
```
 **Run the services in the docker-compose.yml**
  ```
  docker-compose up -d
  ```
 **To stop the application, remove containers, and volumes defined in the docker-compose.yml**
   ```
  docker-compose down
  ```

## Usage
After the API is running, the Swagger UI enables interaction with the endpoints for operations such as search, creation, modification,
and removal of users, items, and orders. The application will be accessible at http://localhost:5000/swagger/index.html.

This API uses token-based authentication. To access protected routes, you need to follow these steps:

1. **Create a User**: In the Swagger UI, look for the /user endpoint. You should see a ‘Try it out’ button.
Click on it, and you will be able to enter the required user details. After entering the details, click on ‘Execute’. 
This will make a POST request to the /user endpoint and create a user.

2. **Login**: After creating a user, navigate to the /login endpoint in the Swagger UI. Again, click on ‘Try it out’, 
and enter your user credentials. Click on ‘Execute’. If the credentials are valid, the API will return an authentication token in the header response.

3. **Use the Token**: To access protected routes, you need to include the authentication token in your requests. 
In the Swagger UI, you can input your token by clicking on the ‘Authorize’ button. A dialog box will appear where you can enter your token.

## Running the Tests

This project contains both unit tests and integration tests. You can run these tests using the following commands

**To run only unit tests**
```
make unit-tests
```

**To run all tests (including unit and integration tests)**
```
make all-tests
```

## License

This project is distributed under the MIT license.

   

