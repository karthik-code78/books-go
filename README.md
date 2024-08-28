# books-go

## Go version - 1.22.5

## Please set up database before running the application - the db details are in the .env files

db creation -> (Please check the values in .env files for each service)

CREATE DATABASE "name of the db in .env file in the directory `environments`";
CREATE USER "username"@"host" IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO "username"@"localhost" WITH GRANT OPTION;
FLUSH PRIVILEGES;

## Dependencies installation

### go mod tidy
### go mod download
### go mod vendor

## How to run the code (After setting up db)

-> go into `cmd` directory -> `main.go` -> run the following command

### go run main.go
-> service will run in a different port.

