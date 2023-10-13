# AuthSystem 🥷🏽🔐 (WIP)

<p align="center">
   <a href="http://makeapullrequest.com"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat" alt=""></a>
   <a href="https://golang.org"><img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="made-with-Go"></a>
   <a href="https://goreportcard.com/report/github.com/goodnessuc/authsystem"><img src="https://goreportcard.com/badge/github.com/goodnessuc/authsystem" alt="GoReportCard"></a>
   <a href="https://github.com/goodnessuc/authsystem"><img src="https://img.shields.io/github/go-mod/go-version/goodnessuc/authsystem.svg" alt="Go.mod version"></a>
   <a href="https://github.com/goodnessuc/authsystem/blob/master/LICENSE"><img src="https://img.shields.io/github/license/goodnessuc/authsystem.svg" alt="LICENSE"></a>
</p>

A Complete Backend Service with User Authentication and Authorization

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation and Setup](#installation-and-setup)
- [Docker Deployment](#docker-deployment)
    - [Prerequisites](#prerequisites-1)
    - [Running with Docker Compose](#running-with-docker-compose)
- [Testing](#testing)
- [Technologies and Tools](#technologies-and-tools)
- [Contributions](#contributions)
- [License](#license)

## Description

This repository contains a backend service developed in GoLang that handles user authentication and authorization.

It includes functionalities like user registration, login, password reset, social authentication (e.g., GitHub, Google),
and JWT authentication.

The service is designed with security in mind, implementing measures to prevent common attacks such as SQL injection,
CSRF attacks, and brute force attacks.

## Features

- [x] User Management
- [x] User Authentication
    - [x] User Registration: Allows new users to create an account.
    - [x] User Login: Allows existing users to log in.
    - [x] User Logout: Allows users to log out.
    - [x] RBAC: Implements Role-Based Access Control (RBAC) for users.
- [x] Session Management
    - [x] Create Session: Creates a session when a user logs in.
    - [x] Destroy Session: Destroys the session when a user logs out.
- [x] Password Management
    - [x] Password Reset: Allows users to reset their password.
    - [x] Password Hashing: Hashes passwords before storing them in the database.
- [x] Social Authentication: Allows users to log in using their social media accounts.
    - [x] GitHub Authentication: Allows users to log in through GitHub
    - [x] Google Authentication: Allows users to log in through Google
- [x] JWT Authentication: Uses JSON Web Tokens (JWT) for secure information transmission.
- [x] Security Measures
    - [x] SQL Injection Prevention: Prevents SQL injection attacks.
    - [x] CSRF Prevention: Prevents Cross-Site Request Forgery (CSRF) attacks.
    - [x] Slowloris Prevention: Prevents Slowloris Attacks

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

- Go (version 1.15 or later)
- Make

### Installation and Setup

1. Clone the repository

```bash
git clone https://github.com/goodnessuc/authsystem.git
```

2. Change directory to the project folder

```bash
cd authsystem
```

3. Install dependencies

```bash
make install
```

4. Build the application

```bash
make build
```

5. Run the application

```bash
./cmd/server/main.go
```

## Docker Deployment

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Running with Docker Compose

1. Ensure that Docker and Docker Compose are installed on your machine.
2. Navigate to the project directory:

```bash
cd authsystem
```

3. Build and run the application and its services using Docker Compose:

```bash
docker-compose up --build
```

To stop the services, simply run:

```bash
docker-compose down
```

## Testing

To run the tests, run the following command:

```bash
make test
```

## Technologies and Tools

- [Go](https://golang.org/) - The programming language used
- [GORM](https://www.gorm.io/gorm) - Database ORM tool
- [The Gin Gonic Framework](https://github.com/gin-gonic/gin) - Go based HTTP routing tool
- [JWT](https://jwt.io/) - JSON Web Tokens for authentication
- [Docker](https://www.docker.com/) - Containerization tool
- [Docker Compose](https://docs.docker.com/compose/) - Tool for defining and running multi-container Docker applications
- [PostgreSQL](https://www.postgresql.org/) - Database
- [OAuth 2.0](https://oauth.net/2/) - Open standard for access delegation
- [Make](https://www.gnu.org/software/make/) - Build automation tool

## Contributions

Contributions, issues, and feature requests are welcome!

## License

MIT

```