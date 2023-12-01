# Tokopedia DevCamp 2023 - Dockerized Golang Web Application

Welcome to the Dockerized Golang Web Application example for Tokopedia DevCamp 2023 participants! This repository contains a simple Golang web application that outputs a welcome message. The application is dockerized for easy deployment.

## Overview

The main components of this repository are:
- `main.go`: Golang code for a basic web server.
- `Dockerfile`: Dockerfile for building a Docker image for the Golang application.

## Getting Started

### Prerequisites

Before you begin, make sure you have the following installed:
- [Docker](https://www.docker.com/)

### Building and Running the Dockerized Web Application

1. Clone this repository:

    ```bash
    git clone https://github.com/ghzmhrm/devcamp-backend-2023.git
    cd devcamp-backend-2023
    ```

2. Build the Docker image:

    ```bash
    docker build -t dockerized-webapp .
    ```

3. Run the Docker container:

    ```bash
    docker run -p 8080:8080 dockerized-webapp
    ```

4. Open your web browser and navigate to [http://localhost:8080](http://localhost:8080). You should see the welcome message.

## Acknowledgments

- Tokopedia DevCamp 2023 organizers

Happy coding!
