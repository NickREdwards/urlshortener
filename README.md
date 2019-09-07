## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)


## About The Project
A simple URL shortener API written in Golang with a Postgres data store.

## Getting Started
To get a local copy up and running follow the below steps.

### Prerequisites
* Docker

### Build & Run
```
cd urlshortener
docker-compose up
```

By default, the API will be exposed at localhost:8080. This can be changed in the docker-compose.yml file.

### Testing
Posting to the /api/create endpoint will return a shortened URL, which, when navigated to, will redirect to the original long URL.

#### Request body format
```
{
    "urlToShorten": "http://www.google.co.uk/search?q=something+really+long"
}
```

#### Invoking
* cURL
```
curl -H "Content-Type: application/json" -d '{"urlToShorten": "http://www.google.co.uk/search?q=something+really+long"}' http://localhost:8080/api/create
```


## License
Distributed under the MIT License. See `LICENSE` for more information.
