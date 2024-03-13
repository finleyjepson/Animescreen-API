# Animescreen-API

## Description

This project is a rework of the backend api from [T3A2-B-Backend](https://github.com/finleyjepson/T3A2-B-Backend) which is the backend for the [T3A-2-B](https://github.com/finleyjepson/T3A2-B-Frontend) project. The original project was built usiing JavaScript with the Express framework, with a MongoDB backend.

The reason for this project is to rewrok the backend to be less resource intensive and be more performant.
for this project i will be using the [GO language](https://go.dev/) with the Fiber framework, and a Postgres Database.

## Setup
To run this project, you will need to have the following installed:
- [GO](https://go.dev/dl/) - This was built using Go 1.22.1

Once you have Go installed, you will need to clone the repository and run the following commands:
```bash
go get
```
This will install all the required dependencies for the project.

you will also need to create an .env file in the root of the project with the following variables
```env
PORT=3000
JWT_SECRET=secret
DB_URI=postgresdburl
```
This will set the port for the server to run on, and the secret for the JWT tokens.
Please note that the JWT_SECRET should be a long random string.

The DB_URI should be the url for the postgres database you are using.

### Database setup

Once u have created a Postgres database, you will need to run the following command to create the tables for the project:
```bash
go run mian.go --seed
```

NOTE: This will also seed the database with some test data.


## Running the project

To run the project, you will need to run the following command:
```bash
go build
```
This will build the project and create an executable file. 

You can then run the executable file using the following command:
```bash
./api.exe
```