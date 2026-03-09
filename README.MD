# GoProductsAPI
![golang](https://github.com/bryandbernhardt/go-products-api/blob/main/assets/golang-banner.png)

![GitHub repo size](https://img.shields.io/github/repo-size/bryandbernhardt/go-products-api?style=for-the-badge)
![GitHub language count](https://img.shields.io/github/languages/count/bryandbernhardt/go-products-api?style=for-the-badge)

> Product ordering API developed in Go.

### Adjustments and improvements

- [x] Initial structuring of endpoints
- [x] Connection to PostgreSQL
- [x] Configure Docker
- [ ] Disable Product
- [ ] Edit Product
- [ ] Authentication

## 💻 Prerequisites

Before you begin, make sure you meet the following requirements:

* Go `1.25`
* Docker Compose

## 🚀 Running the API

```
docker build -t go-api .
docker compose up -d
```

## ☕ Using the API

It is recommended to check the docker compose file.
API will run local on port 8000
PostgreSQL will run local on port 5432


[⬆ Back to top](#GoProductsAPI)<br>