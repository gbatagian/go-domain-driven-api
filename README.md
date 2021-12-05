# Golang Sample API 
---------
## Domain Driven Design Pattern
----------

### 1. About
This sample project presents a custom made domain driven API architecture in Golang using the Gin Web Framework. This domain driven architecture can be extended into other web frameworks as well. Under this architecture every domain acts a separate web application. The design is modular and with a great separation of concerns, making it ideal for large and complex projects in which oftentimes there is the need to extend, replace or remove parts of the application.

### 2. Project Layout
* The `app` folder acts an orchestrator. It provides the structures and the factories to run the whole API and it collects the endpoints from the available domains.
* The `domains` folder contains all the folders that correspond to the API endpoints. Each folder inside the `domains` folder is named based on a domain and it contains the code related only to that domain. Each subfolder can be though of as a separate web application - meaning that every folder inside the `domains` folder can fully run a service completely independently (if needed it can also have dependencies to other domains, but that is not the rule).
* The `settings` folder contains some global configurations, for convenience and better organisation.
* The `utils` folder contains utilities that can be useful into more than one domains - to avoid code duplication and to create modularity and extra separation of concerns.

### 3. Project Setup

1. You need to have Golang installed on your machine.
    - If not installed, download from: https://go.dev/dl/
2. Clone the project's repository.
3. Change directory to the projects root directory.
4. Execute on terminal: `make run`
5. If you have postman you can import the postman collection present in `docs` folder and send requests to the API endpoints.
    - Alternatively, you can try the following commands on a separate terminal window:
        - `curl --location --request GET 'http://localhost:8080/healthcheck'`
        - `curl --location --request POST 'http://localhost:8080/greetme' --header 'Content-Type:application/json' --data-raw '{"name": "George","title": "Mr."}'`
6. During the development process, you can also run: `make test` or `make coverage` to test you code and get a coverage report.
