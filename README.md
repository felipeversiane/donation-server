# Donation Server

This project provides a simple HTTP API for processing donations. It is written in Go and uses PostgreSQL for persistence. A complete development environment is available via Docker Compose, which also spins up a MinIO instance for file storage.

## Getting Started

### 1. Load environment variables

Copy the sample environment file and adjust values if needed:

```bash
cp .env.example .env
```

The application uses [godotenv](https://github.com/joho/godotenv) so variables from the `.env` file are loaded automatically when the server starts.

### 2. Build and run with Docker Compose

Make sure Docker and Docker Compose are installed. Run the following command to start the services defined in `compose.local.yml` (PostgreSQL, migrator, server, pgweb and MinIO):

```bash
docker compose -f compose.local.yml up --build
```

Alternatively you can use `compose.dev.yml` which behaves similarly but loads variables from `.env`.

Once the containers are running the API will be available at `http://localhost:8000/api/v1`.

## API Documentation

Interactive Swagger documentation is generated from the `docs` folder. After the server is up open:

```
http://localhost:8000/api/v1/swagger/index.html
```

Use the credentials defined by `HTTP_SERVER_SWAGGER_USER` and `HTTP_SERVER_SWAGGER_PASSWORD` if authentication is enabled.

## Database Diagram

An Entity–Relationship (ER) diagram of the database schema is included in [`er.png`](er.png). It illustrates the tables created by the migrations under the `migrations` directory.

## License

This project is released under the [MIT License](LICENSE).
