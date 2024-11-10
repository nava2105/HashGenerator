# HashGenerator
## Table of Contents
1. [General Info](#general-info)
2. [Technologies](#technologies)
3. [Installation](#installation)
## General Info
***
This is a project build in golang whose purpose is to provide the user a way to generate the hash code af a pdf file to verify the security and authenticity of the data.
## Technologies
***
A list of technologies used within the project:
* [Golang](https://go.dev): Version 1.23
* [Alpine](https://alpinelinux.org)
## Installation
***
There are two methods to install this project.
### Via GitHub
#### Using Docker
Verify you are running Docker or Docker Desktop and open a terminal in the folder you want to install the application.

Copy the repository
```
git clone https://github.com/nava2105/HashGenerator.git
```
Enter the directory
```
cd ../HashGenerator
```
Build and run the container
```
docker-compose up --build
```
Open a browser and enter to
[http://localhost:8080/upload](http://localhost:8080/upload)
#### Not using Docker
Verify you are using Golang version 1.23 
```
go version
```
Copy the repository
```
git clone https://github.com/nava2105/HashGenerator.git
```
Enter the directory
```
cd ../HashGenerator
```
Compile the project directly form mvn
```
go run main.go
```
Open a browser and enter to
[http://localhost:8080/upload](http://localhost:8080/upload)
### Via Docker-hub
Pull the image from Docker-hub
```
docker pull na4va4/hash-generator
```
Start a container from the image
```
docker run -p 8080:8080 na4va4/hash-generator
```
Open a browser and enter to
[http://localhost:8080/upload](http://localhost:8080/upload)
