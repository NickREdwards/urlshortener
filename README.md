## Table of Contents

* [About the Project](#about-the-project)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Build & Run](#build--run)
* [API](#api)
* [Design Decisions](#design-decisions)
* [Potential Improvements](#potential-improvements)
* [License](#license)


## About The Project
A simple URL shortener API written in Golang with a Postgres database.

## Getting Started
To get a local copy up and running follow the below steps.

### Prerequisites
* Docker

### Build & Run
```
cd urlshortener
docker-compose up
```

By default, the API will be exposed at localhost:8080. Postgres is also exposed outside of the Docker network for debugging purposes on port 5432.
These can be changed in the docker-compose.yml file as required.

## API
### Shorten URL
#### *POST* /api/create
```
{
    "urlToShorten": "http://www.google.co.uk/search?q=something+really+long"
}
```
#### Response
```
{
    "shortenedUrl": "localhost:8080/r/HK5a84"
}
```

#### Invoke
```
curl -H "Content-Type: application/json" -d '{"urlToShorten": "http://www.google.co.uk/search?q=something+really+long"}' http://localhost:8080/api/create
```

### Access Logs
#### *GET* /api/access_logs/{shortCode}?from={from}&to={to}
*from and to are optional and can be independently ommitted*

#### Response
```
{
    "shortCode": "ABCDE",
    "total": 5,
    "logs": [
        { "dateTime": "2019-09-07T17:02:24.123.000000Z" },
        { "dateTime": "2019-09-07T17:03:25.345.000000Z" },
        { "dateTime": "2019-09-07T17:05:32.754.000000Z" },
        { "dateTime": "2019-09-07T17:09:44.876.000000Z" },
        { "dateTime": "2019-09-07T17:18:56.345.000000Z" }
    ]
}
```
#### Invoke
```
curl http://localhost:8080/api/access_logs/ABCDE
curl http://localhost:8080/api/access_logs/ABCDE?from=2019-09-07+17:00:00
curl http://localhost:8080/api/access_logs/ABCDE?to=2019-09-07+17:30:00
curl http://localhost:8080/api/access_logs/ABCDE?from=2019-09-07+17:00:00.000&to=2019-09-07T17:30:00
```

## Design Decisions
* Docker
    * Cross platform and easy to build/run/test all components of the service
* Go
    * Fast, lightweight
* PostgreSQL data storage
    * Created in a Docker container for the purposes of this exercise
    * Simple database capable of meeting the persistent storage requirements for the solution
    * Read/access speed is easily scaled up with the introduction of a non-clustered index on the short code
* The DAL is abstracted away to allow for the data store to be switched out as requirements change (see potential improvements)
* Short code generation
    * 6 characters in length, 62 possible characters = 6^62 permutations
    * With such a large number of possible permutations it's very unlikely that we'd see duplicates. Tests prove that generating 100k codes one after the other does not produce duplicates
    * Quick to generate
* Only store the short code rather than the short URL
    * Saves on space
    * Easier to index
    * Allows for a potential change of host in the future
* Access to a shortened URL is logged in-process
    * While this may degrade performance of URL resolution, it is necessary without having a messaging system in place for guaranteed delivery/eventual consistency
    * Access logs are not indexed as I am prioritizing write performance over read performance
    * At the current scale of the system URL resolution has been tested to return in < 10ms
* Access logs API will return a response regardless of if that code does not exist to help prevent code sniffing

## Potential Improvements
* Introduction of a caching layer (Redis, Memcached) between the API and Database
    * Postgres performs well at the current scale, but the introduction of a caching layer would offer much more flexibility in terms of scale and performance
* Introduction of a message broker (RabbitMQ, Kafka) for recording access logs
    * With a caching layer between API and the DB, the message broker would also allow for Postgres to act as cold storage, fed by the messaging system
* Move configuration values (e.g. Postgres connection details) out of the code
* Allow more flexibility in terms of date format when accessing logs

## License
Distributed under the MIT License. See `LICENSE` for more information.
