# To-do-list
> This application is based on managing the tasks of sign in user. 

## Table of Contents
* [General Info](#general-information)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Setup](#setup)
* [Usage](#usage)
* [Project Status](#project-status)
* [Acknowledgements](#acknowledgements)


## General Information
- This project was created because I wanted to create to do list application in Golang.
- I wanted also to implement user authentication form in Golang.


## Technologies Used
- Go - version 1.19.2


## Features
List the ready features here:
- create server
- create connection with mongo database 
- implementation server and mongodb in docker


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
$ git clone https://github.com/mateuszgua/to-do-list.git
$
$ # Move into folder
$ cd cd to-do-list
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
$ git clone https://github.com/mateuszgua/to-do-list.git
$
$ # Move into folder
$ cd cd to-do-list
$
$ # Start app
$ docker-compose up --build
$ # ...
$ #app_1  | 2022/12/31 23:59:59 Starting server on http://localhost:8080 ...
```

## Project Status
Project is: in_progress 


## Acknowledgements