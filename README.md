# Project Name

## Overview
This project is a web application built using Go and the Chi router. It includes basic routing, form handling, and middleware integration.

## Features
- Basic routing with Chi
- Form validation
- Middleware for request recovery and IP address logging
- Static file serving

## Prerequisites
- Go 1.16 or later

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/yourproject.git
    cd yourproject
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage
1. Run the application:
    ```sh
    go run cmd/web/*.go
    ```

2. Open your browser and navigate to `http://localhost:8080`.

## Project Structure
- `cmd/web/`: Contains the main application code.
    - `forms.go`: Form handling and validation.
    - `routes.go`: Route definitions and middleware setup.
    - `routes_test.go`: Tests for route existence.
- `static/`: Directory for static assets.
- `.gitignore`: Git ignore file for the project.

## Testing
Run the tests using the following command:
```sh
go test ./...