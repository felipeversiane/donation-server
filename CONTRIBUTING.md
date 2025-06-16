# Contributing

Thank you for considering contributing to this project. This repository uses Go and relies on automatic linting before pushing changes. The notes below explain the local setup and required coding standards.

## Local setup

1. Install Go 1.24 or newer.
2. Install [golangci-lint](https://golangci-lint.run) (version 2.x or newer). It is used by the `pre-push` hook and the CI pipeline.
3. Run `make setup` once to enable the git hooks:
   ```bash
   make setup
   ```
   This sets `core.hooksPath` to `.githooks`, enabling `.githooks/pre-push`.
4. Fetch dependencies using `go mod download` if needed.

## Coding standards

- Formatting follows `gofmt`/`goimports`. Tabs are used for Go files as defined in `.editorconfig`.
- The linter configuration lives in `.golangci.yml`. All code must pass `make lint`.
- Keep functions concise and avoid complex logic (`cyclop` and `funlen` linters enforce limits).
- Run `make test` to execute unit tests before submitting a pull request.

## Linting and git hooks

The `pre-push` hook runs:
```bash
make lint
```
If linting fails, the push will be aborted. You can run the command manually before committing to verify.

