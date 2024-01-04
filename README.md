# Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

## `set your .env`
DB_HOST=localhost
DB_PORT=5432
DB_PASSWORD=your_database_password
DB_USER=your_database_user
DB_SSLMODE=disable
DB_NAME=your_database_name

## `used tailwind css for styling`

## `Running the Go Application`

To run the Go application, use the following command:

```bash
go run main.go



// Import the Bootstrap module for initializing the application.
import "github.com/Aysha-Ramzy/Pharmacy-App/bootstrap"

// Import the Repository module for handling data storage and retrieval.
import "github.com/Aysha-Ramzy/Pharmacy-App/repository"

// Import the Fiber web framework for building web applications.
import "github.com/gofiber/fiber/v2"



// Import the required packages for database interaction using Gorm and PostgreSQL.
import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)
