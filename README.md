# Waco Test

This is a Go project created as a test for the job application at "Waco". The project is an application that demonstrates a CRUD (Create, Read, Update, Delete) functionality for user management, authentication and API consumption.

### Accessing the Public Version

The application is also available in railway for public use. You can access the live version by clicking [https://waco-test-go-production.up.railway.app](https://waco-test-go-production.up.railway.app).

### Documentation

For more detailed information on how the project works, you can refer to the <a href="https://documenter.getpostman.com/view/12519140/2s9Xy6qVcC" target="_blank" rel="noopener">Project Documentation</a>. The documentation contains detailed explanations of the endpoints, and other relevant information.

### Libraries Used

The following libraries are used in this project:

- [gin-gonic/gin v1.9.1](https://github.com/gin-gonic/gin)
- [go-sql-driver/mysql v1.7.1](https://github.com/go-sql-driver/mysql)
- [golang-jwt/jwt/v5 v5.0.0](https://github.com/golang-jwt/jwt)
- [golang.org/x/crypto v0.12.0](https://pkg.go.dev/golang.org/x/crypto)

These libraries provide essential functionality for various aspects of the project, such as routing, database connectivity, JWT authentication, and cryptographic operations.

## Getting Started

To run this project locally on your machine, follow these steps:

### Prerequisites

1. Make sure you have Go (1.20) installed on your system.
2. Install MySQL and create a new database for this project.

### Installation

1. Clone the repository to your local machine:

```bash
git clone https://github.com/sebas7603/norte-digital-test-laravel.git
cd norte-digital-test-laravel
```

2. Install project dependencies using Go modules:

```bash
go mod tidy
```

3. Create a copy of the `.env.example` file and rename it to `.env`. Update the database configuration in the `.env` file with your MySQL credentials:

```dotenv
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=waco-test
DB_USERNAME=waco-test
DB_PASSWORD=your_database_password

JWT_SECRET=your_jwt_secret
```

### Running the Application

You can now run the app:

```bash
go run main.go
```

The application should be accessible at `http://localhost:8080` in your web browser and ready to process requests from postman.

## Contact

If you have any questions or need further assistance, you can contact me via email at <a href="mailto:sebastianutpae@gmail.com" target="_blank" rel="noopener">sebastianutpae@gmail.com</a>.

Thank you for reviewing my test project!
