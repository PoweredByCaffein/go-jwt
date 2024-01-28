# Golang JWT Authentication with Redis and MySQL

## Overview

This repository contains a Golang application that provides JWT authentication using Redis for storing JWT tokens and MySQL as the database. The HTTP server is built using Gin Gonic, and Cobra CLI is integrated for command-line functionalities.

## Features

- **JWT Authentication**: Secure your application with JSON Web Token authentication.
- **Redis Integration**: Efficient storage and retrieval of JWT tokens using Redis.
- **MySQL Database**: Store user data securely in a MySQL database.
- **Gin Gonic**: A fast and lightweight HTTP web framework for Go.
- **Cobra CLI**: Command-line interface for easy management and configuration.

## Getting Started

Follow these steps to set up and run the Golang JWT Authentication application:

### Prerequisites

1. [Golang](https://golang.org/doc/install): Make sure you have Go installed on your machine.
2. [Redis](https://redis.io/download): Install and run a Redis server.
3. [MySQL](https://dev.mysql.com/downloads/): Install and configure a MySQL server.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/PoweredByCaffein/go-jwt.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-jwt
    ```

3. Install dependencies:

    ```bash
    go mod download
    ```

### Configuration

1. Copy the example configuration file:

    ```bash
    cp .env.example .env
    ```

2. Update the `.env` file with your Redis and MySQL configuration.

### Usage

#### Running the Server

```bash
go run main.go server