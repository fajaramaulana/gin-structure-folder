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

-   Run `go run main.go app:serve` to start the server.
-   There are other commands available as well. You can run `go run main.go -help` to know about other commands available.

---

## Folder Structure :file_folder:

| Folder Path                      | Description                                                                                         |
| -------------------------------- | --------------------------------------------------------------------------------------------------- |
| `/api`                           | contains all the `middlwares`, `controllers` and `routes` of the server in their respective folders |
| `/api-errors`                    | server error handlers                                                                               |
| `/bootstrap`                     | contains modules required to start the application                                                  |
| `/console`                       | server commands, run `go run main.go -help` for all the available server commands                   |
| `/constants`                     | global application constants                                                                        |
| `/docker`                        | `docker` files required for `docker compose`                                                        |
| `/docs`                          | API endpoints documentation using `swagger`                                                         |
| `/hooks`                         | `git` hooks                                                                                         |
| `/infrastructure`                | third-party services connections like `gmail`, `firebase`, `s3-bucket`, ...                         |
| `/lib`                           | contains library code                                                                               |
| `/migration`                     | database migration files                                                                            |
| `/models`                        | ORM models                                                                                          |
| `/repository`                    | contains repository part of clean architecture. Mainly database queries are added here.             |
| `/seeds`                         | seeds for already migrated tables                                                                   |
| `/services`                      | service layers, contains the functionality that compounds the core of the application               |
| `/tests`                         | includes application tests                                                                          |
| `/utils`                         | global utility/helper functions                                                                     |
| `.env.example`                   | sample environment variables                                                                        |
| `dbconfig.yml`                   | database configuration file for `sql-migrate` command                                               |
| `docker-compose.yml`             | `docker compose` file for service application via `Docker`                                          |
| `main.go`                        | entry-point of the server                                                                           |
| `Makefile`                       | stores frequently used commands; can be invoked using `make` command                                |
| `serviceAccountKey.json.example` | sample credentials file for accessing Google Cloud                                                  |

---

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)