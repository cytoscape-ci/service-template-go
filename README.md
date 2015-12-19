# CX Tools Web Service

![](docs/cytoscape-flat-logo-orange.png) ![](docs/gopher-side_path.png)

## What's this?

This is a microservice to convert CX stream into other formats.


(TBD)

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
If you have a running instance of elsa, the registration process is (semi) automatic.  You can specify the following 
command-line options to register the service to _elsa_.

(TBD)

Options

* id - ID of this service (e.g., _idmapping_)
* ver - version of this API (e.g., _v1_)
* port - Port number of this service
* agent - Location of elsa instance


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

#### _/converter/cx2cyjs_

* Function - Convert CX into Cytoscape.js
* Supported Methods - __POST__

##### Query

```json
{
}
```

### Sample Client Code in Python

```python

import json, requests

SERVICE_URL = "http://192.168.99.100:3000/converter/cx2cyjs" # Replace this to your server location.

res = requests.post(SERVICE_URL, json=query)
print(res.json())
```

## Errors

### Invalid user data

If the query does not contain required parameters (for this version, _ids_ is the only required field) 
the service returns this error.

* Code: 400
* Body:

  ```json
  {
    "code": 400,
    "message": "Invalid query.  Probably you missed ids parameter?",
    "error": "(Any error massage from the system.)"
  }
  ```
* Example

### Unsupported Method

You will see this when you use unsupported HTTP method for an endpoint.
For example you need to use __POST__ method to call this ID Mapping service, and you will get this error 
when you simply call the URL with GET method.

* Code: 405
* Body:

  ```json
  {
    "code": 405,
    "message": "Unsupported HTTP request type.",
    "error": "You need to use POST method to use this endpoint."
  }
  ```

* Example


### Resource Not found

If the service cannot find any result, it returns this error.

* Code: 404
* Body:

  ```json
  {
    "code": 404,
    "message": "No resource found for your query.",
    "error": "No maching IDs for your inputs"
  }
  ```

### Unexpected server side error

Usually this should not happen.  In Go, if critical panic happens due to bugs, this will be returned to the user.

* Code: 500
* Body:

  ```json
  {
    "code": 500,
    "message": "Something wrong happened to the service.  Now is the good time to call admin...",
    "error": "(stack trace, panic message, heap dump, etc.)"
  }
  ```
