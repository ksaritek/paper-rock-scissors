# Developer Guide

That project is a monorepo, and it contains server and client side. This guide will help you to understand the project structure and how to develop and test the grpc endpoints.

## Project Structure

```shell
.
├── README.md
├── proto
│   ├── buf.gen.yaml
│   ├── buf.lock
│   ├── buf.yaml
|   ├── game
│   │   ├── game.proto
│   │   └── game_service.proto
|   internal
│   ├── api/v1
│   │   ├── start/play/end controllers
│   ├── cmd
│   │   ├── grpc server
│   │   └── signal notification
│   ├── domain
│   │   ├── model
│   │   └── repository
│   |   └── service
│   ├── logger
│   ├── observability
```

The project is structured as onion architecture. The `cmd` directory contains the main entry point of the application. The `internal` directory contains the business logic and the domain layer. The `proto` directory contains the proto files and the generated files.

Modules are loosely coupled. The `api` directory contains the controllers and the `domain` directory contains the business logic.

## Requirements

install requirements under Brewfile

```shell
$ brew bundle --file=./Brewfile
```

## proto

`proto` directory contains the required proto and buf files.

### Create buf.lock

```shell
# In the /proto directory
$ make lock
```

### How to generate?

```shell
# In the /proto directory
$ make generate
```
