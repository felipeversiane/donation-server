# Project Roadmap

This document expands on the previous TODO list and tracks the major features planned for the donation server.

## 1. Structured logger with file rotation
- Provide a structured logging solution using `slog`.
- Support log rotation on disk through `lumberjack` with configurable size, backups and retention.
- Allow logs to be written to both a file and stdout based on configuration.
- Add request scoped values like request ID and user ID to each log entry.

## 2. Core architecture with ports, value objects, entities and errors
- Define clear domain entities and value objects under `internal/core`.
- Create port interfaces so external adapters can interact with the domain without knowledge of implementation details.
- Centralize domain level errors for consistent handling by the application layer.

## 3. Application module
- Introduce an `app` layer to orchestrate business rules using the domain ports.
- Expose services for use by handlers and other adapters via an Fx module.
- Implement basic use cases such as user creation and donation processing.

## 4. Repository module
- Implement database repositories that satisfy the domain ports.
- Provide a module with constructors for each repository and database connection.
- Offer mocks or in-memory implementations to simplify testing.

## 5. Handler module
- Create HTTP handlers and routes for the public API.
- Attach middleware such as authentication, rate limiting and logging.
- Validate incoming requests and translate domain errors into HTTP responses.

## 6. Payment gateway integration
- Choose a payment provider and add a client package for it.
- Handle payment creation, success and failure callbacks.
- Store transactions in the repository for later auditing.

## 7. Deployment settings
- Supply Dockerfiles and compose files for local and production environments.
- Document required environment variables and configuration options.
- Add CI/CD scripts or manifests for platforms like Kubernetes.

## 8. Finished
- Complete remaining tests and review the codebase.
- Generate API documentation and update this roadmap.

