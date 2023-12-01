# Tokopedia DevCamp 2023 - NSQ Queue Example

Welcome to the NSQ Queue Example for Tokopedia DevCamp 2023 participants! This repository contains a simple implementation of a queue application written in Golang using the NSQ messaging system.

## Overview

This example demonstrates how to use NSQ for building scalable and distributed message queues in a Golang application. NSQ is a real-time distributed messaging platform designed to operate at scale, handling high throughput and providing fault tolerance.

## Getting Started

### Prerequisites

Before you begin, make sure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Running the Example

1. Clone this repository

    ```bash
    git clone https://github.com/ghzmhrm/devcamp-backend-2023.git
    cd devcamp-backend-2023
    ```

2. Start the NSQ services and run the Golang NSQ producer and consumer using `docker-compose`

    ```bash
    docker-compose up --build
    ```

3. Watch the producer publishing messages and the consumer processing them.

## Acknowledgments

- NSQ: https://nsq.io/
- Tokopedia DevCamp 2023 organizers

Happy coding!
