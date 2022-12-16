# Http_router
> This application base on simple HTTP router using the standard package net/http from Golang. 

## Table of Contents
* [General Info](#general-information)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Setup](#setup)
* [Usage](#usage)
* [Project Status](#project-status)
* [Acknowledgements](#acknowledgements)


## General Information
- This project was created because I wanted to try to implement simple algorithm with Golang.
- I wanted to learn more about the standard routing functions. 


## Technologies Used
- Go - version 1.19.2


## Features
List the ready features here:
- code reading of HTTP server
- implementing tree structure for HTTP router
- implementing HTTP router
- exaple of use HTTP router with endpoints
- implementation in docker


## Setup
For start application with docker you need [Docker](https://docs.docker.com/get-docker/) and [docker-compose](https://docs.docker.com/compose/install/).


## Usage
The application can be build from sources or can be run in docker.

##### Build from sources
```bash
$ # Move to directory
$ cd folder/to/clone-into/
$
$ # Clone the sources
$ git clone https://github.com/mateuszgua/Http_router.git
$
$ # Move into folder
$ cd Http_router
$
$ # Start app
$ go run .
$ #2022/12/31 23:59:59 Starting server on http://localhost:8080 ...  
```

##### Start the app in Docker
```bash
$ # Move to directory
$ cd folder/to/clone-into/
$
$ # Clone the sources
$ git clone https://github.com/mateuszgua/Http_router.git
$
$ # Move into folder
$ cd Http_router
$
$ # Start app
$ docker-compose up --build
$ # ...
$ #app_1  | 2022/12/31 23:59:59 Starting server on http://localhost:8080 ...
```

## Project Status
Project is: complete 


## Acknowledgements
- This project was based on [this tutorial](https://dev.to/bmf_san/introduction-to-golang-http-router-made-with-nethttp-3nmb).