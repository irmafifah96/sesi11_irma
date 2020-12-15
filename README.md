# irmaAssignment11

Sample crud web application project using Golang(http, templates, os, sql), Bootstrap 4, DataTables, MySQL.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software

* Golang, preferably the latest version (13.11).
* MySQL Database

### Installing

1. Clone this repository

```
$ git clone https://github.com/irmafifah96/irmaAssignment11.git
$ cd irmaAssignment11
```

2. Run below command and install MySQL driver's

```
$ go get -u github.com/go-sql-driver/mysql
```

3. Create database on MySQL

```
Execute the scripts in scripts > users.sql
```

4. Edit your environment (or set_env.bat on Windows) file with your parameters and set variables for your environment

```
Change your folder
```

6. Run the application

```
$ go run main.go
```

## Acknowledgments

* This project is in development
