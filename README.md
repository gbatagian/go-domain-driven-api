# Golang Domain-Driven Design (DDD) API 
----------

### 1. About
This project demonstrates a sample DDD API architecture implemented in Golang using the standard library's `net/http` package. This architecture can be adapted to other web frameworks as well.

The DDD approach treats each domain as a separate web application. The design is modular and with a great separation of concerns, making it ideal for large and complex projects in which oftentimes there is the need to extend, replace or remove parts of the application.

### 2. Project Layout
* `app` **folder**: Acts as the project's orchestrator. It defines structures and factories to manage the entire API and gathers the endpoints (URLs) from the available domains.
* `domains` **folder**: Contains subfolders representing individual API domains. Each subfolder is named after a specific domain and contains code relevant to only that domain. A domain subfolder can be though of as a separate web application - meaning that it can fully run a service completely independently (if needed it can also have dependencies to other domains, but that is not the rule).
* `settings` **folder**: Stores global configurations for better organization and convenience.
* `utils` **folder**: Contains utility functions that can be reused across multiple domains.

### 3. Project Setup

1. Clone the project's repository: `git clone https://github.com/gbatagian/go-domain-driven-api.git`
2. Navigate to the project's root directory: `cd /go-domain-driven-api`
3. Run:
   * **Docker**:  `docker compose up`
   * **Local**: 
     * `go run main.go`  
     * `make run`  (sources `.env` before running)
4. API testing:
   * **Postman**: Import the Postman collection provided in the `docs` folder to send requests to API endpoints. 
   * **Command Line**: 
      - `curl http://localhost:8080/healthcheck`
      - `curl -X POST http://localhost:8080/greetme -d '{"name": "George","title": "Mr."}'`
5. Configuration:
   * Customize the API **address** and **port** by editing the `.env` file. Ensure the configurations are valid for your environment.
   * To **update** the address and port after running a Docker container with the API, follow these steps:
     1. Prune the network: `docker compose down && docker network prune`
     2. Restart the container: `docker compose up`
