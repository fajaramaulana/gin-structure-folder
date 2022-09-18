# Gin Clean Architecture

Clean Architecture with [Gin Web Framework](https://gin-gonic.com/)


## Usage

Setup environment variables
- Create env file
  - Key **MYSQL**
  - Key **GIN_MODE**
  - Key **PORT**
  - Key **JWT_SECRET_KEY**

```zsh
MYSQL="[YOUR_DATABASE_USERNAME]:[YOUR_DATABASE_PASS]@tcp(127.0.0.1:[YOUR MYSQL PORT])/[YOUR TABLE]?charset=utf8mb4&parseTime=True&loc=Local"
GIN_MODE=debug
PORT=":[YOUR PORT]"

JWT_SECRET_KEY= [YOUR JWT SECRET KEY]
```
## Package
- [**GIN**](https://gin-gonic.com/)
- [**GORM**](https://gorm.io/index.html)
- [**GODOTENV**](https://github.com/joho/godotenv)
- [**SLUG**](github.com/gosimple/slug)
- [**MYSQL**](github.com/go-sql-driver/mysql)
- [**validator/v10**](github.com/go-playground/validator/v10)

## Run application
### Locally

-   move to folder `/boot/`
-   Run `go run boot.go` to start the server.

---

## Folder Structure :file_folder:

| Folder Path                      | Description                                                                                         |
| -------------------------------- | --------------------------------------------------------------------------------------------------- |
| `/boot`                           | contains package main / entry-point of the server |
| `/cmd`                           | contains `controllers`, `entity`, `repositories`, `request`, and `services` |
| `/cmd/controllers`                           | The `controllers` layer is the last before the data goes to your Frontend or app. You can handle errors or map the response to accommodate your needs. |
| `/cmd/entity`                           | The `entity` layer is which stores the interface and abstraction from data. |
| `/cmd/repositories`                           | The `repositories` layer is where we Create, Get, Update or Delete data within a database. If you get data from the database, the data passing to the next layer should be raw or unprocessed data that comes directly from the database. |
| `/cmd/request`                           | The `request` is which stores the interface and abstraction from request. |
| `/cmd/services`                      | service layers, contains the functionality that compounds the core of the application               |
| `/config`                      | contains `connections`, `middleware`               |
| `/config/connections`                    |  the functionality connection to the database, redis                                                       |
| `/config/middleware`                     | the functionality for JWT validate and create token, and CORS                                                  |
| `/pkg`                       | contains all helper for the app.                   |
| `/routes`                     | contains routes
---

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)