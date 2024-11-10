# Encrypter
## Table of Contents
1. [General Info](#general-info)
2. [Technologies](#technologies)
3. [Installation](#installation)
## General Info
***
This is a project build in php whose purpose is to provide the user a way to encrypt and decrypt, assuring the security and confidentiality of the information in their documents.
## Technologies
***
A list of technologies used within the project:
* [PHP](https://www.php.net): Version 8.2
* [Apache](https://httpd.apache.org)
* [Zip](https://www.php.net/manual/es/book.zip.php)
## Installation
***
There are two methods to install this project.
### Via GitHub
#### Using Docker
Verify you are running Docker or Docker Desktop and open a terminal in the folder you want to install the application.

Copy the repository
```
git clone https://github.com/nava2105/encrypter.git
```
Enter the directory
```
cd ../encrypter
```
Build and run the container
```
docker-compose up --build
```
Open a browser and enter to
[http://localhost:8080](http://localhost:8080)
### Via Docker-hub
Pull the image from Docker-hub
```
docker pull na4va4/encrypter
```
Start a container from the image
```
docker run -p 8080:80 na4va4/encrypter
```
Open a browser and enter to
[http://localhost:8080](http://localhost:8080)
