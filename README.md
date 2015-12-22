# CI Service Template for Go

![](docs/cytoscape-flat-logo-orange.png) ![](docs/gopher-side_path.png)

## What's this?

This is a sample implementation of Go microservice for Cytoscape CI.
 
Essentially, this is just a simple REST API server written in Go, and you can use your own implementation if you prefer.


## How to Run

### 1. In single API server mode
You can run this service as a single RESTful API server independent from Cytoscape CI.

#### On your local machine

* Requirments
  * Go 1.5.x and later

* Clone this repository
* ```cd data```
* Run this command to download and create mapping tables from NCBI and Uniprot:  ```python ./data_table_generator.py```
* ```cd ..```
* ```go build app.go```
* ```./app```

Then access ```http://localhost:3000/``` to check the server is actually working or not.


#### Docker Container
The easiest way to run this application is using Docker.  Suppose you are using Docker host running on ```192.168.99.100```.


### Register to _elsa_
This service works in two modes:

#### Single server mode
If you run this service without elsa instance, it works as a simple RESTful API server.  You can use your server by 
directly send your _POST_ requests to it.


#### CI service mode

## How to use the service

### REST API

#### _/_

* Function - Show service information
* Supported Methods - __GET__
* Output

```json
{
}
```
