
<br />
<div align="center">
  <h3 align="center">Go simple Online shop</h3>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project is a simple online shop Rest-API built using Go-lang, designed following the Domain-Driven Design (DDD) pattern. The project features include authentication (login and registration), product management, and transaction handling.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

This project is based on the following packages:

* [![go][go.js]][go-url]
* [![postgresql][postgresql.js]][postgresql-url]
* [![logrus][logrus.js]][logrus-url]
* [![swagger][swagger.js]][swagger-url]
* [![echo][echo.js]][echo-url]
* [![golang-migration][migration.js]][migration-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This project worker can follow the steps below:

### Prerequisites

1. [Go-lang](https://go.dev/dl/)
2. [Postgres](https://www.postgresql.org/)
3. [Swagger](https://github.com/swaggo/swag)
4. [Golang-Migrate](https://github.com/golang-migrate/migrate)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/sasmeka/backend_coffeshop_with_go.git
   ```
2. Install modules packages
   ```sh
   go mod download
   ```
3. please configure config.yml
4. Run
   ```sh
   go run ./cmd/main.go
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[go.js]: https://img.shields.io/badge/Go-1.16-blue.svg
[go-url]: https://golang.org/

[postgresql.js]: https://img.shields.io/badge/PostgreSQL-13-blue.svg
[postgresql-url]: https://www.postgresql.org/

[logrus.js]: https://img.shields.io/badge/Logrus-1.7-blue.svg
[logrus-url]: https://github.com/sirupsen/logrus

[swagger.js]: https://img.shields.io/badge/Swagger-2.0-green.svg
[swagger-url]: https://github.com/swaggo/swag

[echo.js]: https://img.shields.io/badge/Echo-4.1-lightgrey.svg
[echo-url]: https://echo.labstack.com/

[migration.js]: https://img.shields.io/badge/Migration-1.6-orange.svg
[migration-url]: https://github.com/golang-migrate/migrate