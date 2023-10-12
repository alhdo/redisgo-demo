<h1 align="center">Welcome to redisgo-demo 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-0.0.1-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/castroalhdo" target="_blank">
    <img alt="Twitter: castroalhdo" src="https://img.shields.io/twitter/follow/castroalhdo.svg?style=social" />
  </a>
</p>

> Demonstration of using redis pub sub cannal with go an vue js

## Install

The project consist of 3 components :
  - `simple-go-worker` A GO worker
  - `api` A NodeJS API
  - `vlille` A front-end using Vuejs

To run the project you can either run each component alone and update the env variable accordingly or use `docker-compose.yml` file to orchestrate all the 3 components.
### Running each component independently
```sh
make install
```

## Usage

```sh
make run
```

## Run tests

```sh
make test
```

## Author

👤 **Alhdo**

* Twitter: [@castroalhdo](https://twitter.com/castroalhdo)
* Github: [@alhdo](https://github.com/alhdo)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_