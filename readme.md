# Go Fiber Server

This is a simple Go Fiber server that uses a MySQL database and Redis for caching.

## Features
* Fiber web framework for handling HTTP requests
* G ORORM as anM for database operations
* MySQL as the database
* Redis for caching
* Middleware for CORS configuration

## Endpoints

- `GET /customers`: Get all customers
- `GET /customers/:id`: Get a specific customer by ID
- `POST /customers`: Create a new customer
- `PUT /customers/:id`: Update a customer by ID
- `DELETE /customers/:id`: Delete a customer by ID

## Configuration

The server reads its configuration from a `config.yaml` file at the root of the project. The following options are available:

- `app.port`: The port the server listens on (default: `5000`)
- `db.host`: The MySQL database host (default: `localhost`)
- `db.port`: The MySQL database port (default: `3307`)
- `db.username`: The MySQL database username (default: `root`)
- `db.password`: The MySQL database password (default: `1234`)
- `db.database`: The MySQL database name (default: `customers`)

Docker reads its configuration from a `docker-compose.yml` file at the root of the project.
Edit environment: 
- MARIADB_ROOT_PASSWORD=YOUR_PASSWORD
- MARIADB_DATABASE=YOUR_DATABASE_NAME
If you want to author user you should configure following
- MARIADB_USER=YOUR_USERNAME
IF you don't want to use password you should configure
- MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=yes
